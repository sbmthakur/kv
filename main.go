package main

import (
    "fmt"
    "net"
    "strings"
    "dict/dict"
)

func handleConnection(c net.Conn, con_count int) {
    defer c.Close()
    for {
        fmt.Println("Connection number: %v", con_count)

        data := make([]byte, 0, 1024)

        for {
            buf := make([]byte, 20)
            n, err := c.Read(buf)

            if err != nil {
                fmt.Println("Connection read error")
            }
            d := string(buf)

            data = append(data, buf[:n]...)

            if strings.Contains(d, "\r\n") {
                break
            }
        }


        cmds := strings.Split(string(data), " ")

        /*
        cmd_name := cmds[0]

        cmd_data = cmds[1]

        cmd_data := strings.Replace(cmd_data, "\r\n", "", 1)
        */

        fmt.Println(cmds)

        /*
        if strings.Contains(data, "exit") {
            fmt.Println("got exit!")
            return
        }
        */
    }
}

func main() {
    l, err := net.Listen("tcp", ":9002")

    //d := dict.Dictionary{ "test": "test_val" }

    fmt.Println(d)

    if err != nil {
        fmt.Println(err)
        return
    }

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

        go handleConnection(c, count)
    }
}


