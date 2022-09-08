package main

import (
    "fmt"
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
                fmt.Println("Connection read error")
            }
            str_data := string(buf)

            data = append(data, buf[:n]...)

            if strings.Contains(str_data, "\r\n") {
                break
            }
        }


        cmds := strings.Split(string(data), " ")
        fmt.Println("len %d", len(cmds))
        cmd_name, cmd_key := cmds[0], cmds[1]
        cmd_data := ""

        if len(cmds) == 2 {
            cmd_key = strings.Replace(cmd_key, "\r\n", "", 1)
        }

        if len(cmds) == 3 {
            fmt.Println(cmd_data)
            cmd_data = cmds[2]
            cmd_data = strings.Replace(cmd_data, "\r\n", "", 1)
            fmt.Println("1234")
            fmt.Println(cmd_data)
        }

        if cmd_name == "set" {
            fmt.Println(cmd_key, cmd_data)
            e := d.Add(cmd_key, cmd_data)
            if e != nil {
                fmt.Println("key set error")
            } else {
                _, er := c.Write([]byte("STORED"))
                if er != nil {
                    fmt.Println("SSSSSTORED")
                }
            }
        }

        v, e := d.Search(cmd_key)
        fmt.Println("%%")

        if e == nil {
            fmt.Println(v)

        } else {
            fmt.Println("Key not found %s", cmd_name)
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


