package services

import "cloud-assessment-tool/models"

type VMwareService struct{}

func (v VMwareService) ListVMs() ([]models.VirtualMachine, error) {

	return []models.VirtualMachine{
		{
			Name:     "VMWARE-VM-01",
			CPU:      16,
			RAM:      32,
			Disk:     500,
			CPUUsage: 45,

			MemoryUsage: 60,
			Status:      "RUNNING",
			Owner:       "Infrastructure Team",
		},
	}, nil

}
func (a VMwareService) StartVM(id uint) error {

	return nil
}

func (a VMwareService) StopVM(id uint) error {

	return nil
}
