// Server1은 최소한의 "echo" 서버입니다.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 각 요청은 핸들러를 호출합니다.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler는 요청된 URL r의 Path 구성 요소를 반환합니다.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
