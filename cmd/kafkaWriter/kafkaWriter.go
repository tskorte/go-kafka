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
	address = flag.String("address", "localhost:9092", "Sets the kafka address")
	topic   = flag.String("topic", "my-topic", "Sets the topic of the reader")
	key     = flag.String("key", "key-a", "Sets the key of the writer")
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
		Brokers:  []string{*address},
		Topic:    *topic,
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
