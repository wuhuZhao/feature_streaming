package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

var _ Producer = (*KafkaProducer)(nil)

type KafkaProducer struct {
	topic     string
	partition int
	host      string
	conn      *kafka.Conn
	batch     *kafka.Batch
	queue     chan []byte
}

func NewKafkaProducer(host, topic string, partition int, c chan []byte) *KafkaProducer {
	producer := &KafkaProducer{
		topic:     topic,
		partition: partition,
		host:      host,
		queue:     c,
	}
	return producer
}

func (k *KafkaProducer) Init() error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", k.host, k.topic, k.partition)
	if err != nil {
		return err
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6)
	k.conn = conn
	k.batch = batch
	return nil
}

func (k *KafkaProducer) Produce() error {
	b := make([]byte, 10e3)
	for {
		n, err := k.batch.Read(b)
		if err != nil {
			return err
		}
		k.queue <- b[:n]
	}
}

func (k *KafkaProducer) Close() error {
	if err := k.batch.Close(); err != nil {
		return err
	}
	if err := k.conn.Close(); err != nil {
		return err
	}
	return nil
}
