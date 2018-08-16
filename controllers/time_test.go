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

func TestTimeControllerV1_StartTimer(t *testing.T) {
	var resp *httptest.ResponseRecorder
	var req *http.Request
	var err error
	var body *bytes.Buffer
	var reqJson []byte
	//var respJson jsons.TimeV1
	url := "/v1/time"
	
	controllers.SetupRouter()
	models.SetupTestDatabase()
	defer models.DB.Close()
	
	user := models.SetupTestUser()
	token, _ := user.GenerateToken()
	body = bytes.NewBufferString("not json")
	
	resp = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token+"test token")
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusUnauthorized)
	
	resp = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusBadRequest)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.StartTimerV1{Title: ""})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusUnprocessableEntity)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.StartTimerV1{Title: "test title"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.StartTimerV1{Title: "test title"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusConflict)
	
}

func TestTimeControllerV1_StopTimer(t *testing.T) {
	var resp *httptest.ResponseRecorder
	var req *http.Request
	var err error
	var body *bytes.Buffer
	var reqJson []byte
	//var respJson jsons.TimeV1
	url := "/v1/time"
	
	controllers.SetupRouter()
	models.SetupTestDatabase()
	defer models.DB.Close()
	
	user := models.SetupTestUser()
	token, _ := user.GenerateToken()
	
	user.StartTimer("test title")
	
	resp = httptest.NewRecorder()
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodDelete, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	
	resp = httptest.NewRecorder()
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodDelete, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusConflict)
	
	// TODO cover conflict status
}

func TestTimeControllerV1_Times(t *testing.T) {
	var resp *httptest.ResponseRecorder
	var req *http.Request
	var err error
	var body *bytes.Buffer
	var reqJson []byte
	//var respJson jsons.TimeV1
	url := "/v1/time"
	
	controllers.SetupRouter()
	models.SetupTestDatabase()
	defer models.DB.Close()
	
	user := models.SetupTestUser()
	token, _ := user.GenerateToken()
	
	user.StartTimer("test title")
	user.StopTimer()
	
	user.StartTimer("test title 2")
	user.StopTimer()
	
	
	user.StartTimer("test title 2")
	user.StopTimer()
	
	resp = httptest.NewRecorder()
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodGet, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	
	// TODO check items count
}
