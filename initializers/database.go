package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
  var DB *gorm.DB
  // https://github.com/jackc/pgx
 func ConnectToDb() {
	var err error
	dsn := os.Getenv("DSN")
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err!=nil{
	log.Fatal("failed to connect")
  }
}