package main

import (
	"fmt"
	"github.com/ProtossGenius/str_table"
)

func main() {
	fmt.Println(str_table.TableLineStr(str_table.TableLineType_FIRST_LINE, []int{10, 10, 10, 10}))
	for i := 0 ; i < 100; i++{
		fmt.Print("a")
	}
}
