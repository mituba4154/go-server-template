package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	goapi "github.com/mituba4154/goapi-sdk"
	"github.com/mituba4154/go-server-template/modules/example"
)

func main() {
	javaAddr := getenv("GOAPI_JAVA_ADDR", "http://localhost:8766")
	goapi.Init(javaAddr + "/rpc")

	// ---- モジュール登録 ----
	// 新しいモジュールを追加するときはここに1行追加するだけ
	example.Register()
	// ----------------------

	addr := getenv("GOAPI_ADDR", "localhost:8765")
	http.HandleFunc("/rpc", goapi.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Printf("[GoServer] listening on %s", addr)
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("[GoServer] shutdown")
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
