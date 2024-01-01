package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func WriteJSON(resp http.ResponseWriter, status int, value any) error {
	resp.WriteHeader(status)
	resp.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(resp).Encode(value)
}

// API SERVER LOGIC
type apiFunc func(http.ResponseWriter, *http.Request) error // Function signature of the function currently using

type ApiError struct {
	Error string
}

// Decorator for the apiFunc into a router.httpHandleFunc
func makeHTTPHandleFunc(apiFunction apiFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if err := apiFunction(resp, req); err != nil {
			//FIXME: Figure out the unhandled error
			WriteJSON(resp, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddress string
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		listenAddress: address,
	}
}
func (server *APIServer) Run() {
	router := mux.NewRouter()

	// Currently used as an admin endpoint for all accounts
	router.HandleFunc("/accounts", makeHTTPHandleFunc(server.handleAccounts))
	// Currently used as an individual user account endpoint
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(server.handleGetAccount))

	log.Println("JSON API Server is running on port: ", server.listenAddress)

	//FIXME: Figure out the unhandled error
	http.ListenAndServe(server.listenAddress, router)
}

// HANDLERS
func (server *APIServer) handleAccounts(resp http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		// GET STOOF
		return server.handleGetAccount(resp, req)
	}

	if req.Method == "POST" {
		// UPDATE STOOF
		return server.handleCreateAccount(resp, req)
	}

	if req.Method == "DELETE" {
		// DELETE TARGET ACCOUNTS
		return server.handleDeleteAccount(resp, req)
	}

	return fmt.Errorf("[%s] rquest method not allowed", req.Method)
}

func (server *APIServer) handleGetAccount(resp http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)

	// TODO: db.get(id)
	fmt.Println(id)
	return WriteJSON(resp, http.StatusOK, &Account{})
}

func (server *APIServer) handleCreateAccount(resp http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleDeleteAccount(resp http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleMoneyTransfer(resp http.ResponseWriter, req *http.Request) error {
	return nil
}
