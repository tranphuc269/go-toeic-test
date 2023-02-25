package repositories

import (
	"english_exam_go/utils/app_logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

var conn *gorm.DB

func GetConn() *gorm.DB {
	return conn
}
func OpenDatabase() {
	HOST := "host=" + os.Getenv("DB_HOST")
	USER := " user=" + os.Getenv("DB_USER")
	PASS := " password=" + os.Getenv("DB_PASSWORD")
	DBNAME := " dbname=" + os.Getenv("DB_NAME")
	PORT := " port=" + os.Getenv("DB_PORT")
	SSLMODE := " sslmode=disable"
	TIMEZONE := " timezone=Asia/Tokyo"

	dsn := HOST + USER + PASS + DBNAME + PORT + SSLMODE + TIMEZONE

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 app_logger.Default,
	})

	if err != nil {
		app_logger.Logger.Debug("TransactionImpl Commit")
		app_logger.Logger.Panic("Can not connect DB. error:" + err.Error())
	}

	connPoolSetting(db)

	conn = db
}

func connPoolSetting(db *gorm.DB) *gorm.DB {
	connPoolDB, err := db.DB()
	if err != nil {
		app_logger.Logger.Panic(err.Error())
	}

	// 確立時からのコネクションを保持する最大時間
	maxLifeTime, _ := strconv.Atoi(os.Getenv("MAX_LIFE_TIME"))
	connPoolDB.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Minute)

	// idle状態のコネクションを保持する最大時間
	maxIdleTime, _ := strconv.Atoi(os.Getenv("MAX_IDLE_TIME"))
	connPoolDB.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Minute)

	// idle状態のコネクションの最大数
	maxIdleConns, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	connPoolDB.SetMaxIdleConns(maxIdleConns)

	// プール可能なコネクションの最大数
	maxOpenConns, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	connPoolDB.SetMaxOpenConns(maxOpenConns)

	return db
}
