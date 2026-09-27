package services

import "cloud-assessment-tool/models"

type AWSService struct{}

func (a AWSService) ListVMs() ([]models.VirtualMachine, error) {

	return []models.VirtualMachine{

		{
			Name:        "AWS-VM-01",
			CPU:         4,
			RAM:         8,
			Disk:        100,
			CPUUsage:    45,
			MemoryUsage: 60,
			Status:      "RUNNING",
			Owner:       "Paresh",
		},
	}, nil

}
func (a AWSService) StartVM(id uint) error {

	return nil
}

func (a AWSService) StopVM(id uint) error {

	return nil
}
