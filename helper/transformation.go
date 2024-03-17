package helper

import "gkru-service/entity"

func ToLoginResponse(auth string) entity.LoginResponse {
	return entity.LoginResponse{
		Auth: auth,
	}
}