package repository

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
)

func CreateVM(vm *models.VirtualMachine) error {

	return config.DB.Create(vm).Error
}

func GetAllVMs() ([]models.VirtualMachine, error) {

	var vms []models.VirtualMachine

	err := config.DB.Find(&vms).Error

	return vms, err
}

func UpdateVM(vm *models.VirtualMachine) error {

	return config.DB.Save(vm).Error
}

func GetVMByID(id uint) (models.VirtualMachine, error) {

	var vm models.VirtualMachine

	err := config.DB.First(&vm, id).Error

	return vm, err
}

func StopVM(id uint) error {

	return config.DB.
		Model(&models.VirtualMachine{}).
		Where("id = ?", id).
		Update("status", "STOPPED").Error
}
func StartVM(id uint) error {

	return config.DB.
		Model(&models.VirtualMachine{}).
		Where("id = ?", id).
		Update("status", "RUNNING").Error
}
func GetVMByName(name string) (models.VirtualMachine, error) {

	var vm models.VirtualMachine

	err := config.DB.
		Where("name = ?", name).
		First(&vm).Error

	return vm, err
}
