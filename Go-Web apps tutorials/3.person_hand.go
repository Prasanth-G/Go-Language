
package main

import (
    "net/http"
)
//we can also use an object, instead of binding a path to a func
type person struct{
    name string
    //person can also be an empty struct
}
//the name of the method should be ServeHTTP
func (p *person)ServeHTTP(writer http.ResponseWriter,request *http.Request){
    writer.Write([]byte("message you wanna say"))
}

func main(){
    mux := http.NewServeMux()
    person1 := &person{"your_name"}
    mux.Handle("/name",person1)
    http.ListenAndServe(":8080",mux)
}