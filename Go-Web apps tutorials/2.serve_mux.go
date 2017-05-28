package main
import "net/http"
/******* Handle request using a handler
*******/

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/home",someFunc)
	defer http.ListenAndServe(":8080",mux)
	//we can bind more ports and function using mux
	mux.HandleFunc("/other",someOtherFunc)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	msg := "Welcome to home directory"
	w.Write([]byte(msg))
}
func someOtherFunc(w http.ResponseWriter, req *http.Request){
	
	w.Write([]byte("other directory"))
}