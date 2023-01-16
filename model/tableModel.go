package model

import (
	"httpTable/user_session"
)

// TableModel труктура используется для конструктора контроллер
type TableModel struct {
}

// NewTableModel конструктор контроллера, возращающий экземпляр структуры TableModel
// со свойством model контроллера модели
func NewTableModel() *TableModel {
	return &TableModel{}
}

// ResetTable метод модели возвращающие параметры таблицы 2х2
func (t *TableModel) ResetTable(value string) {
	user_session.Set(value, user_session.Row, user_session.Column)
}

// AddRow метод модели добавляющий 1 строку после каждого обновления URL
func (t *TableModel) AddRow(value string) {
	row, column, _, _ := user_session.Get(value)
	row++
	user_session.Set(value, row, column)
}

// AddColumns AddRow метод модели добавляющий 1 колонку после каждого обновления URL
func (t *TableModel) AddColumns(value string) {
	row, column, value, _ := user_session.Get(value)
	column++
	user_session.Set(value, row, column)
}

// GetCurrentTable метод возвращающий текущие параметры таблицы без изменения
func (t *TableModel) GetCurrentTable(value string) (int, int) {
	row, column, valueNon, _ := user_session.Get(value)
	if valueNon == `` {
		row = user_session.Row
		column = user_session.Column
	}
	user_session.Set(value, row, column)
	return row, column
}
