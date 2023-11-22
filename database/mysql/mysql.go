package mysql

import (
	"database/sql"
	"log"
)

type DBConnector interface {
	Connect() (*sql.DB, error)
}

type MySQLConnector struct {
	Credentials string
}

func (m *MySQLConnector) Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", m.Credentials)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("[InitDB: Conexão bem-sucedida ao AWS RDS]")

	return db, nil
}
