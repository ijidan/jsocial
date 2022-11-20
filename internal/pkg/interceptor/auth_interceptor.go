package interceptor

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		contentType := metautils.ExtractIncoming(ctx).Get("user-agent")
		if !strings.Contains(contentType, "grpcurl") && !strings.Contains(contentType, "grpc-go") {
			return nil, err
		}
		if len(token) == 0 && global.GR.Conf.App.Env != config.EnvProduction {
			return context.WithValue(ctx, "tokenUid", 0), nil
		}
	}
	claim, err1 := jwt.ParseJwtToken(token, global.GR.Conf.Jwt.Secret)
	if err1 != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err1)
	}
	newCtx := context.WithValue(ctx, "tokenUid", claim["uid"])
	return newCtx, nil

}
