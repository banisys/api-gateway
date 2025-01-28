package handlers

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func User(context *gin.Context) {
	target := "http://127.0.0.1:8081"
	proxy(context, target)
}

func proxy(context *gin.Context, target string) {

	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ServeHTTP(context.Writer, context.Request)
}
