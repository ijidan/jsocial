package service

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Register struct {
	Client              *clientv3.Client
	AppName             string
	Service             BasicService
	ServiceLeaseIdMap   map[BasicService]clientv3.LeaseID
	ServiceKeepAliveMap map[BasicService]<-chan *clientv3.LeaseKeepAliveResponse
	ServiceCloseMap     map[BasicService]chan bool
}

func (r *Register) RegisterService(srvList ...BasicService) {
	for _, srv := range srvList {
		LeaseRsp, _ := r.Client.Grant(context.Background(), srv.GetTTL())
		keepAliveCh, _ := r.Client.KeepAlive(context.Background(), LeaseRsp.ID)
		_, _ = r.Client.KV.Put(context.Background(), srv.GetName(), srv.GetAddress(), clientv3.WithLease(LeaseRsp.ID))
		r.ServiceLeaseIdMap[srv] = LeaseRsp.ID
		r.ServiceKeepAliveMap[srv] = keepAliveCh

	}
	for _, srv := range srvList {
		go r.CloseSrv(srv)
	}
	go r.KeepAlive()
}

func (r *Register) UnRegisterService(srv BasicService) {
	leaseId := r.ServiceLeaseIdMap[srv]
	_, _ = r.Client.Revoke(context.Background(), leaseId)
	_, _ = r.Client.KV.Delete(context.Background(), srv.GetName())
}

func (r *Register) Stop(srv BasicService) {
	closeCh := r.ServiceCloseMap[srv]
	closeCh <- true
}

func (r *Register) CloseSrv(srv BasicService) {
	closeCh := r.ServiceCloseMap[srv]
	for {
		select {
		case res := <-closeCh:
			if res {
				r.UnRegisterService(srv)
			}
		}
	}
}

func (r *Register) KeepAlive() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			for _, _ch := range r.ServiceKeepAliveMap {
				_=<-_ch
			}
		}
	}
}

func NewServiceRegister(client *clientv3.Client, appName string) *Register {
	serviceRegister := &Register{
		Client:              client,
		AppName:             appName,
		ServiceLeaseIdMap:   make(map[BasicService]clientv3.LeaseID),
		ServiceKeepAliveMap: make(map[BasicService]<-chan *clientv3.LeaseKeepAliveResponse),
		ServiceCloseMap:     make(map[BasicService]chan bool),
	}
	return serviceRegister
}
