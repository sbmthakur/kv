package main

import (
    "fmt"
    "io"
    "net"
    "strings"
    "dict/dict"
    //"reflect"
)

func handleConnection(c net.Conn, con_count int, d dict.Dictionary) {
    defer c.Close()

    for {
        fmt.Println("Connection number: %v", con_count)

        data := make([]byte, 0, 1024)

        for {
            buf := make([]byte, 20)
            n, err := c.Read(buf)

            if err != nil {

                if err == io.EOF {
                    fmt.Println("Connection %d terminated", con_count)
                    return
                } 

                fmt.Println("Connection read error", err)
            }

            str_data := string(buf)

            fmt.Println("ddd", str_data)

            data = append(data, buf[:n]...)

            if strings.Contains(str_data, "\r\n") {
                break
            }
        }


        cmds := strings.Split(string(data), " ")
        cmd_name, cmd_key := cmds[0], cmds[1]
        cmd_data := ""

        if len(cmds) == 2 {
            cmd_key = strings.Replace(cmd_key, "\r\n", "", 1)
        }

        if len(cmds) == 3 {
            cmd_data = cmds[2]
            cmd_data = strings.Replace(cmd_data, "\r\n", "", 1)
        }

        fmt.Println("Command received %s", cmd_name)

        if cmd_name == "set" {
            //fmt.Println(cmd_key, cmd_data)
            e := d.Add(cmd_key, cmd_data)
            if e != nil {
                _, er := c.Write([]byte("NOT-STORED\r\n"))
                fmt.Println("key set error", er)
            } else {
                _, er := c.Write([]byte("STORED\r\n"))
                if er != nil {
                    fmt.Println("connection write error with set!")
                }
            }
        }

        if cmd_name == "get" {
            val, e := d.Search(cmd_key)

            if e != nil {
                fmt.Println("Key %s not found", cmd_key)
            } else {
                _, er := c.Write([]byte(val))
                if er != nil {
                    fmt.Println("connection write error with get!")
                }
            }
        }
    }
}

func main() {
    l, err := net.Listen("tcp", ":9002")

    if err != nil {
        fmt.Println(err)
        return
    }

    d := dict.Dictionary{}

    defer l.Close()

    count := 0

    for {

        count += 1
        //fmt.Println(count)
        c, err := l.Accept()

        if err != nil {
            fmt.Println(err)
            return
        }

        go handleConnection(c, count, d)
    }
}


