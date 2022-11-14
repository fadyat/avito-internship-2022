package validators

import (
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

func ValidateServiceRequest(service dto.OuterService) error {
	if service.Name == "" {
		return errors.New("service name is empty")
	}

	return nil
}

func ValidateServiceID(id uint64) error {
	if id == 0 {
		return errors.New("service id is empty")
	}

	return nil
}
