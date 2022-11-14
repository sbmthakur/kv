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

func createCommandString(arguments []string) string {
    var commandString strings.Builder

    switch arguments[1] {
    case "set":
        commandString.WriteString(arguments[1])
        commandString.WriteString(" ")
        commandString.WriteString(arguments[2])
        commandString.WriteString(" ")
        commandString.WriteString(arguments[3])
    case "get":
        commandString.WriteString(arguments[1])
        commandString.WriteString(" ")
        commandString.WriteString(arguments[2])
    }
    return commandString.String()
}

func validateCommand(command string, n int) bool {

    command_map := map[string]int{ "set": 4, "get": 3 }

    if command_map[command] != n {
        return false
    }

    return true
}

func main() {

    arguments := os.Args

    command := arguments[1]

    if validateCommand(command, len(arguments)) == false {
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

    commandString := createCommandString(arguments)
    //fmt.Println("Command type")
    //fmt.Println(commandString)

    _, er := con.Write([]byte(commandString + "\r\n"))

    fmt.Println("Key written? ")

    if er != nil {
        fmt.Println("write failure")
    }
}

