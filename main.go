package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	ServerName    string
	ServerURL     string
	tempoExecucao float64
	status        int
}

func criarListaServidores(data [][]string) []Server {
	var servidores []Server
	for i, line := range data {
		if i > 0 {
			servidor := Server{
				ServerName: line[0],
				ServerURL:  line[1],
			}
			servidores = append(servidores, servidor)
		}
	}
	return servidores
}

func checkServer(servidores []Server) {
	for _, servidor := range servidores {
		agora := time.Now()
		get, err := http.Get(servidor.ServerURL)
		if err != nil {
			fmt.Println(err)
		}
		servidor.status = get.StatusCode
		servidor.tempoExecucao = time.Since(agora).Seconds()
		fmt.Printf("Status: [%d] Tempo de carga: [%f] URL: [%s]\n", servidor.status, servidor.tempoExecucao, servidor.ServerURL)
		//fmt.Println(servidor)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	servidores := criarListaServidores(data)
	checkServer(servidores)

}
