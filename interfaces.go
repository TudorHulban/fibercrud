package main

type Authorizer interface {
	IsAuthorized(any) (bool, error)
}

type Publisher interface {
	CreateEvent(any) error
	PublishEvent(any) error
}
