package repository

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
)

func CreateAuditLog(log *models.AuditLog) error {

	return config.DB.Create(log).Error
}

func GetAuditLogs() ([]models.AuditLog, error) {

	var logs []models.AuditLog

	err := config.DB.
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}
