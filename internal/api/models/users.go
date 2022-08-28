package models

import (
	"github.com/vanbien2402/first-web-demo/internal/pkg/rds"
)

type User struct {
	rds.Model
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
