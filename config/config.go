package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ganiszulfa/concise/config/app"
	"github.com/ganiszulfa/concise/internal/gql"
	"github.com/ganiszulfa/concise/internal/models"
	"github.com/ganiszulfa/concise/pkg/inspect"
	"github.com/ganiszulfa/concise/pkg/trace"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Configuration struct {
	Service struct {
		Name      string `yaml:"name"`
		LogLevel  int    `yaml:"loglevel"`
		DebugMode bool   `yaml:"debug"`
	} `yaml:"service"`
	HttpServer struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"httpserver"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"db"`
		LogLevel int    `yaml:"loglevel"`

		MaxIdleConns    int `yaml:"maxidlecon"`
		MaxOpenConns    int `yaml:"maxopencon"`
		ConnMaxLifetime int `yaml:"conmaxlifetime"`
	} `yaml:"database"`
	Jwt struct {
		SecretKey string `yaml:"secretkey"`
		TokenAge  int    `yaml:"age"`
	} `yaml:"jwt"`
}

var Config Configuration

func Initialize(env string) {
	trace.Func()

	setConfig(env)

	initDB()

	gql.Initialize()

	cleanSensitiveConfig()
}

func setConfig(env string) {
	trace.Func()

	confFilePath := fmt.Sprintf("config/%s.yml", env)
	f, err := os.Open(confFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.Level(Config.Service.LogLevel))

	if Config.Service.DebugMode {
		inspect.Do(Config)
	}

}

func initDB() {
	trace.Func()

	if Config.Database.Host == "" || Config.Database.Username == "" {
		panic("no setup for database")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		Config.Database.Host,
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.DBName,
		Config.Database.Port)

	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := app.DB.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(Config.Database.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(Config.Database.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(Config.Database.ConnMaxLifetime) * time.Second)

	models.AutoMigrateAllTables(app.DB)
}

func cleanSensitiveConfig() {
	trace.Func()

	// make sure all username and password has been used
	// to create connection to respective service
	log.Info("Cleaning sensitive data in config..")
	s := "CLEANED"
	Config.Database.Username = s
	Config.Database.Password = s
}
