package main

import (
    "context"
    "fmt"
    "encoding/json"
    _ "io"
    "net/http"
    "os"
    "time"
)

type Cotacao struct {
    Bid string `json:"bid"`
}

func buscaCotacao() (*Cotacao, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
   }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("status code: %d", resp.StatusCode)
    }

    var cotacao Cotacao
    if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
        return nil, err
    }

    return &cotacao, nil
}

func salvarCotacao(cotacao *Cotacao) error {
    f, err := os.Create("cotacao.json")
    if err != nil {
        return err
    }
    defer f.Close()

    if err := json.NewEncoder(f).Encode(cotacao); err != nil {
        return err
    }

    return nil
}

func main() {
    cotacao, err := buscaCotacao()
    if err != nil {
        fmt.Println("Erro ao buscar cotação:", err)
        return
    }

    if err := salvarCotacao(cotacao); err != nil {
        fmt.Println("Erro ao salvar cotação:", err)
        return
    }

    fmt.Println("Cotação salva com sucesso!")
}