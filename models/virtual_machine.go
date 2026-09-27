package models

import "gorm.io/gorm"

type VirtualMachine struct {
	gorm.Model
	Name            string `json:"name"`
	CloudProviderID uint

	CPU int

	RAM int

	Disk int

	CPUUsage float64

	MemoryUsage float64

	Status string

	Owner string
}
