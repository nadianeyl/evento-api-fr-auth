package usecase

import (
	"github.com/nadianeyl/evento-api-fr-auth/internal/domain"
	"github.com/nadianeyl/evento-api-fr-auth/internal/domain/request"
	"github.com/nadianeyl/evento-api-fr-auth/internal/repository"
)

type TicketTypeUsecase struct {
	repository repository.ITicketTypeRepository
}

func NewTicketTypeUsecase(repository repository.ITicketTypeRepository) ITicketTypeUsecase {
	return &TicketTypeUsecase{
		repository: repository,
	}
}

func (u *TicketTypeUsecase) GetAll() ([]*domain.TicketType, error) {
	return u.repository.GetAll()
}

func (u *TicketTypeUsecase) Add(input *request.TicketTypeRequest) (*domain.TicketType, error) {
	ticketType := &domain.TicketType{
		Name:  input.Name,
		Price: input.Price,
	}

	err := u.repository.Add(ticketType)
	if err != nil {
		return nil, err
	}

	return ticketType, nil
}
