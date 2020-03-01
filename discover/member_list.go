package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/hashicorp/memberlist"
	"github.com/pborman/uuid"
)

func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

var (
	MyNode      *memberlist.Node
	mtx         sync.RWMutex
	serviceName = getEnv("SERVICE_NAME", "notset")
	members     = getEnv("MEMBERS_LIST", "")
	port        = getEnv("PORT", "4001")
	proto_port  = getEnv("PROTO_PORT", "5001")
	items       = map[string]string{}
	broadcasts  *memberlist.TransmitLimitedQueue
)

type broadcast struct {
	msg    []byte
	notify chan<- struct{}
}

type delegate struct{}

type update struct {
	Action string // add, del
	Data   map[string]string
}

func init() {
	flag.Parse()
}

func (b *broadcast) Invalidates(other memberlist.Broadcast) bool {
	return false
}

func (b *broadcast) Message() []byte {
	return b.msg
}

func (b *broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}

func (d *delegate) NodeMeta(limit int) []byte {
	return []byte{}
}

func (d *delegate) NotifyMsg(b []byte) {
	if len(b) == 0 {
		return
	}

	switch b[0] {
	case 'd': // data
		var updates []*update
		if err := json.Unmarshal(b[1:], &updates); err != nil {
			return
		}
		mtx.Lock()
		for _, u := range updates {
			for k, v := range u.Data {
				switch u.Action {
				case "add":
					items[k] = v
				case "del":
					delete(items, k)
				}
			}
		}
		mtx.Unlock()
	}
}

func (d *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return broadcasts.GetBroadcasts(overhead, limit)
}

func (d *delegate) LocalState(join bool) []byte {
	mtx.RLock()
	m := items
	mtx.RUnlock()
	b, _ := json.Marshal(m)
	return b
}

func (d *delegate) MergeRemoteState(buf []byte, join bool) {
	if len(buf) == 0 {
		return
	}
	if !join {
		return
	}
	var m map[string]string
	if err := json.Unmarshal(buf, &m); err != nil {
		return
	}
	mtx.Lock()
	for k, v := range m {
		items[k] = v
	}
	mtx.Unlock()
}

type eventDelegate struct{}

func (ed *eventDelegate) NotifyJoin(node *memberlist.Node) {
	fmt.Printf("A node has joined: %s / %s:%d\n", node.String(), node.Addr.String(), node.Port)
}

func (ed *eventDelegate) NotifyLeave(node *memberlist.Node) {
	fmt.Println("A node has left: " + node.String())
}

func (ed *eventDelegate) NotifyUpdate(node *memberlist.Node) {
	fmt.Println("A node was updated: " + node.String())
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := r.Form.Get("val")
	mtx.Lock()
	items[key] = val
	mtx.Unlock()

	b, err := json.Marshal([]*update{
		{
			Action: "add",
			Data: map[string]string{
				key: val,
			},
		},
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	broadcasts.QueueBroadcast(&broadcast{
		msg:    append([]byte("d"), b...),
		notify: nil,
	})
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	mtx.Lock()
	delete(items, key)
	mtx.Unlock()

	b, err := json.Marshal([]*update{{
		Action: "del",
		Data: map[string]string{
			key: "",
		},
	}})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	broadcasts.QueueBroadcast(&broadcast{
		msg:    append([]byte("d"), b...),
		notify: nil,
	})
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	mtx.RLock()
	val := items[key]
	mtx.RUnlock()
	w.Write([]byte(val))
}

var mList *memberlist.Memberlist

func start() (*memberlist.Node, error) {
	var err error
	hostname, _ := os.Hostname()
	c := memberlist.DefaultWANConfig()
	c.Events = &eventDelegate{}
	c.Delegate = &delegate{}
	c.BindPort, _ = strconv.Atoi(proto_port)
	c.Name = serviceName + "-" + hostname + "-" + uuid.NewUUID().String()
	mList, err = memberlist.Create(c)
	if err != nil {
		return nil, err
	}
	if members != "" {
		parts := strings.Split(members, ",")
		_, err := mList.Join(parts)
		if err != nil {
			if !strings.Contains(err.Error(), "connection refused") {
				return nil, err
			}
		}
	}
	broadcasts = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return mList.NumMembers()
		},
		RetransmitMult: 3,
	}
	MyNode = mList.LocalNode()
	fmt.Printf("Local member %s:%d\n", MyNode.Addr, MyNode.Port)
	return MyNode, nil
}

//func main() {
//	if err := start(); err != nil {
//		fmt.Println(err)
//	}
//
//	http.HandleFunc("/add", addHandler)
//	http.HandleFunc("/del", delHandler)
//	http.HandleFunc("/get", getHandler)
//	fmt.Printf("Listening on :%s\n", port)
//	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
//		fmt.Println(err)
//	}
//}
