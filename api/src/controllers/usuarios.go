package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest,&usuario); erro != nil {
		log.Fatal(erro)
	} 

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	repositorio.Criar(usuario)

	w.Write([]byte("Criando Usuário!"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Busca todos os Usuário!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Busca apenas um Usuário!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualiza um Usuário!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deleta um Usuário!"))
}