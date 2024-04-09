package router

import "github.com/gorilla/mux"

//Gerar retorna um router com rotas configuradas
func Gerar() *mux.Router{
	return mux.NewRouter()
}