package services

import "cloud-assessment-tool/models"

type AzureService struct{}

func (a AzureService) ListVMs() ([]models.VirtualMachine, error) {

	return []models.VirtualMachine{

		{
			Name:        "AZURE-VM-01",
			CPU:         8,
			RAM:         16,
			Disk:        200,
			CPUUsage:    30,
			MemoryUsage: 50,
			Status:      "RUNNING",
			Owner:       "John",
		},
	}, nil

}
func (a AzureService) StartVM(id uint) error {

	return nil
}

func (a AzureService) StopVM(id uint) error {

	return nil
}
