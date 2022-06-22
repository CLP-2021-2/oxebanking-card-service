package main

import (
	"net/http" //responsável por nos fornecer os métodos de mapeamento e gerenciamento de requisições
	"fmt"
	"database/sql" //fornece apenas uma interface leve sobre SQL. Deve ser usado em conjunto com um driver de banco de dados
	// _ "github.com/go-sql-driver/mysql" //Driver
) 

type Card struct {
	Num_card    int
	Cod_seg     int
	Name        string
	Date_venc   string
	Status      string
}

func Index(w http.ResponseWriter, r *http.Request) {
	//funçao que vai lidar com as requisições e respostas do servidor
	//(resposta da requisição, tratamento da mesma)
	db := dbConn() //abre a conexão com o banco de dados

	http.ServeFile(w, r, "./static/index.html") //servindo uma página web

	defer db.Close() //fecha a conexão
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser   := "grupo8"
	dbPass   := "123456"
	dbName   := "OxeBanking"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	fmt.Println("Server started on: http://localhost:8080/")

	http.HandleFunc("/", Index) //(caminho, ação)

	http.ListenAndServe(":8080", nil) //especifica em qual porta rodará a aplicaçao e o nil nesse caso informa que utilizaremos a configuração padrão do servidor do Go
}