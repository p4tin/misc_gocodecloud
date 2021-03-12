//lib_disco.go
// - go build -buildmode=c-shared -o lib_disco.so .
package main

import (
	"C"
	"fmt"
	"os"
	"time"
)

var serviceAddresses = map[string]string{
	"checkout": "1.2.3.4",
	"pricing":  "1.2.3.5",
}

func timedEvent() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(listAllNodes())
	}

}

//export startDiscovery
func startDiscovery() {
	var err error
	// Initial setup for Discovery / memberlist stuff
	// Timed events can be sent here + AWS topic-queue messages /GCP pubsub messages
	MyNode, err = start()
	if err != nil {
		panic(err)
	}
	go timedEvent()
}

//export getServiceURL
func getServiceURL(name *C.char) *C.char {
	addr := serviceAddresses[C.GoString(name)]
	return C.CString(addr)
}

//export listAllNodes
func listAllNodes() []string {
	nodeNames := make([]string, 0)
	nodes := mList.Members()
	for _, node := range nodes {
		if node != mList.LocalNode() && mList.GetNodeState(node.Name) == 0 {
			nodeNames = append(nodeNames, node.Name)
		}
	}
	return nodeNames
}

func main() {
	// catch exit signals and leave the cluster
	stop := make(chan os.Signal, 1)
	//signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	startDiscovery()
	fmt.Println(listAllNodes())

	<-stop

	// Leave the cluster with a 5 second timeout. If leaving takes more than 5
	// seconds we return.
	if err := mList.Leave(time.Second * 5); err != nil {
		panic(err)
	}
}
