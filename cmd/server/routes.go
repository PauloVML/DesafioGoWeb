package server

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handlers"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	Bd     *[]domain.Ticket
}

func (r *Router) MapRoutes() {
	r.Engine.Use(gin.Recovery())
	r.Engine.Use(gin.Logger())

	r.SetupTickets()
}

func (r *Router) SetupTickets() {

	repo := tickets.RepositoryImpl{*r.Bd}
	service := tickets.ServiceImpl{&repo}
	handler := handlers.TicketHandler{&service}

	tck := r.Engine.Group("/ticket")

	tck.GET("", handler.GetAll())
	tck.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
	tck.GET("/getAverage/:dest", handler.AverageDestination())

}
