package main

import (
	"encoding/json"
	"fmt"
	"time"

	"crypto/tls"
	"log"
	"net"

	"gopkg.in/mgo.v2"
)

type User struct {
	UserId    string `json:"userId"`
	Yesterday string `json:"yesterday"`
	Today     string `json:"today"`
	Blockers  string `json:"blockers"`
}

type Team struct {
	GroupId string `json:"groupId"`
	Date    string
	Users   []User
}

func main() {
	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"ds039404.mlab.com:39404"},
		Timeout:  5 * time.Second,
		Username: "admin",
		Password: "p8ssw0rd",
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			log.Println(err)
		}
		return conn, err
	}

	//info := &mgo.DialInfo{
	//	Addrs:    []string{"ds039404.mlab.com:39404"},
	//	Timeout:  60 * time.Second,
	//	Database: "gostandups",
	//	Username: "admin",
	//	Password: "p8ssw0rd",
	//}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("gostandups").C("report-data")

	u1 := User{UserId: "1", Yesterday: "msg1", Today: "msg2", Blockers: "msg3"}
	u2 := User{UserId: "2", Yesterday: "msg4", Today: "msg5", Blockers: "msg6"}
	u3 := User{UserId: "3", Yesterday: "msg7", Today: "msg8", Blockers: "msg9"}
	usrs := make([]User, 0)
	usrs = append(usrs, u1)
	usrs = append(usrs, u2)
	usrs = append(usrs, u3)
	t1 := Team{GroupId: "Green", Date: "02/26/2016", Users: usrs}

	_, err = c.UpsertId(t1.GroupId+"-"+t1.Date, &t1)
	if err != nil {
		panic(err)
	}

	// Query One
	result := Team{}
	err = c.FindId(t1.GroupId + "-" + t1.Date).One(&result)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Result ->", result)

	jsonString, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(string(jsonString))
}
