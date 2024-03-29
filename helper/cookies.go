package helper

import "github.com/gofiber/fiber/v2"

func SetCookie(ctx *fiber.Ctx, token string) {
	cookie := fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		HTTPOnly: true, // Hanya bisa diakses oleh server
		Secure:   true, // Aktifkan jika Anda menggunakan HTTPS
		// SameSite: "Strict", // Atau sesuaikan sesuai kebutuhan Anda
	}
	ctx.Cookie(&cookie)
}

func ClearCookie(ctx *fiber.Ctx) {
    ctx.ClearCookie("auth_token")
}

func GetCookieValue(ctx *fiber.Ctx, name string) string {
    cookie := ctx.Cookies(name)
    return cookie
}