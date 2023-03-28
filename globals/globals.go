package globals

import (
	model "DevOps/model/gorm"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Secret = []byte(os.Getenv("SESSION_SECRET"))

const ENV_KEY = "GO_ENV"

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestRequestId int = -1

var db *sqlx.DB

var gormDb *gorm.DB

var logger *zap.SugaredLogger

func GetDatabase() *sqlx.DB {
	return db
}

func GetGormDatabase() *gorm.DB {
	return gormDb
}
func SetDatabase(database *gorm.DB) {

	gormDb = database
	gormDb.AutoMigrate(&model.User{})
	gormDb.AutoMigrate(&model.Message{})
	gormDb.AutoMigrate(&model.Following{})
}

func GetDatabasePath() gorm.Dialector {
	if os.Getenv(ENV_KEY) == "production" {
		connectionString := fmt.Sprintf("postgresql://%s:%s@postgres/db",
			os.Getenv("POSTGRES_USERNAME"),
			os.Getenv("POSTGRES_PASSWORD"))
		return postgres.Open(connectionString)
	}
	return sqlite.Open("itu-minitwit.db")
}

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}

func SetupLogger() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	_logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	sugar := _logger.Sugar()
	logger = sugar
	sugar.Info("Sugared logger initialized.")
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
