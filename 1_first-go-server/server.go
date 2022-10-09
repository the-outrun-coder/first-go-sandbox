package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ">> Go to /api/v0/[endpoint_goes_here]")
}

func HandleAuthRequest_MPform(w http.ResponseWriter, r *http.Request) {

	// TODO - EXCEPTION HANDLING
	// TODO - DATA VALIDATION
	// TODO - TOKEN GEN
	// TODO - RESPONSE PREP
	// TODO - PACKAGE/SUB-ROUTE MANAGEMENT

	user := r.FormValue("user")
	password := r.FormValue("password")
	token := r.FormValue("token")
	
	data := struct{
		User string
		Pass string
		Token string
	}{
		User: user,
		Pass: password,
		Token: token,
	}


  fmt.Println(">> TEST", data)
}

func main() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/api/v0/auth", HandleAuthRequest_MPform)

	fmt.Println("- Server is running on port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}