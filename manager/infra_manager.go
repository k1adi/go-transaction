package manager

import (
	"database/sql"
	"fmt"
	"go-transaction/config"
	"log"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	Connection() *sql.DB
	GetConfig() config.Config
}

type infraManager struct {
	db  *sql.DB
	cfg config.Config
}

func (i *infraManager) initDb() *sql.DB {
	var dbConf = i.cfg.DBConfig
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Password,
		dbConf.Name)

	db, err := sql.Open(i.cfg.Driver, dataSource)
	if err != nil {
		// return err
		log.Fatal("Error when trying connect to DB ", err.Error())
	}

	if err := db.Ping(); err != nil {
		// return err
		log.Fatal("Error when creating connection to DB ", err.Error())
	}

	i.db = db
	return i.db
}

func (i *infraManager) Connection() *sql.DB {
	return i.initDb()
}

func (i *infraManager) GetConfig() config.Config {
	return i.cfg
}

func NewInfraManager(configParam config.Config) InfraManager {
	infra := infraManager{
		cfg: configParam,
	}

	infra.Connection()
	return &infra
}
