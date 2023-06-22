package main

import (
	"io"
	"log"

	"github.com/mo3et/hello-intranet-penetration/define"
	"github.com/mo3et/hello-intranet-penetration/helper"
)

// Client function

// - Connect to local Server
// - Connect to server provided by tunnel
// - Message Forward

func main() {
	conn, err := helper.CreateConn(define.ControlServerAddr)
	if err != nil {
		panic(err)
	}
	for {
		b, err := helper.GetDataFromConnection(define.BufSize, conn)
		if err != nil {
			log.Printf("Get Data Error %v", err)
			continue
		}

		// New Connection
		if string(b) == define.NewConnection {
			go messageForward()

		}
	}
}

func messageForward() {
	// Connection Server tunnel
	tunnelConn, err := helper.CreateConn(define.TunnelServerAddr)
	if err != nil {
		panic(err)
	}

	// Connection Client server
	localConn, err := helper.CreateConn(define.LocalServerAddr)
	if err != nil {
		panic(err)
	}
	// Message Forward
	go io.Copy(localConn, tunnelConn)
	go io.Copy(tunnelConn, localConn)
}
