package utils

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/happierall/l"
	kafka "github.com/segmentio/kafka-go"
)

type Message struct {
	msg    string
	offset int64
}

type MessageReader struct {
	Kr *kafka.Reader
	Rc chan Message
}

func (mr MessageReader) Setup() {
	l.Debug("Ready, reading messages!")
	mr.handleMessage()
	for {
		m, err := mr.Kr.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading messages: %s", err)
			break
		}
		receivedMSG := Message{
			msg:    string(m.Value),
			offset: m.Offset,
		}
		mr.Rc <- receivedMSG
	}
	mr.Kr.Close()
}

func (mh *MessageReader) handleMessage() {
	go func() {
		for {
			msg := <-mh.Rc
			msgString := l.Colorize(fmt.Sprintf("%d - %s", msg.offset, msg.msg), l.LightBlue)
			fmt.Println(msgString)
		}
	}()
}

func (mr *MessageReader) handleExit() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	<-exitSignal
	l.Debug("Received interrupt - shutting down kafka reader")
	mr.Kr.Close()
	l.Debug("Kafka reader closed, see ya!")
}
