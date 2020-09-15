package util

import "github.com/huynhdev/go-gin-structure/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
	jwtRefreshSecret = []byte(setting.AppSetting.JwtRefreshSecret)
}
