package feed

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/model"
	"github.com/ijidan/jsocial/internal/repository"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

const (
	FeedTypeTxt   = 0 //文本
	FeedTypeImage = 1 //图片
	FeedTypeVideo = 2 //视频
)

var FeedTypeDbToProto = map[int32]proto_build.FeedType{
	FeedTypeTxt:   proto_build.FeedType_Txt,
	FeedTypeImage: proto_build.FeedType_Image,
	FeedTypeVideo: proto_build.FeedType_Video,
}

var FeedTypeProtoToDb = map[proto_build.FeedType]int32{
	proto_build.FeedType_Txt:   FeedTypeTxt,
	proto_build.FeedType_Image: FeedTypeImage,
	proto_build.FeedType_Video: FeedTypeVideo,
}

func ConvertFeedToProtoFeed(feed Feed) *proto_build.Feed {
	protoFeed := &proto_build.Feed{
		Id:           cast.ToUint64(feed.ID),
		UserId:       cast.ToUint64(feed.UserID),
		Content:      feed.Content,
		Type:         FeedTypeDbToProto[feed.Type],
		LikeCount:    cast.ToUint64(feed.LikeCount),
		ViewCount:    cast.ToUint64(feed.ViewCount),
		CommentCount: cast.ToUint64(feed.CommentCount),
		Operator:     cast.ToUint64(feed.Operator),
		Remark:       feed.Remark,
		Hot:          cast.ToUint64(feed.Hot),
		IsEnable:     constant.EnableDbToProto[feed.IsEnable],
		ReviewStatus: constant.ReviewStatusDbToProto[feed.ReviewStatus],
		CreatedAt:    nil,
		UpdatedAt:    nil,
		DeletedAt:    nil,
	}
	return protoFeed
}

func CreateTextFeed(db *gorm.DB, userId uint64, content string) (*proto_build.Feed, error) {
	feed := &Feed{
		UserID:       cast.ToInt64(userId),
		Content:      content,
		Type:         FeedTypeTxt,
		IsEnable:     constant.Enable,
		ReviewStatus: constant.ReviewPass,
	}
	err := db.Create(feed).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoFeed := ConvertFeedToProtoFeed(*feed)
	return protoFeed, nil
}

func CreateImageFeed(db *gorm.DB, userId uint64, content string, resource []string) (*proto_build.Feed, error) {
	feed := &Feed{
		UserID:       cast.ToInt64(userId),
		Content:      content,
		Type:         FeedTypeVideo,
		IsEnable:     constant.Enable,
		ReviewStatus: constant.ReviewPass,
		//FeedImage:    nil,
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Create(feed).Error
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		var feedImageList []*FeedImage
		if len(resource) > 0 {
			for _, v := range resource {
				feedImage := &FeedImage{
					FeedID:    feed.ID,
					ImgURL:    v,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				}
				feedImageList = append(feedImageList, feedImage)
			}

		}
		if err := tx.CreateInBatches(feedImageList, len(feedImageList)).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoFeed := ConvertFeedToProtoFeed(*feed)
	return protoFeed, nil
}

func CreateVideoFeed(db *gorm.DB, userId uint64, content string, resource []string) (*proto_build.Feed, error) {
	feed := &Feed{
		UserID:       cast.ToInt64(userId),
		Content:      content,
		Type:         FeedTypeImage,
		IsEnable:     constant.Enable,
		ReviewStatus: constant.ReviewPass,
		//FeedImage:    nil,
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Create(feed).Error
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		var feedVideoList []*FeedVideo
		if len(resource) > 0 {
			for _, v := range resource {
				feedVideo := &FeedVideo{
					FeedID:    cast.ToInt32(feed.ID),
					VideoURL:  v,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				}
				feedVideoList = append(feedVideoList, feedVideo)
			}

		}
		if err := tx.CreateInBatches(feedVideoList, len(feedVideoList)).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoFeed := ConvertFeedToProtoFeed(*feed)
	return protoFeed, nil
}

func DeleteFeed(db *gorm.DB, id uint64) error {
	var feed Feed
	err := db.Where("id=?", id).Delete(&feed).Error
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func LikeFeed(db *gorm.DB, id uint64, userId uint64) error {
	feedLike := &FeedLike{
		FeedID: cast.ToInt64(id),
		UserID: cast.ToInt64(userId),
	}
	err := db.Create(feedLike).Error
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func UnLikeFeed(db *gorm.DB, id uint64, userId uint64) error {
	var feedLike FeedLike
	err := db.Where("id=? and user_id=?", id, userId).Delete(&feedLike).Error
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func GetFeedById(db *gorm.DB, id uint64) (*proto_build.Feed, error) {
	var feed Feed
	err := db.Where("id=?", id).First(&feed).Error
	if err != nil {
		return nil, repository.ConvertModelQueryErrorToGrpcStatusError(err)
	}
	protoFeed := ConvertFeedToProtoFeed(feed)
	return protoFeed, nil
}

func GetFeedByUserId(db *gorm.DB, userId uint64) (*proto_build.Feed, error) {
	var feed Feed
	err := db.Where("user_id=?", userId).First(&feed).Error
	if err != nil {
		return nil, repository.ConvertModelQueryErrorToGrpcStatusError(err)
	}
	protoFeed := ConvertFeedToProtoFeed(feed)
	return protoFeed, nil
}

func QueryFeed(db *gorm.DB, keyword string, lastId uint64, num uint64) ([]*proto_build.Feed, *proto_build.Pager, error) {
	where := "content like ? and id>?"
	whereValue := []interface{}{"%" + keyword + "%", lastId}

	var feed Feed
	var rows []Feed
	order := "id desc"
	pager, err := model.QueryPager(db, &feed, &rows, order, 1, num, where, whereValue)
	if err != nil {
		return nil, nil, repository.ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var protoFeedList []*proto_build.Feed
	if len(rows) > 0 {
		for _, v := range rows {
			protoFeed := ConvertFeedToProtoFeed(v)
			protoFeedList = append(protoFeedList, protoFeed)
		}
	}
	protoPager := repository.ConvertPagerToProtoPager(pager)
	return protoFeedList, protoPager, nil
}

func OwnFeed(db *gorm.DB, userId uint64, reviewStatus int32, keyword string, lastId uint64, num uint64) ([]*proto_build.Feed, *proto_build.Pager, error) {
	imageFormatSlice := []interface{}{constant.ReviewInit, constant.ReviewTo, constant.ReviewPass, constant.ReviewFail}
	imageFormatSet := mapset.NewSetFromSlice(imageFormatSlice)
	var where string
	var whereValue []interface{}
	if imageFormatSet.Contains(reviewStatus) {
		where = "user_id = ? and review_status=? and id>? "
		whereValue = []interface{}{"%" + keyword + "%", reviewStatus, lastId}
	} else {
		where = "user_id = ? and id>? "
		whereValue = []interface{}{"%" + keyword + "%", lastId}
	}
	var feed Feed
	var rows []Feed
	order := "id desc"
	pager, err := model.QueryPager(db, &feed, &rows, order, 1, num, where, whereValue)
	if err != nil {
		return nil, nil, repository.ConvertModelQueryErrorToGrpcStatusError(err)
	}
	var protoFeedList []*proto_build.Feed
	if len(rows) > 0 {
		for _, v := range rows {
			protoFeed := ConvertFeedToProtoFeed(v)
			protoFeedList = append(protoFeedList, protoFeed)
		}
	}
	protoPager := repository.ConvertPagerToProtoPager(pager)
	return protoFeedList, protoPager, nil
}
