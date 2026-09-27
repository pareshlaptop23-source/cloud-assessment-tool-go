package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
)

const VMCacheKey = "vms:all"

func GetCachedVMs() ([]models.VirtualMachine, error) {

	data, err := config.RedisClient.Get(
		config.RedisCtx,
		VMCacheKey,
	).Result()

	if err != nil {
		fmt.Println("Redis GET error:", err)
		return nil, err
	}

	var vms []models.VirtualMachine

	err = json.Unmarshal([]byte(data), &vms)
	if err != nil {
		fmt.Println("Redis UNMARSHAL error:", err)
		// Log a truncated view of the raw cached value to help debugging.
		raw := data
		if len(raw) > 500 {
			raw = raw[:500] + "... (truncated)"
		}
		fmt.Println("Redis raw cached value:", raw)
		return nil, err
	}

	return vms, nil
}

func SetCachedVMs(vms []models.VirtualMachine) error {

	data, err := json.Marshal(vms)

	if err != nil {
		return err
	}

	// Log a short preview of what's being cached to help debugging.
	preview := string(data)
	if len(preview) > 300 {
		preview = preview[:300] + "... (truncated)"
	}
	fmt.Println("Redis SET key=", VMCacheKey, "len=", len(data), "preview=", preview)

	return config.RedisClient.Set(
		context.Background(),
		VMCacheKey,
		string(data),
		5*time.Minute,
	).Err()
}

func DeleteCachedVMs() error {

	return config.RedisClient.Del(
		config.RedisCtx,
		VMCacheKey,
	).Err()
}
