package manager

import (
	"database/sql"
	"fmt"
	"portfolio-api/config"

	_ "github.com/go-sql-driver/mysql"
)

type InfraManager interface {
	Connect() error
	GetDB() *sql.DB
}

type infraManager struct {
	config *config.Config
	db *sql.DB
}

func (i *infraManager) Connect() error {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		i.config.DbUser,
		i.config.DbPass,
		i.config.DbHost,
		i.config.DbPort,
		i.config.DbName,
	)

	db, err := sql.Open(i.config.DbDriver, dataSourceName)
	if err != nil {
		return err
	}

	i.db = db
	err = i.db.Ping()

	if err != nil {
		fmt.Println("Failed to PING to database")
		return err
	}

	return nil
}

func (i *infraManager) GetDB() *sql.DB {
	return i.db
}

func NewInfraManager(config *config.Config) (InfraManager, error) {
	 connection := &infraManager{
		 config: config,
	 }

	 if err := connection.Connect(); err != nil {
		 return nil, err
	 }

	 return connection, nil
}