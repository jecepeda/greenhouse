package cmd

import (
	"fmt"
	"log"

	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/dcontainer"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func getDependencyContainer() (handler.DependencyContainer, func() error) {
	db, close := getDB()
	dc := dcontainer.NewDependencyContainer(db)
	dc.SetEncrypter(crypt.BEncrypter{})
	dc.SetTransactionPool(gsql.NewPool(db))
	dc.Init()

	return dc, close
}

func getDB() (*sqlx.DB, func() error) {
	host := viper.GetString("db_host")
	port := viper.GetInt("db_port")
	user := viper.GetString("db_user")
	password := viper.GetString("db_password")
	dbname := viper.GetString("db_name")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Could not connect mysql: %v", err)
	}
	db.SetMaxOpenConns(200)
	return db, db.Close
}
