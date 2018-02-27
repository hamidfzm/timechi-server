package controllers_test

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/controllers"
)

func TestUserControllerV1_Register(t *testing.T) {
	body := bytes.NewBufferString("hello")
	req, err := http.NewRequest("POST", "/v1/user", body)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()
	
	controllers.SetupRouter()
	controllers.Router.ServeHTTP(resp, req)
	
	helpers.AssertStatus(t, resp.Code, http.StatusCreated)
}
