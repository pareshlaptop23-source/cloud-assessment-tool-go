package services

import (
	"cloud-assessment-tool/models"
	"testing"
)

func TestGetRecommendations(t *testing.T) {

	vms := []models.VirtualMachine{
		{
			Name:        "VM-1",
			CPUUsage:    5,
			MemoryUsage: 10,
		},
	}

	result := GetRecommendations(vms)

	if len(result) != 1 {
		t.Errorf("Expected 1 recommendation but got %d", len(result))
	}

	if result[0].Recommendation != "Terminate VM" {
		t.Errorf(
			"Expected 'Terminate VM' but got '%s'",
			result[0].Recommendation,
		)
	}
}
func TestResizeRecommendation(t *testing.T) {

	vms := []models.VirtualMachine{
		{
			Name:        "VM-2",
			CPUUsage:    20,
			MemoryUsage: 50,
		},
	}

	result := GetRecommendations(vms)

	if len(result) != 1 {
		t.Errorf("Expected 1 recommendation but got %d", len(result))
	}

	if result[0].Recommendation != "Resize VM" {
		t.Errorf(
			"Expected 'Resize VM' but got '%s'",
			result[0].Recommendation,
		)
	}
}
