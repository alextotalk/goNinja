package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

var (
	users = []User{{1, "Vasy"}, {2, "Galy"}}
)

func main() {
	http.HandleFunc("/users", authMiddleware(loggerMiddleware(handleUsers)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("x-id")
		if userID == "" {
			log.Printf("[%s] , %s - error. userID is not provided", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userID)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userId, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s - error. userID is invalid\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("[%s] %s by userID %s \n", r.Method, r.URL, userId)
		next(w, r)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users = append(users, user)
	w.Write([]byte("Good"))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
