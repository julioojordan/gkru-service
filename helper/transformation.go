package helper

import "gkru-service/entity"

func ToLoginResponse(auth string, user entity.User) entity.LoginResponse {
	return entity.LoginResponse{
		Auth: auth,
		Id: user.Id,
		Username: user.Username,
		KetuaLingkungan: user.KetuaLingkungan,
		KetuaWilayah: user.KetuaWilayah,
	}
}