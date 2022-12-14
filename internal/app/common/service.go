package common

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/upload"
	"github.com/ijidan/jsocial/internal/repository"
	"github.com/ijidan/jsocial/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type Service struct {
	service.BasicService
	proto_build.UnimplementedCommonServiceServer
	emailRepository *repository.EmailRepository
	smMsRepository  *repository.SmMsRepository
}

func (s *Service) UploadImage(stream proto_build.CommonService_UploadImageServer) error {
	fileUpload := upload.NewUpload()
	defer fileUpload.Close()

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return status.Error(codes.Internal, err.Error())
		}
		content := req.GetContent()
		_, err1 := fileUpload.WriteContent(content)
		if err1 != nil {
			return status.Error(codes.Internal, "write file failed")
		}
	}

	fileName := fileUpload.GetFileName()
	img, err2 := s.smMsRepository.UploadImage(fileName)
	if err2 != nil {
		return status.Error(codes.Internal, err2.Error())
	}
	rsp := &proto_build.UploadImageResponse{Url: img.Url}
	err3 := stream.SendAndClose(rsp)
	if err3 != nil {
		return status.Error(codes.Internal, err3.Error())
	}
	return nil
}

func (s *Service) SendEmail(c context.Context, req *proto_build.SendEmailRequest) (*proto_build.SendEmailResponse, error) {
	err := s.emailRepository.SendEmail(req.Receiver, req.Cc, req.Subject, req.Content)
	if err != nil {
		return &proto_build.SendEmailResponse{}, err
	}
	rsp := &proto_build.SendEmailResponse{}
	return rsp, nil
}

func NewCommonService(rpcConf *config.Rpc, emailRepository *repository.EmailRepository, smMsRepository *repository.SmMsRepository) *Service {
	instance := &Service{
		BasicService: service.BasicService{
			Name: constant.ServiceCommon,
			Host: rpcConf.Host,
			Port: rpcConf.Port,
			Ttl:  rpcConf.Ttl,
		},
		emailRepository: emailRepository,
		smMsRepository:  smMsRepository,
	}
	return instance
}
