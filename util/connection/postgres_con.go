package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var DB *gorm.DB

func init() {
	godotenv.Load()
}

func DBInit() (*gorm.DB, error) {

	// dbCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
	dbCon := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	fmt.Println(dbCon)
	var err error
	DB, err = gorm.Open("postgres", dbCon)

	if err != nil {
		fmt.Println(fmt.Sprintf("Failed connected to database %s", dbCon))
		fmt.Println(err)
		return DB, err
	}

	fmt.Println(fmt.Sprintf("Successfully connected to database %s", dbCon))
	maxConLifeTime, err := strconv.Atoi(os.Getenv("DB_CONNECTION_LIFETIME_MINUTE"))
	if err != nil {
		fmt.Println("DB_CONNECTION_LIFETIME_MINUTE Failed")
	}
	maxIdleCount, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION_COUNT"))
	if err != nil {
		fmt.Println("DB_MAX_IDLE_CONNECTION_COUNT Failed")
	}
	maxOpenConCount, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION_COUNT"))
	if err != nil {
		fmt.Println("DB_MAX_OPEN_CONNECTION_COUNT Failed")
	}
	DB.DB().SetConnMaxLifetime(time.Duration(maxConLifeTime) * time.Minute)
	DB.DB().SetMaxIdleConns(maxIdleCount)
	DB.DB().SetMaxOpenConns(maxOpenConCount)
	DB.LogMode(true)

	log, err := zap.NewProduction()
	DB.SetLogger(CustomLogger(log))

	fmt.Println("Connection is created")
	return DB, err
}

func GetConnection() *gorm.DB {
	if DB == nil {
		fmt.Println("No Active Connection Found")
		DB, _ = DBInit()
	}
	return DB
}

func CustomLogger(zap *zap.Logger) *Logger {
	return &Logger{
		zap: zap,
	}
}

type Logger struct {
	zap *zap.Logger
}

func (l *Logger) Print(values ...interface{}) {
	var additionalString = ""
	for _, item := range values {
		if _, ok := item.(string); ok {
			additionalString = additionalString + fmt.Sprintf("\n%v", item)
		}
		if err, ok := item.(*mysql.MySQLError); ok {
			err.Message = err.Message + additionalString
		}
	}
}
