package users

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/request/types"
)

type CreateUserRequest struct {
	Name     string     `json:"name" binding:"required"` // TODO, apply validation
	Email    string     `json:"email" binding:"required"`
	Age      uint8      `json:"age" binding:"required"`
	Birthday types.Date `json:"birthday" binding:"required"`
}

type UpdateUserRequest struct {
	ID       uint       `json:"int,omitempty" uri:"id" binding:"required"`
	Name     string     `json:"name" binding:"required"` // TODO, apply validation
	Age      uint8      `json:"age" binding:"required"`
	Birthday types.Date `json:"birthday" binding:"required"`
}
