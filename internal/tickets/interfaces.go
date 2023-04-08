package tickets

import (
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (float64, error)
}

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (float64, error)
}
