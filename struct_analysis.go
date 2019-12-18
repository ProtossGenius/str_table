package str_table

type FStructAnalysis func(name string, val interface{}) TableItf

func MapArrayToTable(name string, val interface{}) TableItf {
	ma := val.([]map[string]interface{})
	//nameArray := []string{}
	if len(ma) == 0{
		return NewTable(name, nil)
	}
	return nil
}
