package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/hamidfzm/timechi-server/models"
	"github.com/hamidfzm/timechi-server/helpers"
	"context"
)

func Authenticate(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		
		if token, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte(helpers.Config.Secret), nil
		}); err == nil {
			if claims, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
				user, err := models.FindUserByID(claims.ID)
				
				if err != nil {
					helpers.Abort(w, http.StatusUnauthorized)
				} else {
					ctx := r.Context()
					ctx = context.WithValue(ctx, "user", user)
					r = r.WithContext(ctx)
					
					handler(w, r, ps)
				}
				
			} else {
				helpers.Abort(w, http.StatusUnauthorized)
			}
			
		} else {
			helpers.Abort(w, http.StatusUnauthorized)
		}
	}
}
