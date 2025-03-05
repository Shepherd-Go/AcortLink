package postgres

import (
	"fmt"

	"acortlink/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
)

type pgOptions struct {
	host     string
	port     int
	user     string
	password string
	database string
}

func (p *pgOptions) getDNS() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.database)
}

func NewPostgresConnection(config config.Config) *gorm.DB {
	configPostgres := config.Postgres
	dns := pgOptions{
		host:     configPostgres.Host,
		port:     configPostgres.Port,
		user:     configPostgres.User,
		password: configPostgres.Password,
		database: configPostgres.Database,
	}

	dbInstance, err := gorm.Open(postgres.Open(dns.getDNS()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err)
	}

	println("Connected to postgres database")

	return dbInstance

}
