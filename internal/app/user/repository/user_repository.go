package repository

import (
	"github.com/ijidan/jsocial/api/proto_build"
	model2 "github.com/ijidan/jsocial/internal/app/user/model"
	"github.com/ijidan/jsocial/internal/pkg/model"
	"github.com/ijidan/jsocial/internal/pkg/passwd"
	"github.com/ijidan/jsocial/internal/pkg/repository"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	GenderMale   = "1"
	GenderFemale = "2"
)

var GenderDbToProto = map[string]proto_build.Gender{
	GenderMale:   proto_build.Gender_Male,
	GenderFemale: proto_build.Gender_Female,
}

var GenderProtoToDb = map[proto_build.Gender]string{
	proto_build.Gender_Male:   GenderMale,
	proto_build.Gender_Female: GenderFemale,
}

func ConvertUserToProtoUser(user model2.User) *proto_build.User {
	protoUser := &proto_build.User{
		Id:        cast.ToUint64(user.ID),
		Nickname:  user.Nickname,
		Gender:    GenderDbToProto[user.Gender],
		AvatarUrl: user.AvatarURL,
		Extra:     user.Extra,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt.Time),
	}
	return protoUser
}

func CreateUser(db *gorm.DB, nickName string, password string, gender string, avatarUrl string, extra string) (*model2.User, error) {
	passwordHash, err := passwd.HashPassword(password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := &model2.User{
		Nickname:  nickName,
		Password:  passwordHash,
		Gender:    gender,
		AvatarURL: avatarUrl,
		Extra:     extra,
	}
	err1 := db.Create(user).Error
	if err1 != nil {
		return nil, status.Error(codes.Internal, err1.Error())
	}
	return user, nil
}

func GetProtoUser(db *gorm.DB, nickName string, password string) (*proto_build.User, error) {
	passwordHash, err := passwd.HashPassword(password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var user model2.User
	err1 := db.Where("nickname=? and password=?", nickName, passwordHash).First(&user).Error
	if err1 != nil {
		return nil, repository.ConvertModelQueryErrorToGrpcStatusError(err1)
	}
	protoUser := ConvertUserToProtoUser(user)
	return protoUser, nil

}

func GetProtoUserByUserId(db *gorm.DB, userId uint64) (*proto_build.User, error) {
	var user model2.User
	err := db.Where("id=?", userId).First(&user).Error
	if err != nil {
		return nil, repository.ConvertModelQueryErrorToGrpcStatusError(err)
	}
	protoUser := ConvertUserToProtoUser(user)
	return protoUser, nil
}

func QueryProtoUser(db *gorm.DB, keyword string, page uint64, pageSize uint64) ([]*proto_build.User, *proto_build.Pager, error) {
	where := "nickname like ?"
	whereValue := []interface{}{"%" + keyword + "%"}

	var user model2.User
	var rows []model2.User
	order := "id desc"
	pager, err := model.QueryPager(db, &user, &rows, order, page, pageSize, where, whereValue)
	if err != nil {
		return nil, nil, repository.ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var protoUserList []*proto_build.User
	if len(rows) > 0 {
		for _, v := range rows {
			protoUser := ConvertUserToProtoUser(v)
			protoUserList = append(protoUserList, protoUser)
		}
	}
	protoPager := repository.ConvertPagerToProtoPager(pager)
	return protoUserList, protoPager, nil
}

func UpdateUserPassword(db *gorm.DB, userId uint64, password string) error {
	err := db.Model(&model2.User{}).Where("id=?", userId).UpdateColumn("password", password).Error
	return repository.ConvertModelQueryErrorToGrpcStatusError(err)
}

func UpdateUserAvatar(db *gorm.DB, userId uint64, avatarUrl string) error {
	err := db.Model(&model2.User{}).Where("id=?", userId).UpdateColumn("avatar_url", avatarUrl).Error
	return repository.ConvertModelQueryErrorToGrpcStatusError(err)
}
