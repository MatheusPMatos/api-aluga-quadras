package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/gin-gonic/gin"
)

type schedule struct {
	serv service.Schedule
}

type data struct {
	Date time.Time `json:"date"`
}

// GetByProductAndDate implements Schedule.
func (s *schedule) GetByProductAndDate(c *gin.Context) {
	var data data

	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}
	sched, err := s.serv.GetByProductWeekDay(uint(id), data.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("erro: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, sched)

}

type Schedule interface {
	GetByProductAndDate(c *gin.Context)
}

func NewScheduleHandler(sv service.Schedule) Schedule {
	return &schedule{serv: sv}
}
