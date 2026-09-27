package services

import "cloud-assessment-tool/models"

type OpenStackService struct{}

func (o OpenStackService) ListVMs() ([]models.VirtualMachine, error) {

	return []models.VirtualMachine{
		{
			Name:        "OPENSTACK-VM-01",
			CPU:         2,
			RAM:         4,
			Disk:        50,
			CPUUsage:    20,
			MemoryUsage: 30,
			Status:      "RUNNING",
			Owner:       "Dev Team",
		},
	}, nil
}

func (o OpenStackService) StartVM(id uint) error {
	return nil
}

func (o OpenStackService) StopVM(id uint) error {
	return nil
}
