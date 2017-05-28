package main
import "net/http"
/******* Handle request by a function
*******/

func main(){
	http.HandleFunc("/home",someFunc)
	http.ListenAndServe(":8080",nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello, world"))
}
