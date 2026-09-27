package models

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model

	UserID uint

	Action string

	Resource string

	ResourceID uint
}
