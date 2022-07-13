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

	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/cardService")

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

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Função Insert, insere valores no banco de dados
func Insert(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		name := r.FormValue("name")
		cod_seg := r.FormValue("cod_seg")
		date_venc := r.FormValue("date_venc")
		status := r.FormValue("status")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO card(cod_seg, name, date_venc, status) VALUES(?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		insForm.Exec(cod_seg, name, date_venc, status)

		// Exibe um log com os valores digitados no formulário
		fmt.Println("INSERT: Name: " + name)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	//Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

// Função Delete, deleta valores no banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {

	// Abre conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	nId := r.URL.Query().Get("id")

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM card WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	delForm.Exec(nId)

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com banco de dados
	db := dbConn()

	// Pega o ID do parametro da URL
	Id := r.URL.Query().Get("Id")

	selDB, err := db.Query("SELECT * FROM card WHERE id=?", Id)
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	structCard := Card{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
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

	// Mostra o template com formulário preenchido para edição
	tmpl.ExecuteTemplate(w, "Edit", structCard)

	// Fecha a conexão com o banco de dados
	defer db.Close()
}

// Função Update, atualiza valores no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		name := r.FormValue("name")
		cod_seg := r.FormValue("cod_seg")
		date_venc := r.FormValue("date_venc")
		status := r.FormValue("status")
		id := r.FormValue("id")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("UPDATE card SET cod_seg=?, name=?, date_venc=?, status=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulário com a SQL tratada e verifica erros
		insForm.Exec(id, cod_seg, name, date_venc, status)

		// Exibe um log com os valores digitados no formulario
		fmt.Println("UPDATE: Name: " + name)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

func main() {
	fmt.Println("Server started on: http://localhost:8080/")

	http.HandleFunc("/", Index) //(caminho, ação)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	//Ações
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)

	http.ListenAndServe(":8080", nil) //especifica em qual porta rodará a aplicaçao e o nil nesse caso informa que utilizaremos a configuração padrão do servidor do Go
}