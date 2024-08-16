package psql

import (
	"fmt"
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func GetDBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

func Connect() (error){

	database_host := os.Getenv("POSTGRES_HOST")
	databale_port := os.Getenv("POSTGRES_PORT")         
	databale_user := os.Getenv("POSTGRES_USER")        
	database_password := os.Getenv("POSTGRES_PASSWORD")
	database_name := os.Getenv("POSTGRES_DB") 
	
	dbConfig := DBConfig{
		Host:     database_host,
		Port:     databale_port,
		User:     databale_user,
		Password: database_password,
		DBName:   database_name,
	}
	dsn := GetDBURL(&dbConfig)

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))
	if err != nil {
		log.Println("Error DB Ping : ", err)
		return err
	}
	log.Println("DB Connection Successful")
	RunDDLScript(DB)
	return nil
}