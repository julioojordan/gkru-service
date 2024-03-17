package helper

import "gkru-service/entity"

func ToUserResponse(user entity.User) entity.UserResponse {
	return entity.UserResponse{
		Id: user.Id,
		Username: user.Username,
	}
}