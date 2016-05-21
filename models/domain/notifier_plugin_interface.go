package domain

type INotifierPlugin interface {
	Notify(healthcheck Healthcheck)
	GetId() string
	GetName() string
}