package postgres

import (
	"fmt"
	"testing"

	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
)

func TestGetDevicesLocation(t *testing.T) {
	db, err := db.New(cfg)
	if err != nil {
		fmt.Println("Failed to create")
	} else {
		fmt.Println("err :", err)
	}
	cfg := config.Load()

	logger := logger.New(cfg.LogLevel)
	pg := New(db, logger, cfg)
	data, err := pg.GetDeviceLocation(models.GetDeviceLocationRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Result :", data)

}
