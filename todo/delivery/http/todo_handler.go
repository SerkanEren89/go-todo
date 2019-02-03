package http

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go-todo/models"
	"go-todo/todo/usecase"
	"log"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func todoAdd(useCase usecase.TodoUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding todo"
		var t *models.Todo
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		useCase.Save(t)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func todoUpdate(useCase usecase.TodoUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating todo"
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		var t *models.Todo
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = useCase.Update(id, t)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func todoFind(useCase usecase.TodoUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading todo"
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		data, err := useCase.FindById(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != models.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func todoDelete(useCase usecase.TodoUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		useCase.DeleteById(id)
		w.Header().Set("Content-Type", "application/json")
	})
}

//MakeTodoHandlers make url handlers
func MakeTodoHandlers(r *mux.Router, n negroni.Negroni, service usecase.TodoUseCase) {
	r.Handle("/v1/todo", n.With(
		negroni.Wrap(todoAdd(service)),
	)).Methods("POST", "OPTIONS").Name("todoAdd")

	r.Handle("/v1/todo/{id}", n.With(
		negroni.Wrap(todoUpdate(service)),
	)).Methods("POST", "OPTIONS").Name("todoAdd")

	r.Handle("/v1/todo/{id}", n.With(
		negroni.Wrap(todoFind(service)),
	)).Methods("GET", "OPTIONS").Name("todoFind")

	r.Handle("/v1/todo/{id}", n.With(
		negroni.Wrap(todoDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("todoDelete")
}
