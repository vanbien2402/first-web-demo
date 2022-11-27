package models

import "github.com/vanbien2402/first-web-demo/internal/pkg/rds"

type Category struct {
	rds.Model
	Name          string `json:"name"`
	Remark        string `json:"remark"`
	ExpiredInDays int    `json:"expiredInDays"`
}
