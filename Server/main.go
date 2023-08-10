package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Quotation struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type QuotationTable struct {
	ID    int `gorm:"primaryKey"`
	Value string
}

func main() {
	// Begin server
	http.HandleFunc("/cotacao", requestQuotationHandler)
	http.ListenAndServe(":8080", nil)
}

func requestQuotationHandler(w http.ResponseWriter, r *http.Request) {

	// Stop if URL is different of /cotacao
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Make request and return quotation
	quotation, error := requestQuotation(r)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Make insertion in cotacoes.db
	err := insertQuotation(quotation.Usdbrl.Bid)
	if err != nil {
		panic(err)
	}

	// Set up the header to application/json
	w.Header().Set("Content-Type", "application/json")

	// set up header the status 200
	w.WriteHeader(http.StatusOK)

	// Set up in JSON the response only of the value dollar
	json.NewEncoder(w).Encode(quotation.Usdbrl.Bid)

}

func requestQuotation(r *http.Request) (*Quotation, error) {
	// Cria um contexto com timeout de 200ms
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	var q Quotation

	// Faz a requisição HTTP usando o contexto com timeout
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &q)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func insertQuotation(q string) error {
	// Cria um contexto com timeout de 10ms
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	db, err := gorm.Open(sqlite.Open("cotacoes.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.WithContext(ctx)

	db.AutoMigrate(&QuotationTable{})

	db.Create(&QuotationTable{
		Value: q,
	})

	log.Printf("Insert: Did with success! value: %v", q)
	return nil
}
