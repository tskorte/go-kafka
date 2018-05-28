package utils

import (
	"context"
	"fmt"
	"io"
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

func NewMessageReader(kr *kafka.Reader, rc chan Message) MessageReader {
	return MessageReader{
		Kr: kr,
		Rc: rc,
	}
}

func (mr MessageReader) Setup() {
	l.Debug("Ready, reading messages!")
	go mr.readMessages()
	mr.handleExit()
}

func (mr *MessageReader) readMessages() {
	mr.handleMessage()
	for {
		m, err := mr.Kr.ReadMessage(context.Background())
		if err == io.EOF {
			break
		} else if err != nil {
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
	mr.waitForExit()
	mr.initiateShutdown()
}

func (mr *MessageReader) waitForExit() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	<-exitSignal
}

func (mr *MessageReader) initiateShutdown() {
	l.Debugf("Initiating shutdown...")
	fine := make(chan bool)
	go mr.closeKafkaReader(fine)
	<-fine
	fmt.Println(' ')
	l.Debug("Kafka reader closed, see ya!")
}

func (mr *MessageReader) closeKafkaReader(fine chan bool) {
	l.Debugf("Closing kafka reader...")
	mr.Kr.Close()
	fine <- true
}
