package persistence

type HealthRepository interface {
	Ping() error
}
