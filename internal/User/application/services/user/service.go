package user

import (
	"UserCrud/internal/User/application/ports/user"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type Service struct {
	trm  trm.Manager
	repo ports.UserRepository
}

func NewService(trm trm.Manager, repo ports.UserRepository) ports.UserService {
	return &Service{
		trm:  trm,
		repo: repo}
}
