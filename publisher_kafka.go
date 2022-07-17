package main

type PublisherToKafka struct{}

var _ Publisher = &PublisherToKafka{}

func NewPublisherToKafka() *PublisherToKafka {
	return &PublisherToKafka{}
}

func (PublisherToKafka) createEvent(any) error {
	// TODO: logic

	return nil
}

// no error returned. local error handling.
func (p *PublisherToKafka) PublishEvent(data any) {
	p.createEvent(data)

	// TODO: logic
}
