package users

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/request/types"
)

type CreateUserRequest struct {
	Name     string     `json:"name" binding:"required" example:"Name"` // TODO, apply validation
	Email    string     `json:"email" binding:"required" example:"Name@name.com"`
	Age      uint8      `json:"age" binding:"required" example:"50"`
	Birthday types.Date `json:"birthday" binding:"required" example:"2022-06-10"`
}

type UpdateUserRequest struct {
	ID       uint       `json:"int,omitempty" uri:"id" binding:"required"`
	Name     string     `json:"name" binding:"required"` // TODO, apply validation
	Age      uint8      `json:"age" binding:"required"`
	Birthday types.Date `json:"birthday" binding:"required"`
}
