package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/banisys/api-gateway/internal/models"
	"github.com/banisys/api-gateway/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/banisys/api-gateway/user_service_grpc"
)

type MockGrpcSignup struct {
}

func (g MockGrpcSignup) GrpcSignup(user models.User) *pb.UserServiceRes {
	return &pb.UserServiceRes{}
}

func TestSignup(t *testing.T) {

	route := gin.Default()

	var MockGrpcSignup MockGrpcSignup

	routes.RegisterRoutes(route, MockGrpcSignup)

	testBody := []byte(`{"name":"asali","email":"test@example.com","password":"123"}`)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, req)

	// assert userId not null

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var responseBody map[string]any
	json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	require.NotNil(t, responseBody["user_id"])

}

// test signup validation
