package postgres

import (
	"fmt"
	"testing"

	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
)

func TestCreateAdmin(t *testing.T) {
	db, err := db.New(cfg)
	if err != nil {
		fmt.Println("Failed to create")
	} else {
		fmt.Println("err :", err)
	}
	cfg := config.Load()

	logger := logger.New(cfg.LogLevel)
	pg := New(db, logger, cfg)
	data, err := pg.CreateAdmin(models.Admin{
		Firstname: "Kamron",
		Lastname:  "Bahtiyorov",
		Login:     "kamron1",
		Password:  "kamron18",
		Phone:     "+998938460418",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
func TestGetAdmins(t *testing.T) {
	db, err := db.New(cfg)
	if err != nil {
		fmt.Println("Failed to create")
	} else {
		fmt.Println("err :", err)
	}
	cfg := config.Load()

	logger := logger.New(cfg.LogLevel)
	pg := New(db, logger, cfg)
	data, err := pg.GetAdmins(models.GetAdminsRequest{
		Id:    "b7b92a2e-e02e-11ee-85a4-c8b29b7e512b",
		Limit: 10,
		Page:  1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Data :", data)
}
