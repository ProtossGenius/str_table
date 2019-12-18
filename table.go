package str_table

import (
	"fmt"
	"strings"
	"github.com/mattn/go-runewidth"
)


type TableRow []string
type TableCols []TableColumnItf

type TableItf interface {
	Name() string
	Rows() int
	Cols() int
	ColWidthList() []int
	RowAt(line int) (TableRow, error)
	Set(colNo, row int, val string) error
	SetByName(colName string, row int, val string) error
	Column(colNo int)(TableColumnItf, error)
	ColumnByName(colName string) (TableColumnItf, error)
	AddRow(row TableRow)
}

type TableColumnItf interface {
	Name() string
	ColumnAt(pos int) string
	SetName(name string)
	SetColumnList(list []string) error
	SetColumnAt(pos int, val string) error
	MaxWidth() int
	SetTableItf(itf TableItf)
	Len() int
}

func inSize(pos, size int) bool {
	return  pos >= 0 && pos < size
}

type TableColumn struct {
	table      TableItf
	name       string
	columnList []string
	maxWidth   int
}

func NewTableColumn(name string, table TableItf) TableColumnItf {
	return &TableColumn{name:name, table:table}
}

func (this *TableColumn) checkString(str string) {
	l := runewidth.StringWidth(str)
	if l > this.maxWidth {
		this.maxWidth = l
	}
}

func (this *TableColumn) Name() string {
	return this.name
}

func (this *TableColumn) ColumnAt(pos int) string {
	if pos < 0 || pos >= this.Len() {
		return this.columnList[pos]
	}
	return ""
}

func (this *TableColumn) SetTableItf(itf TableItf){
	this.table = itf
}

func (this *TableColumn) SetName(name string) {
	this.name = name
	this.checkString(name)
}

func (this *TableColumn) SetColumnList(list []string) error {
	tl := this.Len()
	if tl < len(list) {
		return fmt.Errorf(ErrLengthMismatch, tl, len(list))
	}
	this.columnList = list
	for _, str := range list {
		this.checkString(str)
	}
	return nil
}

func (this *TableColumn) SetColumnAt(pos int, val string) error {
	if pos < 0 || pos >= this.Len() {
		this.columnList[pos] = val
		return nil
	}
	return fmt.Errorf(ErrUndefinedPos, this.Name(), this.Len(), pos)
}

func (this *TableColumn) MaxWidth() int {
	return this.maxWidth
}

func (this *TableColumn) Len() int {
	l := this.table.Rows()
	for l > len(this.columnList) {
		this.columnList = append(this.columnList, "")
	}
	return l
}

type table_ struct {
	name        string
	cols        TableCols
	rows        int
	columnNames []string
	kMap        map[string]int
}

func NewTable(name string, colNames []string) TableItf {
	t := &table_{name:name, columnNames:colNames}
	t.kMap = map[string]int{}
	for i, cn := range colNames{
		t.kMap[cn] = i
		t.cols = append(t.cols, NewTableColumn(cn, t))
	}
	return t
}

func (this *table_) ColWidthList() []int {
	wl := make([]int, 0, this.Cols())
	for _, col := range this.cols {
		wl = append(wl, col.MaxWidth())
	}
	return wl
}

func (this *table_) Name() string {
	return this.name
}
func (this *table_) Rows() int {
	return this.rows
}

func (this *table_) Cols() int {
	return len(this.cols)
}

func (this *table_) RowAt(line int) (TableRow, error) {
	if !inSize(line, this.Rows()) {
		return nil, fmt.Errorf(ErrUndefinedPos, this.Name(), this.Rows(), line)
	}
	row := make([]string, 0, this.Cols())
	for _, col := range this.cols {
		row = append(row, col.ColumnAt(line))
	}
	return row, nil
}


func (this *table_) Set(col, row int, val string) error {
	c , err := this.Column(col)
	if err != nil{
		return err
	}
	return c.SetColumnAt(row, val)
}

func (this *table_) SetByName(colName string, row int, val string) error {
	col , err := this.ColumnByName(colName)
	if err != nil{
		return err
	}
	return col.SetColumnAt(row, val)
}

func (this *table_) ColumnByName(colName string) (TableColumnItf, error) {
	var id int
	var ok bool
	if id, ok = this.kMap[colName]; !ok {
		return nil, fmt.Errorf(ErrUnknownTableColumn, colName, strings.Join(this.columnNames, ", "))
	}
	return this.cols[id], nil
}
func (this *table_) Column(col int) (TableColumnItf, error) {
	if !inSize(col ,this.Cols()){
		return nil, fmt.Errorf(ErrUndefinedPos, this.Name(), this.Rows(), col)
	}
	return this.cols[col], nil
}
func (this *table_) AddRow(row TableRow) {
	pos := this.rows
	rowLen := len(row)
	this.rows++
	for i := range this.cols{
		v := ""
		if i < rowLen{
			v = row[i]
		}
		this.cols[i].SetColumnAt(pos, v)
	}
}
