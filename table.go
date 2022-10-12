package console

import (
	"fmt"
	"io"
	"reflect"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
)

//WriteTable render the provided object as a table, object can be map, struct or array of struct(s)
func WriteTable(w io.Writer, obj interface{}, header ...string) {
	if isArray(obj) {
		renderArray(w, obj, header)
		return
	}
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

//isArray returns true if obj is array
func isArray(obj interface{}) bool {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Invalid {
		return false
	}

	return (v.Kind() == reflect.Slice || v.Kind() == reflect.Array)
}

func renderArray(w io.Writer, obj interface{}, header []string) {
	table := tablewriter.NewWriter(w)
	arr := reflect.ValueOf(obj)
	if header == nil {
		table.SetHeader(getArrayHeader(arr))
	} else {
		table.SetHeader(header)
	}

	for i := 0; i < arr.Len(); i++ {
		item := arr.Index(i)
		v := reflect.Indirect(item)
		var row []string
		for j := 0; j < v.NumField(); j++ {
			row = append(row, fmt.Sprintf("%v", v.Field(j)))
		}
		table.Append(row)
	}

	table.Render()
}

func getArrayHeader(val reflect.Value) []string {
	var header []string

	if val.Len() == 0 {
		return header
	}

	t := reflect.Indirect(val.Index(0)).Type()

	if t.Kind() != reflect.Struct {
		return header
	}

	for i := 0; i < t.NumField(); i++ {
		header = append(header, t.Field(i).Name)
	}
	return header
}
