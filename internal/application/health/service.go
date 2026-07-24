package health

import (
	"github.com/barlus-developer/go-simple-http/internal/domain/health"
)

type Service interface {
	Status() health.Status
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Status() health.Status {
	return health.Status{
		Status:  "ok",
		Message: "Hello, World!!!",
	}
}
