package main

import (
	"dict/dict"
	"flag"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func connWriter(c net.Conn, message string) {
	c.Write([]byte(message))
}

func handleSet(c net.Conn, cmds []string, d dict.Dictionary) string {
	valsize, _ := strconv.Atoi(cmds[4])
	valbuf := make([]byte, valsize+2)
	c.Read(valbuf)
	cmd_key := cmds[1]
	cmd_data := string(valbuf)[0:valsize]
	e := d.Add(cmd_key, cmd_data)
	if e != nil {
		return "NOT-STORED\r\n"
	} else {
		return "STORED\r\n"
	}
}

func handleGet(cmds []string, d dict.Dictionary) (string, string) {
	cmd_key := cmds[1]
	val, _ := d.Search(cmd_key)
	ws := "VALUE " + cmd_key + " 0 " + strconv.Itoa(len(val)) + "\r\n"
	return ws, val
}

func handleConnection(c net.Conn, con_count int, d dict.Dictionary) {
	defer c.Close()
	fmt.Println("Connection number:", con_count)

	for {

		data := make([]byte, 0, 1024)

		for {
			buf := make([]byte, 20)
			n, err := c.Read(buf)

			if err != nil {

				if err == io.EOF {
					fmt.Println("Connection ", con_count, " terminated")
					return
				}

				//fmt.Println("Connection read error", err)
			}

			fmt.Println(buf)
			str_data := string(buf)

			fmt.Println("ddd", str_data)

			data = append(data, buf[:n]...)

			if strings.Contains(str_data, "\r\n") {
				//fmt.Println("Got new line chars!")
				break
			}
		}

		input_string := strings.Replace(string(data), "\r\n", "", 1)
		//fmt.Println("input str", input_string)
		cmds := strings.Split(input_string, " ")

		cmd_name := cmds[0]

		if cmd_name == "set" {
			res := handleSet(c, cmds, d)
			connWriter(c, res)
		}

		if cmd_name == "get" {
			info, val := handleGet(cmds, d)
			connWriter(c, info)
			connWriter(c, val)
			connWriter(c, "END\r\n")
		}
	}
}

func main() {

	portPtr := flag.String("port", "4000", "port number to be used")
	flag.Parse()

	port := ":" + *portPtr

	fmt.Println("Using port ", port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println(err)
		return
	}

	d := dict.Dictionary{}

	defer l.Close()

	count := 0

	for {

		count += 1
		c, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c, count, d)
	}
}
