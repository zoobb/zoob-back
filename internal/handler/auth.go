package handler

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"net/http"
	"time"
	"zoob-back/internal/auth"
	"zoob-back/internal/db"
	"zoob-back/internal/types"
)

func SignUp() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var ReqBody types.AuthReqBody
		err := json.NewDecoder(req.Body).Decode(&ReqBody)
		if err != nil {
			log.Println(err)
			return
		}
		err = db.SignUp(ReqBody.Login, ReqBody.Pass)
		if err != nil {
			log.Println(err)
		}
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		if pgErr.Code == pgerrcode.UniqueViolation {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
func LogIn() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var ReqBody types.AuthReqBody
		err := json.NewDecoder(req.Body).Decode(&ReqBody)
		if err != nil {
			return
		}
		passHash, err := db.GetPassHash(ReqBody.Login)
		err = json.NewEncoder(rw).Encode(auth.CheckPass(ReqBody.Pass, passHash))
		if err != nil {
			return
		}

		token, err := auth.GenerateToken(ReqBody.Login)
		if err != nil {
			log.Println(err)
			return
		}

		http.SetCookie(rw, &http.Cookie{
			Name:     "JWT",
			Value:    token,
			Expires:  time.Now().Add(10 * time.Minute),
			HttpOnly: true,
		})

		rw.WriteHeader(http.StatusMovedPermanently)
		_, err = rw.Write([]byte("Logged in successfully"))
		if err != nil {
			return
		}
	}
}
