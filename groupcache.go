// Simple groupcache example: https://github.com/golang/groupcache
// Running 3 instances:
// go run groupcache.go -addr=:8080 -pool=http://127.0.0.1:8080,http://127.0.0.1:8081,http://127.0.0.1:8082
// go run groupcache.go -addr=:8081 -pool=http://127.0.0.1:8081,http://127.0.0.1:8080,http://127.0.0.1:8082
// go run groupcache.go -addr=:8082 -pool=http://127.0.0.1:8082,http://127.0.0.1:8080,http://127.0.0.1:8081
// Testing:
// curl localhost:8080/color?name=red
package main

import (
	"time"
	//"errors"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/golang/groupcache"
	"github.com/hashicorp/memberlist"
	"fmt"
)

var list *memberlist.Memberlist
var name *string
var addr *string
var poolArr []string
var pool *groupcache.HTTPPool

var Store = map[string][]byte{
	"red":   []byte("#FF0000"),
	"green": []byte("#00FF00"),
	"blue":  []byte("#0000FF"),
}
//
//var Group = groupcache.NewGroup("foobar", 64<<20, groupcache.GetterFunc(
//	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
//		log.Println("looking up", key)
//		v, ok := Store[key]
//		if !ok {
//			//Load from DB or Error
//			return errors.New("color not found")
//		}
//		dest.SetBytes(v)
//		return nil
//	},
//))

func broadcast(key string) {
	fmt.Println("NumMembers:", list.NumMembers())
	for _, member := range list.Members() {

		url:= fmt.Sprintf("http://%s:%d/replicate?name=%s&code=%s", member.Addr, member.Port, key, string(Store[key]))
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}

	}
}

func main() {
	addr = flag.String("addr", ":8080", "server address")
	name = flag.String("name", "MemberList-Node-01", "memberlist node name")
	peers := flag.String("pool", "localhost", "server pool list")
	memberPort := flag.Int("memberport", 7946, "memberlist port number")
	//join := flag.String("join", "10.27.10.224:7946", "current member address to join memberlist")
	flag.Parse()

	var err error
	//str := *addr
	//advertPort, _ := strconv.Atoi(str[1:])
	list, err = memberlist.Create(&memberlist.Config{
		Name: *name,
		BindPort: *memberPort,
		BindAddr: "10.27.10.224",
		ProtocolVersion: 4,
		TCPTimeout: 500 * time.Millisecond,
		PushPullInterval: 5 * time.Second,
		//AdvertiseAddr: "10.27.10.224",
		//AdvertisePort: advertPort,
	})
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	if len(*peers) > 0 {
		parts := strings.Split(*peers, ",")
		_, err = list.Join(parts)
		if err != nil {
			panic("Failed to join cluster: " + err.Error())
		}
	}

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s %d\n", member.Name, member.Addr.String(), member.Port)
	}


	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		color := strings.TrimSpace(r.FormValue("name"))
		b, ok := Store[color]
		if !ok {
			w.Write([]byte("Not Found\n"))
			return
		}
		//err := Group.Get(nil, color, groupcache.AllocatingByteSliceSink(&b))
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusNotFound)
		//	return
		//}
		w.Write(b)
		w.Write([]byte{'\n'})
	})
	http.HandleFunc("/replicate", func(w http.ResponseWriter, r *http.Request) {
		color := strings.TrimSpace(r.FormValue("name"))
		code := strings.TrimSpace(r.FormValue("code"))
		Store[color] = []byte(code)
		str := color + " added or modified."
		w.Write([]byte(str))
	})
	http.HandleFunc("/addcolor", func(w http.ResponseWriter, r *http.Request) {
		color := strings.TrimSpace(r.FormValue("name"))
		code := strings.TrimSpace(r.FormValue("code"))
		Store[color] = []byte(code)
		str := color + " added or modified."

		w.Write([]byte(str))
	})
	http.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		// Ask for members of the cluster
		str := ""
		for _, member := range list.Members() {
			str = fmt.Sprintf("%sMember: %s %s %d\n", str, member.Name, member.Addr, member.Port)
		}
		w.Write([]byte(str))
	})
	poolArr = strings.Split(*peers, ",")
	pool = groupcache.NewHTTPPool(poolArr[0])
	pool.Set(poolArr...)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
