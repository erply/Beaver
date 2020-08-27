package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/clivern/beaver/internal/pkg/logger"
)

//Cors handles CORS requests
func Cors() gin.HandlerFunc {
	corsHeaders := viper.GetStringSlice("app.cors.headers")
	corsMethods := viper.GetStringSlice("app.cors.methods")
	corsOrigins := viper.GetStringSlice("app.cors.origins")

	corsConfig := cors.DefaultConfig()
	//Configured headers are added to the default list of allowed headers
	//By default the following headers are allowed: "Origin", "Content-Length", "Content-Type", "X-Correlation-ID", "X-AUTH-TOKEN"
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-Correlation-ID", "X-AUTH-TOKEN")
	if len(corsHeaders) > 0 {
		corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, corsHeaders...)
	}
	//Configured allowed methods replace the default list of allowed methods
	if len(corsMethods) > 0 {
		corsConfig.AllowMethods = corsMethods
	}
	//In case of no specified origins, allow all origins. If specified, allow only those origins
	if len(corsOrigins) > 0 {
		corsConfig.AllowOrigins = corsOrigins
	} else {
		corsConfig.AllowAllOrigins = true
	}

	logger.Infoln("Cors Headers: " + strings.Join(corsConfig.AllowHeaders, ","))
	logger.Infoln("Cors Methods: " + strings.Join(corsConfig.AllowMethods, ","))
	logger.Infoln("Cors Origins: " + strings.Join(corsConfig.AllowOrigins, ","))
	logger.Infoln("Cors AllOrigins: " + strconv.FormatBool(corsConfig.AllowAllOrigins))

	return cors.New(corsConfig)
}
