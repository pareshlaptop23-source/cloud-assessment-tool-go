package seed

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
)

func SeedCloudProviders() {

	var count int64
	config.DB.Model(&models.CloudProvider{}).Count(&count)

	if count > 0 {
		return
	}

	providers := []models.CloudProvider{
		{
			Name:   "AWS",
			Type:   "Public",
			Status: "ACTIVE",
		},
		{
			Name:   "Azure",
			Type:   "Public",
			Status: "ACTIVE",
		},
		{
			Name:   "VMware",
			Type:   "Private",
			Status: "ACTIVE",
		},
		{
			Name:   "OpenStack",
			Type:   "Private",
			Status: "ACTIVE",
		},
	}
	config.DB.Create(&providers)
}
