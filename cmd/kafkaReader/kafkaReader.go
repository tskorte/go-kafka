package main

import (
	"flag"
	"go-kafka/utils"

	"github.com/happierall/l"
	"github.com/segmentio/kafka-go"
)

var (
	address = flag.String("address", "localhost:9092", "Sets the kafka address")
	key     = flag.String("key", "key-a", "Sets the key of the reader")
	topic   = flag.String("topic", "my-topic", "Sets the topic of the reader")
)

func init() {
	flag.Parse()
}

func main() {
	l.Debug("*******************************************")
	l.Debug("* Initializing kafka reader...please wait *")
	l.Debug("*         cmd/ctrl+c to quit              *")
	l.Debug("*******************************************")
	partition := 0
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{*address},
		Topic:     *topic,
		Partition: partition,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(-1)
	c := make(chan utils.Message)
	handler := utils.MessageReader{
		Kr: r,
		Rc: c,
	}
	initialize(handler)
}

func initialize(r Reader) {
	r.Setup()
}

type Printer interface {
}

type Reader interface {
	Setup()
}
