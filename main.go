package main

import (
	"flag"
	"fmt"

	"github.com/JohnGrimm/Stress-Test-fg/cmd"
	"github.com/JohnGrimm/Stress-Test-fg/internal/stresstest"
)

func main() {
	url := flag.String("url", "", "URL a ser testada")
	request := flag.Int("requests", 1, "Número de requisições a serem feitas")
	concorrencias := flag.Int("concurrency", 1, "Número de requisições concorrentes")

	flag.Parse()

	config := stresstest.Config{
		Url:          *url,
		Requisicoes:  *request,
		Concorrencia: *concorrencias,
	}

	if config.Url == "" {
		fmt.Println("URL é obrigatória")
		return
	}

	cmd.Execute(config)
}
