package config

import (
	"fmt"
	"os"
	"time"

	env "github.com/caarlos0/env/v6"
	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/gql"
	"github.com/ganiszulfa/concise/backend/internal/models/migrations"
	"github.com/ganiszulfa/concise/backend/pkg/inspect"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Configuration struct {
	DbUser string `env:"CON_DB_USER"`
	DbPw   string `env:"CON_DB_PW"`
	DbPort int    `env:"CON_DB_PORT"`
	DbHost string `env:"CON_DB_HOST"`
	DbName string `env:"CON_DB_NAME"`

	DbMaxIdleConns    int `env:"CON_DB_MAX_IDLE_CON" envDefault:"10"`
	DbMaxOpenConns    int `env:"CON_DB_MAX_OPEN_CON" envDefault:"20"`
	DbConnMaxLifetime int `env:"CON_DB_CON_MAX_LIFETIME" envDefault:"3600"`

	DebugMode bool `env:"CON_DEBUG_MODE" envDefault:"false"`
}

var Config Configuration

func Initialize(env string) {
	trace.Func()

	setConfig(env)

	initDB()

	gql.Initialize()

	cleanSensitiveConfig()
}

func setConfig(envName string) {
	trace.Func()

	if err := godotenv.Load(".env"); err != nil {
		log.Infof("%+v. Will try to get from ENVs directly\n", err)
	}

	if err := env.Parse(&Config); err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

	if Config.DebugMode {
		log.SetLevel(log.TraceLevel)
		inspect.Do(Config)
	}

}

func initDB() {
	trace.Func()

	if Config.DbHost == "" || Config.DbUser == "" {
		panic("no setup for database")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		Config.DbHost,
		Config.DbUser,
		Config.DbPw,
		Config.DbName,
		Config.DbPort,
	)

	newLogger := logger.New(
		log.New(),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	sqlDB, err := app.DB.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(Config.DbMaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(Config.DbMaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(Config.DbConnMaxLifetime) * time.Second)

	err = migrations.Migrate(app.DB)
	if err != nil {
		panic(err)
	}
}

func cleanSensitiveConfig() {
	trace.Func()

	// make sure all username and password has been used
	// to create connection to respective service
	log.Info("Cleaning sensitive data in config..")
	s := "CLEANED"
	Config.DbUser = s
	Config.DbHost = s
	Config.DbPw = s
}
