package helpers

import (
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func DecodeJson(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func EncodeJson(w http.ResponseWriter, data interface{}) error {
	if j, err := json.Marshal(data); err == nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", j)
		
		return nil
	} else {
		return err
	}
}

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	
	return hashedPassword
}
