package streamlib

type StreamBuilder interface {
	Consumer(consumerBuilder KafkaConsumerBuilder) StreamBuilder
	Producer(producerBuilder KafkaProducerBuilder) StreamBuilder
	Concurrency(concurrency int) StreamBuilder
	MessageFunc(messageFunc func(messageBody []byte) ([]byte, error)) StreamBuilder
	Build() Stream
}

func NewStreamBuilder() StreamBuilder {
	return &streamBuilder{}
}

type streamBuilder struct {
}

func (b *streamBuilder) Consumer(consumerBuilder KafkaConsumerBuilder) StreamBuilder {
	return b
}

func (b *streamBuilder) Producer(producerBuilder KafkaProducerBuilder) StreamBuilder {
	return b
}

func (b *streamBuilder) Concurrency(concurrency int) StreamBuilder {
	return b
}

func (b *streamBuilder) MessageFunc(messageFunc func(messageBody []byte) ([]byte, error)) StreamBuilder {
	return b
}

func (b *streamBuilder) Build() Stream {
	return &stream{}
}
