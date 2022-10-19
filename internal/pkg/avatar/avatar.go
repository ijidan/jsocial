package avatar

import (
	"fmt"
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"strk.kbt.io/projects/go/libravatar"
)

type Avatar struct {
	nickName string
}

func (avatar *Avatar) GetUrl(nickName string, cf config.App) (string, error) {
	avt := libravatar.New()
	email := fmt.Sprintf("%s@%s", nickName, cf.Name)
	avatarUrl, err := avt.FromEmail(email)
	return avatarUrl, err
}

func NewAvatar() *Avatar {
	return &Avatar{}
}

var Provider = wire.NewSet(NewAvatar)
