package main

import (
	"go-kafka/utils"

	"github.com/happierall/l"
	"github.com/segmentio/kafka-go"
)

func main() {
	l.Debug("*******************************************")
	l.Debug("* Initializing kafka reader...please wait *")
	l.Debug("*         cmd/ctrl+c to quit              *")
	l.Debug("*******************************************")
	topic := "my-topic"
	partition := 0
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
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
