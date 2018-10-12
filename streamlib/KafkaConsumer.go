package streamlib

type KafkaConsumerBuilder interface {
	Brokers(brokers []string) KafkaConsumerBuilder
	Topic(topic string) KafkaConsumerBuilder
	ConsumerGroup(consumerGroup string) KafkaConsumerBuilder
}

func NewKafkaConsumerBuilder() KafkaConsumerBuilder {
	return &kafkaConsumerBuilder{}
}

type kafkaConsumerBuilder struct {
	brokers       []string
	topic         string
	consumerGroup string
}

func (b *kafkaConsumerBuilder) Brokers(brokers []string) KafkaConsumerBuilder {
	b.brokers = brokers
	return b
}

func (b *kafkaConsumerBuilder) Topic(topic string) KafkaConsumerBuilder {
	b.topic = topic
	return b
}

func (b *kafkaConsumerBuilder) ConsumerGroup(consumerGroup string) KafkaConsumerBuilder {
	b.consumerGroup = consumerGroup
	return b
}
