package tickets

import (
	"context"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cxt = context.Background()

var tickets = []domain.Ticket{
	{
		Id:      "1",
		Name:    "Tait Mc Caughan",
		Email:   "tmc0@scribd.com",
		Country: "Finland",
		Time:    "17:11",
		Price:   785.00,
	},
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

var ticketsByDestination = []domain.Ticket{
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

type stubRepo struct {
	db *DbMock
}

type DbMock struct {
	db  []domain.Ticket
	spy bool
	err error
}

func (r *stubRepo) GetAll() ([]domain.Ticket, error) {
	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}
	return tickets, nil
}

func (r *stubRepo) GetTotalTickets(destination string) (int, error) {
	r.db.spy = true
	var total int
	for _, ticket := range tickets {
		if ticket.Country == destination {
			total++
		}
	}
	return total, nil
}

func (r *stubRepo) AverageDestination(destination string) (float64, error) {
	r.db.spy = true

	var totalDest int
	var total int = len(tickets)
	var avg float64

	for _, ticket := range tickets {
		if ticket.Country == destination {
			totalDest++
		}
	}

	avg = float64(totalDest) / float64(total)

	return avg * 100, nil

}

func (r *stubRepo) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	var tkts []domain.Ticket

	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	for _, t := range r.db.db {
		if t.Country == destination {
			tkts = append(tkts, t)
		}
	}

	return tkts, nil
}

func TestGetTicketByDestination(t *testing.T) {

	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := stubRepo{dbMock}
	service := ServiceImpl{&repo}

	tkts, err := service.GetTotalTickets("China")

	assert.Nil(t, err)
	assert.True(t, dbMock.spy)
	assert.Equal(t, len(ticketsByDestination), tkts)
}

func TestAverageDestination(t *testing.T) {

	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := stubRepo{dbMock}
	service := ServiceImpl{&repo}

	avr, err := service.AverageDestination("China")

	expetedValue := 66

	assert.Nil(t, err)
	assert.Equal(t, expetedValue, int(avr))
	assert.True(t, dbMock.spy)
}

func TestGetTotalTickets(t *testing.T) {
	//Arrange
	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := stubRepo{db: dbMock}
	service := ServiceImpl{&repo}

	//Act
	total, err := service.GetTotalTickets("China")

	//Assert
	expectedValue := 2

	assert.True(t, dbMock.spy)
	assert.Nil(t, err)
	assert.Equal(t, expectedValue, total)
}
