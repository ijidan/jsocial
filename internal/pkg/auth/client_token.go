package auth

import (
	"context"
	"github.com/google/wire"
)

type ClientToken struct {
	Value string
}
const headerAuthorize string = "authorization"

func (t *ClientToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize:t.Value},nil
}

func (t *ClientToken)RequireTransportSecurity() bool  {
	return false
}

func NewClientToken(token string) *ClientToken  {
	return &ClientToken{
		Value: token,
	}
}

var Provider=wire.NewSet(NewClientToken)
