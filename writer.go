package main

import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
    "github.com/streadway/amqp"
)

type Rabittmq struct {
    Name string
    Data string
}

func main() {

    // connecte with mongo
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("mqtt").C("rabittmq")

    // setting the channel
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    checkeError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    checkeError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
      "hello", // name
      false,   // durable
      false,   // delete when usrused
      false,   // exclusive
      false,   // no-wait
      nil,     // arguments
    )
    checkeError(err, "Failed to declare a queue")



    // add consumer
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    checkeError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            err = c.Insert(&Rabittmq{"massage", string(d.Body)})
            if err != nil {
                log.Fatal(err)
            }
        log.Printf("Received a message: %s", d.Body)
      }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever



    fmt.Println("is ok")
}





func checkeError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", err, msg)
  }
}




//result := Person{}
//err = c.Find(bson.M{"n": ""}).One(&result)
//if err != nil {
//      log.Fatal(err)
//}

//fmt.Println("data: ", result.Phone)
