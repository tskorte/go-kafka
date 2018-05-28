# go-kafka

## Requirements
A running Kafka instance (https://kafka.apache.org/quickstart)

## Installation
```
Run go install go-kafka/cmd/kafkaReader
Run go install go-kafka/cmd/kafkaWriter
Run $GOPATH/bin/kafkaReader to start the consumer
Run $GOPATH/bin/kafkaWriter to produce a message
```

## Configuration
The reader accepts `address`, `key` and `topic`
``` 
./kafkaReader -address="localhost:9092" -key="myKey" -topic="my-topic"
```

The writer accepts `address`, `key` and `topic`
``` 
./kafkaWriter -address="localhost:9092" -key="myKey" -topic="my-topic"
```

Default values are:
```
address = "localhost:9092"
key = "key-a"
topic = "my-topic"
```
## About
I used [kafka-go](https://github.com/segmentio/kafka-go) from [Segment](www.segment.io) for this example repo.