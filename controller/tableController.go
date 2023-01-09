package controller

import (
	"httpTable/model"
	"net/http"
	"strconv"
)

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
func (t *TableController) ResetTable(res http.ResponseWriter, req *http.Request) {
	t.model.ResetTable()
	row, column := t.model.GetCurrentTable()
	//заголовок таблицы
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>СБРОС ТАБЛИЦЫ</caption>
	`
	tableMidle := t.tableMiddle(row, column)
	tableHTML = tableHTML + tableMidle
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(tableHTML)
	res.Write(html)
}

// AddRow метод контроллера по выводу таблицы с добавлением при каждом клике дополнительной строки
func (t *TableController) AddRow(res http.ResponseWriter, req *http.Request) {
	t.model.AddRow()
	row, column := t.model.GetCurrentTable()
	//заголовок таблицы
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>ТАБЛИЦА КОГДА ПРИБАВЛЯЕТСЯ СТРОКА</caption>
	<tbody>
    <tr>`
	//цикл по конкотинации таблицы в зависимости от количества строк и столбцов
	tableMain := t.tableMiddle(row, column)
	table := tableHTML + tableMain
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(table)
	res.Write(html)
}

//AddColoms  метод контроллера по выводу таблицы с добавлением при каждом клике дополнительной строки
func (t *TableController) AddColumn(res http.ResponseWriter, req *http.Request) {
	t.model.AddColumns()
	row, column := t.model.GetCurrentTable()
	//начало таблицы
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>ТАБЛИЦА КОГДА ПРИБАВЛЯЕТСЯ КОЛОНКА</caption>
	<tbody>
    <tr>`
	tableMain := t.tableMiddle(row, column)
	table := tableHTML + tableMain
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(table)
	res.Write(html)
}

// GetCurrentTable  метод контроллера по выводу таблицы в текущем состоянии
func (t *TableController) GetCurrentTable(res http.ResponseWriter, req *http.Request) {
	row, column := t.model.GetCurrentTable()
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<caption>ТЕКУЩАЯ ТАБЛИЦА</caption>
	<tbody>
    <tr>`
	tableMain := t.tableMiddle(row, column)
	table := tableHTML + tableMain
	//отправка таблицы в браузер
	res.Header().Set("Content-Type", "text/html")
	html := []byte(table)
	res.Write(html)
}
