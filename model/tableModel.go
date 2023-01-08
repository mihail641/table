package model

import (
	"httpTable/table_metadata"
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
func (t *TableModel) ResetTable() {
	table_metadata.TableInf.Row = 2
	table_metadata.TableInf.Column = 2
}

// AddRow метод модели добавляющий 1 строку после каждого обновления URL
func (t *TableModel) AddRow() {
	table_metadata.TableInf.Row++
}

// AddColumns AddRow метод модели добавляющий 1 колонку после каждого обновления URL
func (t *TableModel) AddColumns() {
	table_metadata.TableInf.Column++
}

// GetCurrentTable метод возвращающий текущие параметры таблицы без изменения
func (t *TableModel) GetCurrentTable() (row int, colum int) {
	return table_metadata.TableInf.Row, table_metadata.TableInf.Column
}
