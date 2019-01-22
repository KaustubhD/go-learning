package main

import ("fmt"
        "net/http")

func home_handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, `<h1>Welcome home</h1>
  <h3>Pretty neat, huh!</h3>`)
}

func another_handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "See all of this coolness ? huh ?")
}

func main(){

http.HandleFunc("/", home_handler)
http.HandleFunc("/yo", another_handler)
http.ListenAndServe(":8080", nil)
}
