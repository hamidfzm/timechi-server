package controllers_test

import (
	"testing"
	"net/http"
	"bytes"
	"github.com/hamidfzm/timechi-server/router"
	"net/http/httptest"
	"github.com/hamidfzm/timechi-server/helpers"
)

func TestUserControllerV1_Register(t *testing.T) {
	body := bytes.NewBufferString("hello")
	req, err := http.NewRequest("POST", "/v1/user", body)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()
	
	router.Router.ServeHTTP(resp, req)
	
	helpers.AssertStatus(t, resp.Code, http.StatusCreated)
}
