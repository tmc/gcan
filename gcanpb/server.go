package gcanpb

type Server interface {
	ProducerServer
	ConsumerServer
}
