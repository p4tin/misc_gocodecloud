package main

import (
	"fmt"
	"log"
	"net"

	"crypto/tls"

	"crypto/x509"
	"io/ioutil"

	"gopkg.in/mgo.v2"
)

// Person ...
type Person struct {
	Hero  string
	Skill int
}

func preferSslConnect() (*mgo.Session, error) {
	// --sslCAFile
	rootCerts := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./mongo-ssl/ca_cert.pem")
	if err == nil {
		rootCerts.AppendCertsFromPEM(ca)
	}

	// --sslPEMKeyFile
	clientCerts := []tls.Certificate{}
	if cert, err := tls.LoadX509KeyPair("./mongo-ssl/client.crt", "./mongo-ssl/client.key"); err == nil {
		clientCerts = append(clientCerts, cert)
	}

	// Dial with TLS
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"d1shmondb01.dev.ecomm.urbanout.com:27017"},
		Database: "test",
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{
				RootCAs:      rootCerts,
				Certificates: clientCerts,
			})
		},
	})

	return session, err
}

func noSslConnect() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session, err
}

func main() {
	// Connect
	session, err := preferSslConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	fmt.Println("Connected to mongo with SSL")

	//// Drop collection
	//c := session.DB("test").C("mgo_test")
	//c.DropCollection()
	//
	//// Insert a struct
	//err = c.Insert(Person{"John Rambo", 90})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Insert a series of docs
	//docs := make([]interface{}, 5)
	//for i := 0; i < 5; i++ {
	//	docs[i] = bson.M{"a": i, "b": i * i, "c": math.Pow(2, float64(i))}
	//}
	//err = c.Insert(docs...)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Get results
	//result := Person{}
	//err = c.Find(bson.M{"hero": "John Rambo"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
}
