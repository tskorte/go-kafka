package main

import (
	"bufio"
	"flag"
	"go-kafka/utils"
	"os"

	"github.com/happierall/l"

	"github.com/segmentio/kafka-go"
)

var (
	key  = flag.String("key", "key-a", "Sets the key of the writer")
	body = flag.String("body", "body-message", "Sets the body of the message")
)

func init() {
	flag.Parse()
}

type test struct {
	name string
}

func main() {
	l.Debug("*******************************************")
	l.Debug("* Initializing kafka writer...please wait *")
	l.Debug("*         cmd/ctrl+c to quit              *")
	l.Debug("*******************************************")
	config := kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}
	sw := utils.StdinWriter{
		Key: *key,
		Kw:  kafka.NewWriter(config),
		R:   bufio.NewReader(os.Stdin),
	}
	initialize(sw)
}

func initialize(w Writer) {
	w.Setup()
}

type Writer interface {
	Setup()
}
