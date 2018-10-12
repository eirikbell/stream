package streamlib

type KafkaProducerBuilder interface {
	Brokers(brokers []string) KafkaProducerBuilder
	Topic(topic string) KafkaProducerBuilder
}

func NewKafkaProducerBuilder() KafkaProducerBuilder {
	return &kafkaProducerBuilder{}
}

type kafkaProducerBuilder struct {
	brokers []string
	topic   string
}

func (b *kafkaProducerBuilder) Brokers(brokers []string) KafkaProducerBuilder {
	b.brokers = brokers
	return b
}

func (b *kafkaProducerBuilder) Topic(topic string) KafkaProducerBuilder {
	b.topic = topic
	return b
}
