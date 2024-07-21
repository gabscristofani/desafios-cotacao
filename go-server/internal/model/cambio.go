package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type Cambio struct {
	USDBRL struct {
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

func (c *Cambio) SalvarCotacao(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Commit()

	data := time.Now().Format("2006-01-02 15:04:05")
	_, err = tx.ExecContext(ctx, "INSERT INTO cotacoes (valor, data) VALUES (?, ?)", c.USDBRL.Bid, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c *Cambio) NewTaxaCambio(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return err
	}

	return nil
}
