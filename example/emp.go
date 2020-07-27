package main

import (
	"fmt"

	"github.com/ProtossGenius/str_table"
	"github.com/mattn/go-runewidth"
)

type Struct struct {
	A string `str_table:"A哈哈哈"`
	B string `str_table:"B呵呵呵"`
}

func main() {
	fmt.Println(runewidth.StringWidth("A哈哈哈"))
	fmt.Println("012345678901234567890123456789")
	fmt.Println(str_table.TableLineStr(str_table.TableLineType_FIRST_LINE, []int{6, 6, 10, 10}))
	t := str_table.StructToTable("hello", []interface{}{&Struct{A: "hello", B: "World"}})
	t.AddRow([]string{"cccccc", "dddddd"})
	fmt.Println(t.RowAt(0))
	fmt.Println(str_table.TableToString(t))
}
