package models

import (
	"fmt"
	"github.com/filedrive-team/filplus-info/settings/settingtypes"
	"github.com/filedrive-team/filplus-info/types"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

var autoMigrateModels = make([]interface{}, 0)

type Model struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt types.UnixTime  `json:"created_at"`
	UpdatedAt types.UnixTime  `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Setup initializes the database instance
func Setup(conf *settingtypes.AppConfig) {
	var err error
	var dialector gorm.Dialector
	switch conf.Database.Type {
	case "postgres":
		dialector = postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Name))
	case "mysql":
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Name))
	}

	newLogger := gormlogger.New(
		log.New(os.Stdout, "\n", log.LstdFlags), // io writer
		gormlogger.Config{
			SlowThreshold: 5 * time.Second,   // slow SQL threshold
			LogLevel:      gormlogger.Silent, // Log level
			Colorful:      false,             // disable color printing
		},
	)
	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logger.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("db.DB() err: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	autoMigrate()
	logger.Info("### model.Setup finished! ###")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer func() {
		logger.Info("close db")
	}()
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func autoMigrate() {
	if len(autoMigrateModels) > 0 {
		logger.Info("auto migrate start...")
		db.AutoMigrate(autoMigrateModels...)
		logger.Info("auto migrate finish")
	}
}
