package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host           string
	Port           string
	Admin_User     string
	Admin_Password string
	DBName         string
	SSLMode        string
	App_User       string
	App_Password   string
}

var (
	DB    *gorm.DB // полное подключение к DB для миграций
	AppDB *gorm.DB // ограниченное подключение к DB
)

func CreateAppUser(db *gorm.DB, cfg Config) {
	err := db.Exec(fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", cfg.App_User, cfg.App_Password)).Error
	if err != nil {
		log.Println(err)
		return
	}
	err = db.Exec(fmt.Sprintf("GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO %s", cfg.App_User)).Error
	if err != nil {
		log.Println(err)
		return
	}
	err = db.Exec(fmt.Sprintf("GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO %s", cfg.App_User)).Error
	if err != nil {
		log.Println(err)
		return
	}
}

func InitDB(cfg Config) {
	adminDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Admin_User, cfg.Admin_Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	adminDB, err := gorm.Open(postgres.Open(adminDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := adminDB.AutoMigrate(&User{}, &Album{}, &Song{}); err != nil {
		panic(err)
	}
	DB = adminDB
	log.Println("Migrated database")

	CreateAppUser(DB, cfg)

	appDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.App_User, cfg.App_Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	appDB, err := gorm.Open(postgres.Open(appDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	AppDB = appDB

	log.Println("Database initialized with restricted user")
}
