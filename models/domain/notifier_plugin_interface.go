package domain

type INotifierPlugin interface {
	Notify(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier)
	GetId() string
	GetName() string
}
