package model

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Cambio struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func (c *Cambio) NewCambioServer(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cotacao", nil)
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

func (c *Cambio) SaveCotacao() error {
	return ioutil.WriteFile("cotacao.txt", []byte("Dólar: "+c.USDBRL.Bid), 0644)
}
