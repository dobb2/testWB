package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"testWB-app/internal/config"
)

func ConnectToPostgres(conf config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
