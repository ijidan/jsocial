package info

import (
	"fmt"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
	"os"
)

func BuildApiAddr(conf config.Http) string {
	host := conf.Host
	port := conf.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	return addr
}
func BuildApiTable(conf config.Http, pid int) *tablewriter.Table {
	port := conf.Port
	addr := BuildApiAddr(conf)
	data := [][]string{
		[]string{"1", "Application", "Jim"},
		[]string{"2", "Address", addr},
		[]string{"3", "Port", cast.ToString(port)},
		[]string{"4", "Pid", fmt.Sprintf("%d", pid)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Info", "Desc"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	return table
}
