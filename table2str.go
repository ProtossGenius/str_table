package str_table

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

type TableComp struct {
	TablePoint2LeftUp    string
	TablePoint2LeftDown  string
	TablePoint2RightUp   string
	TablePoint2RightDown string
	TablePoint4          string
	TablePoint3Down      string
	TablePoint3Left      string
	TablePoint3Right     string
	TablePoint3Up        string
	TableLineHor         string
	TableLineVer         string
}

var defaultTableComp = &TableComp{
	TablePoint2LeftUp:    "┌",
	TablePoint2LeftDown:  "└",
	TablePoint2RightUp:   "┐",
	TablePoint2RightDown: "┘",
	TablePoint4:          "┼",
	TablePoint3Down:      "┴",
	TablePoint3Left:      "├",
	TablePoint3Right:     "┤",
	TablePoint3Up:        "┬",
	TableLineHor:         "─",
	TableLineVer:         "│",
}

var defaultTableCompAscii = &TableComp{
	TablePoint2LeftUp:    "+",
	TablePoint2LeftDown:  "+",
	TablePoint2RightUp:   "+",
	TablePoint2RightDown: "+",
	TablePoint4:          "+",
	TablePoint3Down:      "+",
	TablePoint3Left:      "+",
	TablePoint3Right:     "+",
	TablePoint3Up:        "+",
	TableLineHor:         "-",
	TableLineVer:         "|",
}

type TableLinePoint struct {
	FirstPoint string
	MidPoint   string
	LastPoint  string
}

type TableDrawComp struct {
	firstLine *TableLinePoint
	midLine   *TableLinePoint
	lastLine  *TableLinePoint
}

var defaultTableDrawComp = GetTableDrawComp(defaultTableComp)
var defaultTableDrawCompAscii = GetTableDrawComp(defaultTableCompAscii)

type TableLineType int

const (
	TableLineType_FIRST_LINE TableLineType = iota
	TableLineType_MID_LINE
	TableLineType_LAST_LINE
)

func GetTableDrawComp(tableComp *TableComp) *TableDrawComp {
	return &TableDrawComp{&TableLinePoint{tableComp.TablePoint2LeftUp, tableComp.TablePoint3Up, tableComp.TablePoint2RightUp},
		&TableLinePoint{tableComp.TablePoint3Left, tableComp.TablePoint4, tableComp.TablePoint3Right},
		&TableLinePoint{tableComp.TablePoint2LeftDown, tableComp.TablePoint3Down, tableComp.TablePoint2RightDown}}
}

func fill(str string, width int) (string, error) {
	sWidth := runewidth.StringWidth(str)
	if width%sWidth != 0 {
		return "", fmt.Errorf(ErrFillSizeMismatch, width, sWidth)
	}
	num := width / sWidth
	return strings.Repeat(str, num), nil
}

//width should be even number
func TableLineStr(t TableLineType, widthList []int) (string, error) {
	return TableLineStrU(defaultTableDrawComp, defaultTableComp, t, widthList)
}

func TableLineStrAscii(t TableLineType, widthList []int) (string, error) {
	return TableLineStrU(defaultTableDrawCompAscii, defaultTableCompAscii, t, widthList)
}

func TableLineStrU(tdc *TableDrawComp, tc *TableComp, t TableLineType, widthList []int) (string, error) {
	var tlp *TableLinePoint
	switch t {
	case TableLineType_FIRST_LINE:
		tlp = tdc.firstLine
	case TableLineType_MID_LINE:
		tlp = tdc.midLine
	case TableLineType_LAST_LINE:
		tlp = tdc.lastLine
	default:
		return "", fmt.Errorf(ErrUnknownTableLineType, t)
	}
	res := tlp.FirstPoint
	for i, width := range widthList {
		if i != 0 {
			res += tlp.MidPoint
		}
		f, err := fill(tc.TableLineHor, width)
		if err != nil {
			return "", err
		}
		res += f
	}
	return res + tlp.LastPoint, nil
}

func TableToString(t TableItf) (string, error) {
	lines := []string{}
	widthList := t.ColWidthList()
	for i := range widthList {
		if widthList[i]&1 != 0 {
			widthList[i]++
		}
	}
	fmt.Println(widthList)
	tmp, err := TableLineStr(TableLineType_FIRST_LINE, widthList)
	if err != nil {
		return "", err
	}
	lines = append(lines, tmp)
	return strings.Join(lines, "\n"), nil
}
