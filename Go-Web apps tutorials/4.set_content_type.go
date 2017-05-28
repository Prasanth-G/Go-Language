package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

//"path/filepath"
//"io/ioutil"
//ioutil    -to read files from the given path(from local machine)
//log       -to view what happen

type MyHandler struct {
}

func (p *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	path := "client_server.py"
	log.Println(request)

	data, err := ioutil.ReadFile(string(path))
	if err == nil {
		var contentType string
		ext := filepath.Ext(path)

		m := make(map[string]string)
		m = map[string]string{
			".html": "text/html",
			".css":  "text/css",
			".js":   "application/javascript",
			".png":  "image/png",
			".svg":  "image/svg+xml",
			".gif":  "image/gif",
			".mp4":  "video/mp4",
			".mp3":  "audio/mp3"}
		var ok bool
		if contentType, ok = m[ext]; !ok {
			contentType = "text/plain"
		}
		log.Println("log for writer start")
		log.Println(writer)
		log.Println("end")
		writer.Header().Add("Content Type", contentType)
		writer.Write(data)
	} else {
		writer.WriteHeader(404)
		writer.Write([]byte("File not Found \nERROR : " + http.StatusText(404)))
	}
	//writer.Write([]byte("data to be send"))
}

func main() {
	mux := http.NewServeMux()
	handler1 := &MyHandler{}
	mux.Handle("/", handler1)
	http.ListenAndServe(":8080", mux)
}
