package main

import (
	"github.com/eirikbell/stream/streamlib"
)

func main() {
	consumerBuilder := streamlib.NewKafkaConsumerBuilder().
		Brokers("localhost:9092").
		Topic("test").
		ConsumerGroup("testconsumer")
	producerBuilder := streamlib.NewKafkaProducerBuilder().
		Brokers("localhost:9092").
		Topic("test2")
	streamBuilder := streamlib.NewStreamBuilder().
		Consumer(consumerBuilder).
		Producer(producerBuilder).
		Concurrency(100).
		MessageFunc(testMessageFunc)

	stream := streamBuilder.Build()

	_ = stream.Start()
}

func testMessageFunc(messageBody []byte) ([]byte, error) {
	return messageBody, nil
}
