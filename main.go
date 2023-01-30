package main

import (
    "fmt"
    "flag"
//    "io"
    "net"
    "strings"
    "dict/dict"
    //"os"
    //"reflect"
    "strconv"
)

func handleConnection(c net.Conn, con_count int, d dict.Dictionary) {
    defer c.Close()
    fmt.Println("Connection number:", con_count)

    for {

        data := make([]byte, 0, 1024)

        for {
            buf := make([]byte, 20)
            n, err := c.Read(buf)

            if err != nil {

                /*
                if err == io.EOF {
                    fmt.Println("Connection %d terminated", con_count)
                    return
                } 

                fmt.Println("Connection read error", err)
                */
            }

            fmt.Println(buf)
            str_data := string(buf)

            fmt.Println("ddd", str_data)

            data = append(data, buf[:n]...)

            if strings.Contains(str_data, "\r\n") {
                fmt.Println("Got new line chars!")
                break
            }
        }


        input_string := strings.Replace(string(data), "\r\n", "", 1)
        fmt.Println("input str", input_string)
        cmds := strings.Split(input_string, " ")

        cmd_name, cmd_key := cmds[0], cmds[1]

        if cmd_name == "set" {
            valsize, _ := strconv.Atoi(cmds[4])
            fmt.Println("valsize", valsize)
            valbuf := make([]byte, valsize + 2)
            fmt.Println("expecting val for set")
            c.Read(valbuf)
            //cmd_data := strings.Replace(string(valbuf), "\r\n", "", 1)
            cmd_data := string(valbuf)[0:valsize]
            fmt.Println("Storing ", cmd_key," for ", cmd_data, len(cmd_data))
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
            val, _ := d.Search(cmd_key)
            fmt.Println("retrieved val ", len(val))
            ws := "VALUE " + cmd_key + " 0 " + strconv.Itoa(len(val)) + "\r\n"
            c.Write([]byte(ws))
            c.Write([]byte(val + "\r\n"))
            c.Write([]byte("END\r\n"))

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

