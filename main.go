package main

import (
	"net/http" //responsável por nos fornecer os métodos de mapeamento e gerenciamento de requisições
	"fmt"
	"text/template" // Gerencia templates
	"database/sql" //fornece apenas uma interface leve sobre SQL. Deve ser usado em conjunto com um driver de banco de dados
	_ "github.com/go-sql-driver/mysql" //Driver
) 

type Card struct {
	Id    int
	Cod_seg     int
	Name        string
	Date_venc   string
	Status      string
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cardService")

	if err != nil {
		panic(err.Error())
	}

	return db
}

var tmpl = template.Must(template.ParseGlob("tmpl/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	//funçao que vai lidar com as requisições e respostas do servidor
	//(resposta da requisição, tratamento da mesma)

	db := dbConn() //abre a conexão com o banco de dados

	// Realiza a consulta com banco de dados e trata erros
	selDB, err := db.Query("SELECT * FROM card")
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	structCard := Card{}

	// Monta um array para guardar os valores da struct
	res := []Card{}

	// Pega todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var id, cod_seg int
		var name, date_venc, status string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &cod_seg, &name, &date_venc, &status)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		structCard.Id = id
		structCard.Cod_seg = cod_seg
		structCard.Name = name
		structCard.Date_venc = date_venc
		structCard.Status = status

		// Junta a Struct com Array
		res = append(res, structCard)
	}

	// Abre a página Index e exibe todos os registrados na tela
	tmpl.ExecuteTemplate(w, "Index", res)

	// http.ServeFile(w, r, "./static/index.html") //servindo uma página web

	defer db.Close() //fecha a conexão
}

// Função Show exibe apenas um resultado
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Pega o ID do parametro da URL
	id := r.URL.Query().Get("id")

	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM card WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	// Monta a strcut para ser utilizada no template
	structCard := Card{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	// Pega todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var id, cod_seg int
		var name, date_venc, status string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &cod_seg, &name, &date_venc, &status)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		structCard.Id = id
		structCard.Cod_seg = cod_seg
		structCard.Name = name
		structCard.Date_venc = date_venc
		structCard.Status = status

	}

	// Mostra o template
	tmpl.ExecuteTemplate(w, "Show", structCard)

	// Fecha a conexão
	defer db.Close()

}

func main() {
	fmt.Println("Server started on: http://localhost:8080/")

	http.HandleFunc("/", Index) //(caminho, ação)
	http.HandleFunc("/show", Show)

	http.ListenAndServe(":8080", nil) //especifica em qual porta rodará a aplicaçao e o nil nesse caso informa que utilizaremos a configuração padrão do servidor do Go
}