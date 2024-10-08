package api

import (
    "encoding/json"
    "net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
    Username string // This struct defines the parameters to retrieve a coin balance, with "Username" being the key input.
}

// Coin Balance Response
type CoinBalanceResponse struct {
    // Success Code, Usually 200
    Code int // Code is used to represent the success status of the API call.
    
    // Account Balance
    Balance int64 // This holds the account balance for the user.
}

// Error Response
type Error struct {
    // Error code
    Code int // A general error code to indicate failure.

    // Error message
    Message string // A descriptive error message to help identify the problem.
}
func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var(
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandlerErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Ocurred", http.StatusInternalServerError)
	}
)