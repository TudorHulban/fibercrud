package pubk

import "github.com/TudorHulban/fibercrud/infra"

type PublisherToKafka struct{}

var _ infra.Publisher = &PublisherToKafka{}

func NewPublisherToKafka() *PublisherToKafka {
	return &PublisherToKafka{}
}

func (PublisherToKafka) createEvent(any) error {
	// TODO: logic
	return nil
}

// no error returned. local error handling.
func (p *PublisherToKafka) PublishEvent(data any) {
	_ = p.createEvent(data)
	// TODO: logic
}
