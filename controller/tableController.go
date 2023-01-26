package controller

import (
	"fmt"
	"httpTable/cookies"
	"httpTable/model"
	"net/http"
	"strconv"
)

const nameCookie = "html-table_coookie"

// TableController cтруктура используется для конструктора контроллер
type TableController struct {
	model *model.TableModel
}

// NewTableController конструктор контроллера, возращающий экземпляр структуры TableController
// со свойством model контроллера модели
func NewTableController() *TableController {
	return &TableController{
		model: model.NewTableModel(),
	}
}

//приватный метод для вывода повторяющихся частей таблицы на печать
func (t *TableController) tableMiddle(row int, column int) string {
	var numberRow int
	var tableMain string
	//цикл по конкатенации таблицы в зависимости от количества строк и столбцов
	for i := 0; i < row; i++ {
		var tableMiddle string
		for j := 0; j < column; j++ {
			numberRow = numberRow + 1
			stringNumberRow := strconv.Itoa(numberRow)
			tableColumn := `<td>Значение` + stringNumberRow + `</td>`
			tableMiddle = tableMiddle + tableColumn
		}
		//начало новой строки в таблице
		lineBreak := `</tr>
		<tr>`
		tableMain = tableMain + tableMiddle + lineBreak
	}
	return tableMain
}

// ResertTable метод сброса таблицы до состояния 2х2
func (t *TableController) ResetTable(res http.ResponseWriter, req *http.Request) error {
	value, err := TokenVerification(res, req)
	if err != nil {
		fmt.Sprintf("Проблема с токеном в контроллера в методе ResetTable", err)
		return err
	}
	t.model.ResetTable(value)
	row, column, name := t.model.GetCurrentTable(value)
	//заголовок таблицы
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>СБРОС ТАБЛИЦЫ ` + name + `</caption>
	`
	tableMidle := t.tableMiddle(row, column)
	tableFinish :=
		`<a href="../add_row"> ДОБАВЛЕНИЕ СТРОК </a> 
		 <a href="../add_col">ДОБАВЛЕНИЕ КОЛОНОК  </a> 
		 <a href="../">ТЕКУЩЕЕ СОСТОЯНИЕ ТАБЛИЦЫ  </a> 
	     <a href="../reset">СБРОС ТАБЛИЦЫ  </a> `
	tableHTML = tableHTML + tableMidle + tableFinish
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(tableHTML)
	res.Write(html)
	return err
}

// AddRow метод контроллера по выводу таблицы с добавлением при каждом клике дополнительной строки
func (t *TableController) AddRow(res http.ResponseWriter, req *http.Request) error {
	value, err := TokenVerification(res, req)
	if err != nil {
		fmt.Sprintf("Проблема с токеном в контроллере в методе AddRow", err)
		return err
	}
	t.model.AddRow(value)
	row, column, name := t.model.GetCurrentTable(value)
	//заголовок таблицы
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>ДОБАВЛЕНИЕ СТРОК ` + name + `</caption>
	<tbody>
	<tr>`
	//цикл по конкотинации таблицы в зависимости от количества строк и столбцов
	tableMain := t.tableMiddle(row, column)
	tableFinish :=
		`<a href="../add_row"> ДОБАВЛЕНИЕ СТРОК  </a>
		 <a href="../add_col"> ДОБАВЛЕНИЕ КОЛОНОК  </a>
		 <a href="../">ТЕКУЩЕЕ СОСТОЯНИЕ ТАБЛИЦЫ  </a>
	     <a href="../reset">СБРОС ТАБЛИЦЫ  </a>`

	table := tableHTML + tableMain + tableFinish
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(table)
	res.Write(html)
	return nil
}

//AddColoms  метод контроллера по выводу таблицы с добавлением при каждом клике дополнительной строки
func (t *TableController) AddColumn(res http.ResponseWriter, req *http.Request) error {
	value, err := TokenVerification(res, req)
	if err != nil {
		fmt.Sprintf("Проблема с токеном в контроллера в методе AddColumn", err)
		return err
	}
	t.model.AddColumns(value)
	row, column, name := t.model.GetCurrentTable(value)
	//начало таблицы
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>ДОБАВЛЕНИЕ СТОЛБЦОВ ` + name + `</caption>
	<tbody>
	<tr>`
	tableMain := t.tableMiddle(row, column)
	tableFinish :=
		`<a href="../add_row">ДОБАВЛЕНИЕ СТРОК  </a>
		 <a href="../add_col">ДОБАВЛЕНИЕ КОЛОНОК  </a>
		 <a href="../">ТЕКУЩЕЕ СОСТОЯНИЕ ТАБЛИЦЫ  </a>
	     <a href="../reset">СБРОС ТАБЛИЦЫ  </a>`

	table := tableHTML + tableMain + tableFinish
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(table)
	res.Write(html)
	return nil
}

// GetCurrentTable  метод контроллера по выводу таблицы в текущем состоянии
func (t *TableController) GetCurrentTable(res http.ResponseWriter, req *http.Request) error {
	value, err := TokenVerification(res, req)
	if err != nil {
		fmt.Sprintf("Проблема с токеном в контроллере в методе GetCurrentTable", err)
		return err
	}
	row, column, name := t.model.GetCurrentTable(value)
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>ТЕКУЩАЯ ТАБЛИЦА ` + name + `</caption>
	<tbody>
	<tr>`
	tableMain := t.tableMiddle(row, column)
	tableFinish :=
		`<a href="../add_row">ДОБАВЛЕНИЕ СТРОК</a>
		 <a href="../add_col">ДОБАВЛЕНИЕ КОЛОНОК   </a>
		 <a href="../">ТЕКУЩЕЕ СОСТОЯНИЕ ТАБЛИЦЫ   </a>
	     <a href="../reset">СБРОС ТАБЛИЦЫ  </a>`

	table := tableHTML + tableMain + tableFinish

	if value == "" {
		table = ""
	}
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(table)
	res.Write(html)
	return nil
}

// TokenVerification функция для проверки соответсвия куки файлам- файлам пользователя
func TokenVerification(res http.ResponseWriter, req *http.Request) (string, error) {
	token, err := cookies.GetCookie(req, nameCookie)
	if err != nil {
		fmt.Printf("Нет cookies файлов или неправильный токен метод tokenVerification")
		return "", err
	}
	return token, err
}

// TokenGet метод для полученияТокенов и имен файлов- cookie
func TokenGet(res http.ResponseWriter, req *http.Request) (string, error) {
	cookie := &http.Cookie{Name: nameCookie}
	value, err := cookies.SetCookie(res, *cookie)
	if err != nil {
		fmt.Printf("Проблема с токеном в контроллере в методе tokenVerification%v\n\t", err)
		return value, err
	} else {
		fmt.Printf("ВСЕ ХОРОШО")
	}
	return value, err
}
