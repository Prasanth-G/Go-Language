
package main

import (
    "net/http"
    "log"   
    "path/filepath"
    "bufio"
    "os"
)
//ioutil    -to read files from the given path(from local machine)
//log       -to view what happen

type MyHandler struct{
}
func (p *MyHandler)ServeHTTP(writer http.ResponseWriter,request *http.Request){
    path := request.URL.Path[1:]
    log.Println(path)
    file, err := os.Open(path)
    if err == nil{
        bufferedReader :=  bufio.NewReader(file)
        var contentType string
        ext := filepath.Ext(path)
        
        m := make(map[string]string)
        m = map[string]string{
            ".html" : "text/html",
            ".css" : "text/css",
            ".js" : "application/javascript",
            ".png":"image/png",
            ".svg" :"image/svg+xml",
            ".gif" : "image/gif",
            ".mp4" : "video/mp4"}
        var ok bool
        if contentType , ok = m[ext]; ! ok{
            contentType = "text/plain"
    	}
        writer.Header().Add("Content Type", contentType)
        bufferedReader.WriteTo(writer)
    }else{
        writer.WriteHeader(404)
        writer.Write([]byte("File not Found \nERROR : " + http.StatusText(404)))
    }
}

func main(){
    mux := http.NewServeMux()
    handler1 := &MyHandler{}
    mux.Handle("/",handler1)
    http.ListenAndServe(":8080",mux)
}