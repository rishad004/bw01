package db

import (
	"fmt"

	"github.com/rishad004/bw01/microservice-01/pkg/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PsqlConn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(viper.GetString("PSQL_URL")))

	if err != nil {
		return nil, err
	} else {
		fmt.Println("Connect to Psql")
	}

	db.AutoMigrate(&domain.User{})

	return db, nil
}
