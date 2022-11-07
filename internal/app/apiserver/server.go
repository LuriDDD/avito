package apiserver

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"

	"encoding/json"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  *store.Store
}

func newServer(store *store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/replenishbalance", s.handleReplenishBalance()).Methods("PUT")
	s.router.HandleFunc("/reservefunds", s.handleReserveFunds()).Methods("POST")
	s.router.HandleFunc("/recognizerevenue", s.handleRecognizeRevenue()).Methods("DELETE")
	s.router.HandleFunc("/getbalance", s.handleGetBalance()).Methods("GET")
}

func (s *server) handleReplenishBalance() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var p model.User

		err := DecodeJSONBody(w, r, &p)
		if err != nil {
			var mr *MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		s.store.User().ReplenishOrCreate(&p)
		e, _ := json.Marshal(p)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(e))
	}
}

func (s *server) handleReserveFunds() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var p model.ReservedFund

		err := DecodeJSONBody(w, r, &p)
		if err != nil {
			var mr *MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		s.store.ReservedFunds().Create(&p)

		e, _ := json.Marshal(p)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(e))
	}
}

func (s *server) handleRecognizeRevenue() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var p model.ReservedFund

		err := DecodeJSONBody(w, r, &p)
		if err != nil {
			var mr *MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		s.store.ReservedFunds().Recognize(&p)

		e, _ := json.Marshal(p)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(e))
	}
}

func (s *server) handleGetBalance() http.HandlerFunc {
	type GetBalance struct {
		IDUser int
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var p GetBalance

		err := DecodeJSONBody(w, r, &p)
		if err != nil {
			var mr *MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		user, _ := s.store.User().GetBalance(p.IDUser)
		e, _ := json.Marshal(user)
		w.Header().Add("Content-Type", "application/json")

		fmt.Fprint(w, string(e))
	}
}
