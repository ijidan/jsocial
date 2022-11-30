package server

import (
	"context"
	"github.com/ijidan/jsocial/internal/global"
	_ "github.com/mkevac/debugcharts"
	_ "net/http/pprof"
)

func StartRpc(ctx context.Context, cancel context.CancelFunc) {
	go func() {
		err := RunHttp(*global.GR.Conf, ctx)
		if err != nil {
			cancel()
			global.GR.Logger.Fatalf("run http server:%s", err.Error())
		}
	}()
	go func() {
		err := RunPprof(*global.GR.Conf, ctx)
		if err != nil {
			cancel()
			global.GR.Logger.Fatalf("run pprof server:%s", err.Error())
		}
	}()
	go func() {
		err := RunGrpc(*global.GR.Conf, ctx)
		if err != nil {
			cancel()
			global.GR.Logger.Fatalf("run grpc server:%s", err.Error())
		}
	}()

	go func() {
		err := RunGoPs(*global.GR.Conf, ctx)
		if err != nil {
			cancel()
			global.GR.Logger.Fatalf("run gops:%s", err.Error())
		}
	}()
	//go func() {
	//	err := dispatch.SubscribeMessage(ctx)
	//	if err != nil {
	//		cancel()
	//		global.GR.Logger.Fatalf("run kafka client:%s", err.Error())
	//	}
	//}()
}
