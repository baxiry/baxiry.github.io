package main

import (
    "fmt"
    "log"
    "github.com/streadway/amqp"
)

func checkErr(msg string, err error) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}
func main() {

    connect , err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    checkErr("Failed to connect to RabittMQ", err)
    defer connect.Close()


    ch, err := connect.Channel()
    checkErr("Failed to opon a channel", err)
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello", // name
        false, // durable
        false, // delet when unused
        false, // exclusive
        false, // no wait
        nil, // arguments
    )

    body := "Hello World. This is a test massage"

    for i := 0; i< 100; i++ {
        err = ch.Publish(
            "", // exchange
            q.Name, // routine key
            false, // mandatory
            false, // immediate
            amqp.Publishing {
                ContentType: "text/plain",
                Body: []byte(string(i)+" : "+body),
            },
        )
    }
    checkErr("Failed to publish a message", err)

    fmt.Println("is ok")
}
