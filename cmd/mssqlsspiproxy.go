package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/glards/mssqlsspiproxy"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Print("MSSQL SSPI Proxy")

	conf := configure()

	l, err := net.Listen("tcp", fmt.Sprint("0.0.0.0:", conf.ListenPort))
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		go proxy(conn, conf)
	}
}

func configure() *Configuration {
	conf := NewConfiguration()
	return conf
}

func proxy(client net.Conn, config *Configuration) {
	server, err := net.Dial("tcp", config.ServerDial)
	if err != nil {
		client.Close()
		return
	}

	p := tds.NewProxy(server, client)
	go p.Handle()
}
