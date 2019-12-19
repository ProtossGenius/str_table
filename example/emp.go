package main

import (
	"fmt"

	"github.com/ProtossGenius/str_table"
)

type Struct struct {
	A string `str_table:"A哈哈哈"`
	B string `str_table:"B呵呵呵"`
}

func main() {
	fmt.Println(str_table.TableLineStr(str_table.TableLineType_FIRST_LINE, []int{6, 6, 10, 10}))
	t := str_table.StructToTable("hello", []interface{}{&Struct{A: "hello", B: "World"}})
	t.AddRow([]string{"cccccc", "dddddd"})
	fmt.Println(t.RowAt(0))
	fmt.Println(str_table.TableToString(t))
}
