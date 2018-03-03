package helpers

import (
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"testing"
	"gopkg.in/go-playground/validator.v9"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func DecodeJsonRequest(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func EncodeJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	if Config.Debug {
		if j, err := json.MarshalIndent(data, "", "\t"); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			fmt.Fprintf(w, "%s", j)
			
			return nil
		} else {
			return err
		}
	} else {
		if j, err := json.Marshal(data); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			fmt.Fprintf(w, "%s", j)
			
			return nil
		} else {
			return err
		}
	}
	
}

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	
	return hashedPassword
}

func AssertStatus(t *testing.T, responseCode int, assertCode int) {
	if responseCode != assertCode {
		t.Errorf("Wrong status code: Got %d, expected %d", responseCode, assertCode)
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%s != %s", a, b)
		t.Fatal()
	}
}

func Abort(w http.ResponseWriter, statusCode int) {
	message := http.StatusText(statusCode)
	http.Error(w, message, statusCode)
	panic(ErrorAbort{message})
}
