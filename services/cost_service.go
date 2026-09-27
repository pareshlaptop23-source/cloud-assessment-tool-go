package services

import "cloud-assessment-tool/models"

type Recommendation struct {
	Name           string `json:"name"`
	Recommendation string `json:"recommendation"`
}

func GetRecommendations(vms []models.VirtualMachine) []Recommendation {

	var recommendations []Recommendation

	for _, vm := range vms {

		if vm.CPUUsage < 10 && vm.MemoryUsage < 20 {

			recommendations = append(
				recommendations,
				Recommendation{
					Name:           vm.Name,
					Recommendation: "Terminate VM",
				},
			)

		} else if vm.CPUUsage < 30 {

			recommendations = append(
				recommendations,
				Recommendation{
					Name:           vm.Name,
					Recommendation: "Resize VM",
				},
			)
		}
	}

	return recommendations
}
