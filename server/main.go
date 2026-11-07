package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	port := flag.Int("port", 6380, "port to listen on")
	flag.Parse()
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go handle(c)
	}
}

// readCommand parses one RESP array of bulk strings.
func readCommand(r *bufio.Reader) ([]string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) == 0 || line[0] != '*' {
		return nil, fmt.Errorf("expected array")
	}
	n, _ := strconv.Atoi(line[1:])
	args := make([]string, 0, n)
	for i := 0; i < n; i++ {
		hdr, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		l, _ := strconv.Atoi(strings.TrimRight(hdr, "\r\n")[1:])
		buf := make([]byte, l+2)
		if _, err := r.Read(buf[:0]); err != nil {
			return nil, err
		}
		got := 0
		for got < l+2 {
			m, err := r.Read(buf[got:])
			if err != nil {
				return nil, err
			}
			got += m
		}
		args = append(args, string(buf[:l]))
	}
	return args, nil
}

func handle(c net.Conn) {
	defer c.Close()
	r := bufio.NewReader(c)
	for {
		args, err := readCommand(r)
		if err != nil {
			return
		}
		if len(args) > 0 && strings.EqualFold(args[0], "PING") {
			c.Write([]byte("+PONG\r\n"))
			continue
		}
		c.Write([]byte("-ERR unknown command\r\n"))
	}
}
