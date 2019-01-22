package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	startIDService()
}

// start http service of id generator
func startIDService() {
	idService := &http.Server{Addr: ":9999"}
	http.HandleFunc("/", idServiceResponse) // 设置访问的路由

	err := idService.ListenAndServe()
	if err != nil {
		log.Fatal("StartIDService: ", err)
	}
}

// Response
func idServiceResponse(w http.ResponseWriter, r *http.Request) {
	id := GetIDInstance().NextID()
	fmt.Fprintf(w, "%d", id) // 这个写入到w的是输出到客户端的
}
