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
	"fmt"
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
	
	user := models.User{
		Name:     "test2",
		Email:    "test@test.com",
		Password: helpers.HashPassword("test"),
	}
	user.Create()
	token, _ := user.GenerateToken()
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPut, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
	
	resp = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
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
	
	user := models.User{
		Name:     "test2",
		Email:    "test@test.com",
		Password: helpers.HashPassword("test"),
	}
	user.Create()
	token, _ := user.GenerateToken()
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPut, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
	
	resp = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
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
	
	user := models.User{
		Name:     "test2",
		Email:    "test@test.com",
		Password: helpers.HashPassword("test"),
	}
	user.Create()
	token, _ := user.GenerateToken()
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPost, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
	
	resp = httptest.NewRecorder()
	reqJson, _ = json.Marshal(jsons.RegisterV1{Email: "test@test.com", Name: "test", Password: "123456"})
	body = bytes.NewBuffer(reqJson)
	req, err = http.NewRequest(http.MethodPut, url, body)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
	
	resp = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(helpers.AuthorizationHeader, token)
	if err != nil {
		t.Error(err)
	}
	controllers.Router.ServeHTTP(resp, req)
	helpers.AssertStatus(t, resp.Code, http.StatusOK)
	//json.NewDecoder(resp.Body).Decode(&respJson)
	fmt.Println(resp.Body)
}
