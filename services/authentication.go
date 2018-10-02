package services

import (
	"github.com/hamidfzm/timechi-server/stores"
	"github.com/hamidfzm/timechi-server/entities"
	"github.com/hamidfzm/timechi-server/config"
	"github.com/hamidfzm/timechi-server/factories"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"context"
	"github.com/hamidfzm/timechi-server/models"
	"github.com/hamidfzm/timechi-server/errors"
)

type Authentication struct {
	stores.User
	*config.Config
	*Handler
}

func (s *Authentication) Register(input *entities.RegisterV1) (output *entities.PublicProfileV1, err error) {
	user := factories.GetRegisterV1User(input)
	err = s.User.Create(user)
	output = factories.GetUserPublicProfileV1(user)
	return
}

func (s *Authentication) Login(input *entities.LoginV1) (output *entities.TokenV1, err error) {
	if user, err := s.User.FindByEmail(input.Email); err == nil {
		
		if err = user.VerifyPassword(input.Password); err == nil {
			output, err = factories.GetUserTokenV1(user, s.Config.Secret)
		}
	}
	return
	
}

func (s *Authentication) Profile(r *http.Request) *entities.PublicProfileV1 {
	user := s.CurrentUser(r)
	return factories.GetUserPublicProfileV1(user)
}

func (s *Authentication) TokenRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if token, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), &entities.TokenClaimsV1{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected siging method")
			}
			return []byte(s.Config.Secret), nil
		}); err == nil {
			if claims, ok := token.Claims.(*entities.TokenClaimsV1); ok && token.Valid {
				user, err := s.User.FindByID(claims.ID)
				
				if err != nil {
					s.Abort(w, http.StatusUnauthorized)
				} else {
					ctx := r.Context()
					ctx = context.WithValue(ctx, "user", user)
					r = r.WithContext(ctx)
					
					handler(w, r)
				}
				
			} else {
				s.Abort(w, http.StatusUnauthorized)
			}
			
		} else {
			s.Abort(w, http.StatusUnauthorized)
		}
	}
}

func (s *Authentication) CurrentUser(r *http.Request) *models.User {
	if user, ok := r.Context().Value("user").(*models.User); ok {
		return user
	} else {
		panic(errors.Authentication{fmt.Sprintf("Use authenticate method for %s", r.URL.String())})
	}
}
