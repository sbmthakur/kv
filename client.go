package main

import (
    "fmt"
    "net"
    "os"
    "strings"
)

func read(con net.Conn) {
    buf := make([]byte, 20)
    con.Read(buf)
    fmt.Println(string(buf))
    con.Close()
}

func validate_command(command string, n int) bool {

    command_map := map[string]int{ "set": 4, "get": 3 }

    if command_map[command] != n {
        return false
    }

    return true
}

func main() {

    arguments := os.Args

    command := arguments[1]

    if validate_command(command, len(arguments)) == false {
        fmt.Println("Invalid number of arguments")
        return
    }

    if strings.Contains(arguments[1], "\r\n") {
        fmt.Println(arguments[1])
    }

    con, e := net.Dial("tcp", ":9002")

    if e != nil {
        fmt.Println("connection failure")
    }

    defer read(con)
    // Concat string...
    _, er := con.Write([]byte("set key val\r\n"))

    fmt.Println("Key written? ")

    if er != nil {
        fmt.Println("write failure")
    }

    /*
    for {
        read(con)
    }
    */
}

