package streamlib

type KafkaProducerBuilder interface {
	Brokers(brokers string) KafkaProducerBuilder
	Topic(topic string) KafkaProducerBuilder
	ConsumerGroup(consumerGroup string) KafkaProducerBuilder
}

func NewKafkaProducerBuilder() KafkaProducerBuilder {
	return &kafkaProducerBuilder{}
}

type kafkaProducerBuilder struct {
}

func (b *kafkaProducerBuilder) Brokers(brokers string) KafkaProducerBuilder {
	return b
}

func (b *kafkaProducerBuilder) Topic(topic string) KafkaProducerBuilder {
	return b
}

func (b *kafkaProducerBuilder) ConsumerGroup(consumerGroup string) KafkaProducerBuilder {
	return b
}
