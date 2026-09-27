package models

import "gorm.io/gorm"

type CloudProvider struct {
	gorm.Model
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}
