package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w,http.StatusUnprocessableEntity,erro)
		return
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest,&usuario); erro != nil {
		respostas.Erro(w,http.StatusBadRequest,erro)
		return
	} 

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w,http.StatusBadRequest,erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return
	}

	respostas.JSON(w,http.StatusCreated,usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError,erro)
		return
	}
	defer db.Close()


	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return
	}

	respostas.JSON(w, http.StatusOK,usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil{
		respostas.Erro(w,http.StatusBadRequest,erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return
	}
	respostas.JSON(w,http.StatusOK,usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível alterar informações que não sejam suas"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao,&usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()
	
	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request){
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden,errors.New("Não é possível parar de seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest,erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10,64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	seguindo, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguindo)

}

func AtualizarSenha(w http.ResponseWriter, r *http.Request){
	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return	
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"],10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	if usuarioIDNoToken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de um usuário sem ser você"))
	}


	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	var senha modelos.Senha

	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}

	if erro = seguranca.VerificarSenha(senha.Atual,senhaSalvaNoBanco); erro != nil {
		respostas.Erro(w,http.StatusUnauthorized,errors.New("A senha atual não condiz com a que está salva no banco"))
		return	
	}


	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w,http.StatusBadRequest,erro)
		return	
	}

	if erro := repositorio.AtualizarSenha(usuarioID,string(senhaComHash)); erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return	
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}