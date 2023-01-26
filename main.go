//Package main входная точка АПИ, запуск роутеров
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpTable/controller"
	"httpTable/userData"
	"httpTable/user_session"
	"log"
	"net/http"
)

func main() {
	//инициализация map-пользователей при запуске приложения
	userData.UsersInitializationData()
	//инициация map-cookies при начале работы приложения
	user_session.InitializationSession()
	router := mux.NewRouter()
	fmt.Println("Сервер запустился")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/reset" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера ResertTable
	router.HandleFunc(
		"/reset",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			err := table.ResetTable(res, req)
			if err != nil {
				http.Redirect(res, req, "/auth", http.StatusSeeOther)
				return
			}
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/add_row" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера AddRow
	router.HandleFunc(
		"/add_row",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			err := table.AddRow(res, req)
			if err != nil {
				http.Redirect(res, req, "/auth", http.StatusSeeOther)
				return
			}
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/add_col" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера AddColumns
	router.HandleFunc(
		"/add_col",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			err := table.AddColumn(res, req)
			if err != nil {
				http.Redirect(res, req, "/auth", http.StatusSeeOther)
				return
			}
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера GetCurrentTable
	router.HandleFunc(
		"/",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewTableController()
			err := table.GetCurrentTable(res, req)
			if err != nil {
				http.Redirect(res, req, "/auth", http.StatusSeeOther)
				return
			}
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/auth" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера GetAuthorizationUsers
	router.HandleFunc(
		"/auth",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewAuthorizationUsersController()
			num := table.GetAuthorizationUsers(res, req)
			//перенаправление на другой роут в случае если пользовать  зарегистрировался
			if num == 1 {
				http.Redirect(res, req, "/", http.StatusSeeOther)
				return
			}
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/register" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера AddNewUser
	router.HandleFunc(
		"/register",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewAuthorizationUsersController()
			num := table.AddNewUser(res, req)
			//перенаправление на другой роут в случае если пользовать зарегистрировался
			if num == 1 {
				http.Redirect(res, req, "/", http.StatusSeeOther)
				return
			}
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/delete" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера DeleteUser
	router.HandleFunc(
		"/delete",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewAuthorizationUsersController()
			table.DeleteUser(res, req)
		},
	).Methods("GET")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/logout" и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров функции контроллера LogOutUser
	router.HandleFunc(
		"/logout",
		func(res http.ResponseWriter, req *http.Request) {
			table := controller.NewAuthorizationUsersController()
			table.LogOutUser(res, req)
		},
	).Methods("GET")

	log.Println("Starting HTTP server on :4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
