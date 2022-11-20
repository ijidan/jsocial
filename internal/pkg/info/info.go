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
		{"1", "Application", "JSocial API"},
		{"2", "Address", addr},
		{"3", "Port", cast.ToString(port)},
		{"4", "Pid", fmt.Sprintf("%d", pid)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Info", "Desc"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	return table
}

func BuildRpcTable(config config.Config) *tablewriter.Table {
	httpAddress := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)
	pprofAddress := fmt.Sprintf("%s:%d", config.Pprof.Host, config.Pprof.Port)
	grpcAddress := fmt.Sprintf("%s:%d", config.Rpc.Host, config.Rpc.Port)
	goPsAddress := fmt.Sprintf("%s:%d", config.Ps.Host, config.Ps.Port)

	data := [][]string{
		{"1", "Application", "Jim_Service"},
		{"2", "Pprof", pprofAddress},
		{"3", "Grpc", grpcAddress},
		{"4", "GoPs", goPsAddress},
		{"5", "Http", httpAddress},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Info", "Desc"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	return table
}
