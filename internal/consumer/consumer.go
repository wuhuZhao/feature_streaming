package consumer

import (
	"github.com/bytedance/gopkg/util/gopool"
	fe "github.com/wuhuZhao/feature_extractor/pkg"
	"sync"
	"unsafe"
)

type Consume interface {
	Init(config string) error
	Consume() error
	Close() error
}

type DefaultConsumer struct {
	queue     chan []byte
	config    []string
	extractor []*fe.Handler
	mu        *sync.Mutex
}

func NewDefaultConsumer(c chan []byte) *DefaultConsumer {
	return &DefaultConsumer{queue: c, mu: &sync.Mutex{}}
}

func (d *DefaultConsumer) Init(config string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	handler := fe.NewHandler()
	handler.Init(config)
	d.extractor = append(d.extractor, handler)
	d.config = append(d.config, config)
	return nil
}

func (d *DefaultConsumer) Consume() error {
	ll := len(d.extractor)
	for ll == len(d.extractor) {
		data := <-d.queue
		for i := 0; i < len(d.extractor); i++ {
			handler := d.extractor[i]
			gopool.Go(func() {
				handler.Handle(*(*string)(unsafe.Pointer(&data)))
			})
		}
		ll = len(d.extractor)
	}
	return nil
}

func (d *DefaultConsumer) Close() error {
	return nil
}
