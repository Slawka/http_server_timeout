package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)
//  rm main.exe ; go build main.go ; .\main.exe
// $Env:GOOS = "linux"; $Env:GOARCH = "amd64"; $Env:CGO_ENABLED=0; go build

func handler(w http.ResponseWriter, r *http.Request) {
	// Получение значения параметра "name" из GET-запроса
	timeParam := r.URL.Query().Get("timeout")
	if timeParam == "" {
		timeParam = "1"
	}
	timeout, _ := strconv.Atoi(timeParam)
	log.Print("timeout: ", timeout)
	time.Sleep(time.Duration(timeout) * time.Second)
	w.Write([]byte("Hello, World! \n timeout:" + timeParam + "s"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  600 * time.Second, // ReadTimeout устанавливает максимальное время, в течение которого сервер будет ожидать полного чтения запроса.
		WriteTimeout: 600 * time.Second, // WriteTimeout устанавливает максимальное время, в течение которого сервер будет ожидать завершения записи ответа.
		IdleTimeout:  600 * time.Second, // IdleTimeout устанавливает максимальное время ожидания следующего запроса, когда соединение остается неактивным.
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not listen on :8080: %v\n", err)
	}
}
