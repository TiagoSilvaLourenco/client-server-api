package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}

	res, err := c.Get("http://localhost:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := "Dólar: " + string(body)

	rec, err := f.Write([]byte(result))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso: Tamanho: %d bytes", rec)

}
