package repository

import (
	"log"
	"random-string-generator-service/internal/app/domain"

	"github.com/aidarkhanov/nanoid"
)

type CmdNanoidRepository struct {
	config domain.NanoidConfig
}

func NewCmdNanoidRepository(limit int, prefix string) *CmdNanoidRepository {
	return &CmdNanoidRepository{
		config: domain.NanoidConfig{
			Prefix:         prefix,
			LimitCharacter: limit,
		},
	}
}

func (cnr CmdNanoidRepository) GenerateUniqueID() string {
	defaultAlphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id, err := nanoid.Generate(defaultAlphabet, cnr.config.LimitCharacter)
	if err != nil {
		log.Println(err)
		return ""
	}
	return cnr.config.Prefix + id
}
