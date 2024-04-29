package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/imanudd/clean-arch-pattern/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// init postgresql with sqlx
func InitPostgreSQLSqlx(cfg *config.MainConfig) *sqlx.DB {
	db, err := sqlx.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	log.Printf("Successfully connected to database server")

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	return db
}

// init postgress with gorm
func InitPostgreSQL(cfg *config.MainConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", cfg.PostgresHost, cfg.PostgresUsername, cfg.PostgresPassword, cfg.DBName, cfg.PostgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	log.Printf("Successfully connected to database server")

	rdb, err := db.DB()
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	rdb.SetMaxIdleConns(cfg.MaxIdleConns)
	rdb.SetMaxOpenConns(cfg.MaxOpenConns)
	rdb.SetConnMaxLifetime(time.Duration(int(time.Minute) * cfg.ConnMaxLifetime))

	return db
}
