package service

import (
	"random-string-generator-service/internal/app/domain"
	"random-string-generator-service/internal/app/nanoid/config"
)

//go:generate moq -out nanoid_service_mock_test.go . NanoidRepository
type NanoidRepository interface {
	GenerateUniqueID() string
}

type NanoidService struct {
	nanoidRepo NanoidRepository
}

func NewNanoidService(nanoidRepo NanoidRepository) *NanoidService {
	return &NanoidService{nanoidRepo}
}

func (ns NanoidService) RunGenerateRandomString() (domain.NanoidResult, bool) {
	randomResult := ns.nanoidRepo.GenerateUniqueID()
	isValidPrefix := len(randomResult) == len(config.Prefix)+config.LimitCharacter
	return domain.NanoidResult{
		Value: randomResult,
	}, isValidPrefix
}
