package producer

type Producer interface {
	Init() error
	Produce() error
	Close() error
}
