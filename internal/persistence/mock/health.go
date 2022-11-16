package mock

type HealthRepo struct{}

func (h HealthRepo) Ping() error {
	return nil
}
