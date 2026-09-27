package handlers

import (
	"cloud-assessment-tool/models"
	"cloud-assessment-tool/repository"
	"cloud-assessment-tool/services"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// SyncCloud godoc
// @Summary Synchronize cloud resources
// @Description Synchronizes cloud resources with the Cloud Assessment Tool
// @Tags Cloud
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cloud/sync [post]
func SyncCloud(c *gin.Context) {
	var wg sync.WaitGroup
	vmChannels := make(chan []models.VirtualMachine)

	aws := services.AWSService{}
	azure := services.AzureService{}
	vmware := services.VMwareService{}
	openstack := services.OpenStackService{}

	clouds := []services.CloudService{

		aws,

		azure,

		vmware,

		openstack,
	}
	var allVMs []models.VirtualMachine
	wg.Add(len(clouds))

	for _, cloud := range clouds {

		go func(cloud services.CloudService) {
			defer wg.Done()
			vms, err := cloud.ListVMs()

			if err != nil {
				return
			}
			vmChannels <- vms
		}(cloud)
	}
	go func() {
		wg.Wait()
		close(vmChannels)
	}()
	for vms := range vmChannels {
		allVMs = append(allVMs, vms...)
	}
	for _, vm := range allVMs {

		existingVM, err := repository.GetVMByName(vm.Name)

		if err != nil {

			// VM doesn't exist
			repository.CreateVM(&vm)

		} else {

			// VM already exists
			existingVM.CPU = vm.CPU
			existingVM.RAM = vm.RAM
			existingVM.Disk = vm.Disk
			existingVM.CPUUsage = vm.CPUUsage
			existingVM.MemoryUsage = vm.MemoryUsage
			existingVM.Status = vm.Status
			existingVM.Owner = vm.Owner

			repository.UpdateVM(&existingVM)
		}
	}
	c.JSON(

		http.StatusOK,

		gin.H{

			"message": "Cloud Sync Successful",

			"data": allVMs,
		},
	)
}
