package utils

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func PrintTable(cols []interface{}, rows [][]interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(convertToRow(cols))
	t.AppendRows(convertToRows(rows))
	t.SetStyle(table.StyleLight)
	t.Render()
}

func convertToRow(row []interface{}) table.Row {
	var r table.Row
	for _, col := range row {
		r = append(r, col)
	}
	return r
}

func convertToRows(inputRows [][]interface{}) []table.Row {
	var rows []table.Row
	for _, row := range inputRows {
		rows = append(rows, convertToRow(row))
	}
	return rows
}
