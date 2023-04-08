package tickets

import (
	"errors"
	"fmt"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type RepositoryImpl struct {
	DB []domain.Ticket
}

func (r *RepositoryImpl) GetAll() ([]domain.Ticket, error) {

	if len(r.DB) == 0 {
		return []domain.Ticket{}, ErrEmptyList
	}

	return r.DB, nil
}

func (r *RepositoryImpl) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.DB) == 0 {
		return []domain.Ticket{}, ErrEmptyList
	}

	for _, t := range r.DB {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *RepositoryImpl) GetTotalTickets(destination string) (int, error) {
	var totalTickets int

	for _, ticket := range r.DB {
		if ticket.Country == destination {
			totalTickets++
		}
	}

	if len(r.DB) == 0 || totalTickets == 0 {
		return totalTickets, ErrEmptyList
	}

	return totalTickets, nil

}

func (r *RepositoryImpl) AverageDestination(destination string) (float64, error) {

	total, err := r.GetTotalTickets(destination)

	fmt.Printf("El total de tickets emitidos a %v es de %v \n", destination, total)
	fmt.Println(float64((total / (len(r.DB) + 1)) * 100))

	if err != nil {
		return 0, err
	}

	var promedio float64 = float64(total) / float64(len(r.DB))

	return promedio * 100, nil

}

var (
	ErrEmptyList = errors.New("No hay elementos disponibles")
)
