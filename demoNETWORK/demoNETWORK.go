package demoNETWORK

import (
	"fmt"
	"net"
)

func UDPClientUsage() {
	var msg string

	msg = "hello, world!"

	addr, err := net.ResolveUDPAddr("udp", ":49736")
	if err != nil {
		fmt.Println("demoUDPClient ResolveUDPAddr fail")
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("demoUDPClient DialUDP fail")
		return
	}

	conn.Write([]byte(msg))

	conn.Close()
}

func UDPServerUsage() {

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("demoUDPServer ResolveUDPAddr fail")
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("demoUDPServer ListenUDP fail")
		conn.Close()
		return
	}

	for {
		data := make([]byte, 20)
		_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("demoUDPServer ReadFronUDP fail")
			conn.Close()
			return
		}

		fmt.Println("Recv UDP packets from ", rAddr.IP.String())
		fmt.Println("Contents: ", string(data))
	}
}

func TCPClientUsage() {

	var msg string

	msg = "hello, world"

	s, err := net.Dial("tcp", "127.0.0.1:52335")
	if err != nil {
		fmt.Println("demoTCPClient Dial fail", err)
		return
	}

	count, err := s.Write([]byte(msg))
	if err != nil {
		fmt.Println("demoTCPClient Write fail")
		return
	}

	fmt.Println("count = ", count)

	s.Close()
}

func TCPServerUsage() {

	var b []byte

	b = make([]byte, 20)

	fmt.Println("TCP Server at localhost:8080, only 20 bytes allowed to receive =.=!")

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("demoTCPServer Listen fail")
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("demoTCPServer Accept fail")
			ln.Close()
			return
		}

		count, err := c.Read(b)
		if err != nil {
			fmt.Println("demoTCPServer Read fail")
			ln.Close()
			return
		}

		fmt.Println("count = ", count)
		fmt.Println("content = ", string(b))
	}
}
