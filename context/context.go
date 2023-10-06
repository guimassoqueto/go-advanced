package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	timeout := time.Duration(5000) * time.Millisecond

	// define um contexto de timeout
	timeoutContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// cria uma requisicao http usando o contexto de timeout
	// se a requisicao demorar mais que o tempo defindo, ela retorn um erro
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://picsum.photos/2000", nil)
	if err != nil {
		panic(err)
	}

	// executa a requisicao http
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// get data from http response
	imageData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Downloaded image of size %d\n", len(imageData))
}