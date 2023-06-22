package test

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
	"testing"

	"github.com/mo3et/hello-intranet-penetration/helper"
)

const (
	ControlServerAddr = "0.0.0.0:8080"
	RequestServerAddr = "0.0.0.0:8081"
	// KeepAliveStr      = "KeepAlive\n"
)

var wg sync.WaitGroup
var clientConn *net.TCPConn

// Server
func TestUserServer(t *testing.T) {
	wg.Add(1)
	// 监听控制中心
	go ControlServer()
	// 监听用户的请求
	go RequestServer()

	wg.Wait()
}
func ControlServer() {
	tcpListener, err := helper.CreateListen(ControlServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("ControlServer is listening on %s\n", ControlServerAddr)
	for {
		clientConn, err = tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		go helper.KeepAlive(clientConn)
	}
}

func RequestServer() {
	tcpListener, err := helper.CreateListen(RequestServerAddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		go func() {
			if _, err := io.Copy(conn, clientConn); err != nil {
				panic(err)
			}
		}()
		go func() {
			if _, err := io.Copy(clientConn, conn); err != nil {
				panic(err)
			}
		}()

	}
}

// Client
func TestUserClient(t *testing.T) {
	conn, err := helper.CreateConn(ControlServerAddr)
	if err != nil {
		log.Printf("[Connection Fail.] %s", err)
	}
	for {
		s, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Get Data Err: %v", err)
			continue
			// return
		}
		log.Printf("Get Data: %v", s)
	}
}

// User request client, Send data to Server.
func TestUserRequestClient(t *testing.T) {
	conn, err := helper.CreateConn(RequestServerAddr)
	if err != nil {
		log.Printf("[Connection Fail.] %s", err)
	}
	_, err = conn.Write([]byte("KeepAliveStr \n"))
	if err != nil {
		log.Printf("[Send Fail] %s", err)
	}

}
