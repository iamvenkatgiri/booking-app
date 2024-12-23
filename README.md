### Booking App

This is a simple booking application written in Go. It allows users to book tickets for a conference.

### How to Run

To initialize the module:
```
go mod init booking-app
```

To run the application:
```
go run main.go
```

### Important Syntax

- Define a list of maps with type `<String, String>`:
```
make([]map[string]string, 0)
```

- Define a custom data type:
```go
type UserData struct {
    firstName   string
    lastName    string
    email       string
    noOfTickets uint
}

### Go Routines and Channels

Go routines are lightweight threads managed by the Go runtime. They allow you to run functions concurrently.

To start a new Go routine, use the `go` keyword:
```go
go myFunction()
```

Channels are used to communicate between Go routines. They can send and receive values of a specific type.

To create a channel:
```go
ch := make(chan int)
```

To send a value to a channel:
```go
ch <- 42
```

To receive a value from a channel:
```go
value := <-ch
```

Here is a simple example demonstrating Go routines and channels:
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    messages := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        messages <- "Hello, World!"
    }()

    msg := <-messages
    fmt.Println(msg)
}
```
In this example, a Go routine sends a message to the `messages` channel after a 2-second delay, and the main function receives and prints the message.