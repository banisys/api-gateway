package handlers

import (
	"bytes"
	"io"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {

	target := "127.0.0.1:8081"
	proxy(context, target)

}

func proxy(c *gin.Context, target string) {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// خواندن و بازنشانی Body درخواست
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(body)) // بازنشانی Body

	// اضافه کردن یک هدر نمونه
	c.Request.Header.Set("X-Example", "MyGateway")

	// ارسال درخواست به مقصد
	proxy.ServeHTTP(c.Writer, c.Request)
}
