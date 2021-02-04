package main
 
import (
    "fmt"
    "log"
    "net"
)
 
func main() {
    conn, err := net.Dial("tcp", "localhost:3001")
    if err != nil {
        log.Fatalln(err)
    }
 
    _, err = conn.Write([]byte("Hello Server balakrishna!"))
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Message sent: Hello Server!")
 
    _, err = conn.Write([]byte("How are you balakrishna?"))
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Message sent: How are you?")
 
    for {
        buffer := make([]byte, 1400)
        dataSize, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("connection closed")
            return
        }
 
        data := buffer[:dataSize]
        fmt.Println("received message: ", string(data))
    }
}
 