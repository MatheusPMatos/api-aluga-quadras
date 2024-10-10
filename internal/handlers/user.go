package handlers

import "github.com/gin-gonic/gin"

type userHandle struct {
}

// Delete implements UserHandler.
func (u *userHandle) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Edit implements UserHandler.
func (u *userHandle) Edit(c *gin.Context) {
	panic("unimplemented")
}

// GetById implements UserHandler.
func (u *userHandle) GetById(c *gin.Context) {
	panic("unimplemented")
}

func (u *userHandle) Get(c *gin.Context) {
	panic("unimplemented")
}

type UserHandler interface {
	Get(c *gin.Context)
	Edit(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUserHandle() UserHandler {
	return &userHandle{}
}
