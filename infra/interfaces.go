package infra

type Authorizer interface {
	IsAuthorized(any) (bool, error)
}

type Publisher interface {
	PublishEvent(data any)
}
