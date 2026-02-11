package url

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection(DSN string) (*gorm.DB, error) {

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{Logger: gormLogger})

	if err != nil {
		return nil, err
	}

	return db, nil
}
