package call

import "C"
import (
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/sirupsen/logrus"
)

type PingCall struct {
	*BasicCall
}

func (c *PingCall) Ping() (*proto_build.PingResponse, error) {
	client := proto_build.NewPingServiceClient(c.Conn)
	rsp, err := client.Ping(c.DialCtx, &proto_build.PingRequest{})
	if err != nil {
		logrus.Fatalf("could not greet: %v", err)
	}
	return rsp, err
}

func NewPingCall(host string, port uint) *PingCall {
	basic := NewBasicCall(host, port)
	return &PingCall{BasicCall: basic}
}
