package streamlib

type KafkaConsumerBuilder interface {
	Brokers(brokers string) KafkaConsumerBuilder
	Topic(topic string) KafkaConsumerBuilder
	ConsumerGroup(consumerGroup string) KafkaConsumerBuilder
}

func NewKafkaConsumerBuilder() KafkaConsumerBuilder {
	return &kafkaConsumerBuilder{}
}

type kafkaConsumerBuilder struct {
}

func (b *kafkaConsumerBuilder) Brokers(brokers string) KafkaConsumerBuilder {
	return b
}

func (b *kafkaConsumerBuilder) Topic(topic string) KafkaConsumerBuilder {
	return b
}

func (b *kafkaConsumerBuilder) ConsumerGroup(consumerGroup string) KafkaConsumerBuilder {
	return b
}
