package converter

import (
	"github.com/asp3cto/auth/internal/model"
	desc "github.com/asp3cto/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserFromService ...
func ToUserFromService(user *model.User) *desc.User {
	serviceUserInfo := ToUserInfoFromService(&user.Info)
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Name:      serviceUserInfo.Name,
		Email:     serviceUserInfo.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserInfoFromService ...
func ToUserInfoFromService(info *model.UserInfo) desc.UserInfo {
	return desc.UserInfo{
		Name:  info.Name,
		Email: info.Email,
	}
}

// ToUserInfoFromDesc ...
func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
	}
}
