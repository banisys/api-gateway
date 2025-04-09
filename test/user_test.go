package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/banisys/api-gateway/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_sigup(t *testing.T) {

	route := gin.Default()
	routes.RegisterRoutes(route)

	testBody := []byte(`{"name":"asali","email":"test@example.com","password":"123"}`)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

}
