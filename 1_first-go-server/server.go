package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// TODO - IMPLEMENT FROM A POST BODY EXAMPLE

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


  fmt.Println(">> TEST INCOMING MP FORM", data)
}

func HandleAuthRequest_json(w http.ResponseWriter, r *http.Request) {
	// TODO - same as above handle

	reqBody, _ := ioutil.ReadAll(r.Body)

	type AuthReq struct{
		User string
		Pass string
		Token string
	}
	var authReq AuthReq

	json.Unmarshal(reqBody, &authReq)
	
	fmt.Println(">> TEST INCOMING JSON", authReq)
}

func main() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/api/v0/auth-mpf", HandleAuthRequest_MPform)
	http.HandleFunc("/api/v0/auth-json", HandleAuthRequest_json)

	fmt.Println("- Server is running on port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}