package execute

import (
	"github.com/wuhuZhao/feature_streaming/internal/consumer"
	"github.com/wuhuZhao/feature_streaming/internal/producer"
	"log"
)

type Execute interface {
	Init() error
	Exec(config string) error
	Close() error
}

type DefaultExecute struct {
	producer producer.Producer
	consumer consumer.Consume
}

func NewDefaultExecute(p producer.Producer, c consumer.Consume) *DefaultExecute {
	return &DefaultExecute{
		producer: p,
		consumer: c,
	}
}

func (d *DefaultExecute) Init() error {
	if err := d.producer.Init(); err != nil {
		return err
	}
	return nil
}

func (d *DefaultExecute) Exec(config string) error {
	if err := d.consumer.Init(config); err != nil {
		return err
	}
	go func() {
		err := d.consumer.Consume()
		if err != nil {
			log.Fatalf("consumer error: %v\n", err)
		}
	}()
	return nil
}

func (d *DefaultExecute) Close() error {
	if err := d.producer.Close(); err != nil {
		return err
	}
	if err := d.consumer.Close(); err != nil {
		return err
	}
	return nil
}
