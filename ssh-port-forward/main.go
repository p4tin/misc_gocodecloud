package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/labstack/echo"

	gossh "golang.org/x/crypto/ssh"
)

func getPrivateKey(file string) (ssh.Signer, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := gossh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return key, nil
}

var router *echo.Echo

func main() {
	signer, err := getPrivateKey(os.Getenv("SSH_HOST_PRIVATE_KEY_FILE"))
	if err != nil {
		panic(err)
	}

	server := ssh.Server{
		LocalPortForwardingCallback: ssh.LocalPortForwardingCallback(func(ctx ssh.Context, dhost string, dport uint32) bool {
			log.Println("Accepted forward", dhost, dport)
			return true
		}),
		HostSigners: []ssh.Signer{signer},
		Addr:        ":2222",
		Handler: ssh.Handler(func(s ssh.Session) {
			ticker := time.NewTicker(30 * time.Second)
			io.WriteString(s, "Remote forwarding available...\n")
			log.Println("User", s.User(), "is connected")
			for {
				select {
				case <-ticker.C:
					io.WriteString(s, "Time Elapsed message\n")
				}
			}
		}),
		ConnCallback: func(conn net.Conn) net.Conn {
			log.Printf("Connect: %+v\n", conn)
			return conn
		},
		ReversePortForwardingCallback: ssh.ReversePortForwardingCallback(func(ctx ssh.Context, host string, port uint32) bool {
			log.Println("attempt to bind", host, port, "granted")
			log.Printf("Context: %+v\n", ctx)
			//route := host
			//if host != "" || host != "localhost" {
			//	host = ""
			//}
			return true
		}),
	}
	log.Println("starting ssh server on port 2222...")
	go server.ListenAndServe()

	router = echo.New()
	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	log.Println("starting http server on port 9090...")
	router.Start(":9090")
}
