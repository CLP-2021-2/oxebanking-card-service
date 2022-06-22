package main

import ("net/http") //responsável por nos fornecer os métodos de mapeamento e gerenciamento de requisições

func handler(w http.ResponseWriter, r *http.Request) {
	//funçao que vai lidar com as requisições e respostas do servidor
	//(resposta da requisição, tratamento da mesma)

	http.ServeFile(w, r, "./static/index.html") //servindo uma página web
}

func main() {
	http.HandleFunc("/", handler) //(caminho, ação)
	http.ListenAndServe(":8080", nil) //especifica em qual porta rodará a aplicaçao e o nil nesse caso informa que utilizaremos a configuração padrão do servidor do Go
}