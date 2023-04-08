package tickets

import (
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type ServiceImpl struct {
	Repo Repository
}

func (s *ServiceImpl) GetAll() ([]domain.Ticket, error) {
	return s.Repo.GetAll()
	//TODO: Convertir a un DTO
}

func (s *ServiceImpl) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	return s.Repo.GetTicketByDestination(destination)
}

func (s *ServiceImpl) GetTotalTickets(destination string) (int, error) {
	return s.Repo.GetTotalTickets(destination)
}

func (s *ServiceImpl) AverageDestination(destination string) (float64, error) {
	return s.Repo.AverageDestination(destination)
}
