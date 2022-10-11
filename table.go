package console

import (
	"fmt"
	"io"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
)

//ToTable render the provided object as a table
func ToTable(w io.Writer, obj interface{}, header ...string) {
	var m map[string]interface{}
	if structs.IsStruct(obj) {
		m = structs.Map(obj)
	} else {
		var ok bool
		m, ok = convertMap(obj)
		if !ok {
			m = map[string]interface{}{
				"Value": fmt.Sprintf("%v", obj),
			}
		}
	}
	table := tablewriter.NewWriter(w)
	if header != nil {
		table.SetHeader(header)
	}
	for k, v := range m {
		table.Append([]string{k, fmt.Sprintf("%v", v)})
	}
	table.Render()
}

func convertMap(obj interface{}) (map[string]interface{}, bool) {
	m, ok := obj.(map[string]interface{})
	if ok {
		return m, ok
	}
	m1, ok := obj.(map[string]string)
	if ok {
		m = make(map[string]interface{})
		for k, v := range m1 {
			m[k] = v
		}
		return m, true
	}
	return nil, false
}
