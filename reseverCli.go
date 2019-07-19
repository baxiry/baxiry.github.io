package main
import (
  "fmt"
  "log"

  "github.com/streadway/amqp"
)


func main() {

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    checkError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    checkError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello", // name
      false,   // durable
      false,   // delete when usused
      false,   // exclusive
      false,   // no-wait
      nil,     // arguments
    )
    checkError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
      q.Name, // queue
      "ahmed",     // consumer
      true,   // auto-ack
      false,  // exclusive
      false,  // no-local
      false,  // no-wait
      nil,    // args
    )
    checkError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func() {
      for d := range msgs {
        log.Printf("Received a message: %s", d.Body)
      }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever
    fmt.Println("is ok")
}





func checkError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}
