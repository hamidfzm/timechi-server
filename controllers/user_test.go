package controllers_test

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/controllers"
	"github.com/hamidfzm/timechi-server/models"
	"encoding/json"
	"github.com/hamidfzm/timechi-server/jsons"
)

func TestUserControllerV1_Register(t *testing.T) {
	var resp *httptest.ResponseRecorder
	var req *http.Request
	var err error
	var body *bytes.Buffer
	var reqJson []byte
	var respJson jsons.PublicProfileV1
	url := "/v1/user/register"
	
	controllers.SetupRouter()
	models.SetupTestDatabase()
	defer models.DB.Close()
	
	resp = httptest.NewRecorder()
	body = bytes.NewBufferString("not json")
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusBadRequest)
	
	resp = httptest.NewRecorder()
	body = bytes.NewBufferString("{\"email\":\"not valid json\"}")
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusUnprocessableEntity)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "not valid email"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusUnprocessableEntity)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusCreated)
	json.NewDecoder(resp.Body).Decode(&respJson)
	helpers.AssertEqual(t, respJson.Email, "test@test.com")
	helpers.AssertEqual(t, respJson.Name, "test")
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusConflict)
}
