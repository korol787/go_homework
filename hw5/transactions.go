package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Transaction struct {
	ID          int
	Amount      float64
	Currency    string
	Type        string
	Category    string
	Date        string
	Description string
}

var transactions []Transaction

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&newTransaction)

	newTransaction.ID = len(transactions) + 1
	transactions = append(transactions, newTransaction)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransaction)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func getTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, transaction := range transactions {
		if transaction.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transaction)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedTransaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&updatedTransaction)

	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions[i] = updatedTransaction
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTransaction)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTransaction(w, r)
		case http.MethodGet:
			getTransactions(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/transactions/get", getTransactionByID)
	http.HandleFunc("/transactions/update", updateTransaction)
	http.HandleFunc("/transactions/delete", deleteTransaction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
