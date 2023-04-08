package handlers

import (
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	Service tickets.Service
}

func (handler *TicketHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := handler.Service.GetTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (handler *TicketHandler) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := handler.Service.AverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, gin.H{"Promedio": avg})
	}
}

func (handler *TicketHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		tickets, err := handler.Service.GetAll()

		if err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, tickets)

	}
}
