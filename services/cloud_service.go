package services

import "cloud-assessment-tool/models"

type CloudService interface {
	ListVMs() ([]models.VirtualMachine, error)

	StartVM(id uint) error

	StopVM(id uint) error
}
