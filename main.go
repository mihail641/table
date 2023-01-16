//Package main входная точка АПИ, запуск роутеров
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpTable/controller"
	"httpTable/user_session"
	"log"
	"net/http"
)

func main() {
	//инициация map при начале работы приложения
	user_session.Init()
	router := mux.NewRouter()
	fmt.Println("Сервер запустился")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/reset" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера ResertTable
	router.HandleFunc(
		"/reset",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			table.ResetTable(res, req)
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/add_row" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера AddRow
	router.HandleFunc(
		"/add_row",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			table.AddRow(res, req)
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/add_col" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера AddColumns
	router.HandleFunc(
		"/add_col",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			table.AddColumn(res, req)
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера GetCurrentTable
	router.HandleFunc(
		"/",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			table.GetCurrentTable(res, req)
		},
	).Methods("GET")
	log.Println("Starting HTTP server on :4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
