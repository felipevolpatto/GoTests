package main

import ( 

  "fmt"
	"net/http"
)


/* handles http requests for / */
func handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w , "Hello world");
}

func main(){

	http.HandleFunc( "/" , handler )
	http.ListenAndServe(":8080" , nil)
}
