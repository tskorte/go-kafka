package utils

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/happierall/l"
	kafka "github.com/segmentio/kafka-go"
)

type StdinWriter struct {
	Key string
	Kw  *kafka.Writer
	R   *bufio.Reader
}

func (w StdinWriter) Setup() {
	go w.setupMessageHandling(w.R)
	w.setupExitHandling()
}

func (wh *StdinWriter) setupMessageHandling(reader *bufio.Reader) {
	l.Debug(l.Colorize("Waiting for input...", l.LightBlue))
	read, _ := reader.ReadString('\n')
	m := kafka.Message{
		Key:   []byte(wh.Key),
		Value: []byte(strings.TrimSpace(read)),
	}
	wh.writeMessages(wh.Kw, m)
	wh.setupMessageHandling(reader)
}

func (wh *StdinWriter) setupExitHandling() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	<-exitSignal
	l.Debug("Received interrupt - shutting down kafka writer")
	wh.Kw.Close()
	l.Debug("Kafka writer closed, see ya!")
}

func (wh *StdinWriter) writeMessages(w *kafka.Writer, msg kafka.Message) {
	l.Debug(l.Colorize("Sending message...", l.Green))
	err := w.WriteMessages(
		context.Background(),
		msg,
	)
	if err != nil {
		log.Fatalf("Error while writing messages: %s", err)
	}
}
