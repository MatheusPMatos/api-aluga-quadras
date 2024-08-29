package handlers

import "github.com/gin-gonic/gin"

type userHandle struct {
}

func (u *userHandle) GetUsers(c *gin.Context) {
	panic("unimplemented")
}

type UserHangle interface {
	GetUsers(c *gin.Context)
}

func NewUserHandle() UserHangle {
	return &userHandle{}
}
