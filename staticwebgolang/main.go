package main

import (
	"flag"
	"fmt"
	"net/http"
)

// /Users/lixingyu/Desktop/Fourier
func main() {
	staticpath := flag.String("path", "", "静态文件路径")
	flag.Parse()
	fmt.Println(*staticpath)
	http.Handle("/", http.FileServer(http.Dir(*staticpath)))
	http.ListenAndServe(":8080", nil)
}
