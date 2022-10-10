package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type AuthReqSubmission struct{
	User string
	Pass string
	Token string
	Valid bool
}

func validateAuthRequest(data AuthReqSubmission, w http.ResponseWriter) {
	fmt.Println(">> Auth Req Validation w/")
	fmt.Println("- User:", data.User)
	fmt.Println("- Password:", data.Pass)
	fmt.Println("- token", data.Token)

	// VALIDATE - Email
	emailVerified := data.User == "c137@onecause.com"
	// VALIDATE - Password
	passVerified := data.Pass == "#th@nH@rm#y#r!$100%D0p#"
	// VALIDATE - Token
	ctMatcher := time.Now().Format("0304")
	ctTokenMatches := ctMatcher == data.Token

	fmt.Println(">> CURRENT TIME TOKEN MATCH:", ctMatcher, data.Token, ctTokenMatches)

	data.Valid = emailVerified && passVerified && ctTokenMatches

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

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

	var requestData AuthReqSubmission

	json.Unmarshal(reqBody, &requestData)
	
	// fmt.Println(">> TEST INCOMING JSON", authReqSubmission)
	validateAuthRequest(requestData, w)
}

func main() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/api/v0/auth-mpf", HandleAuthRequest_MPform)
	http.HandleFunc("/api/v0/auth-json", HandleAuthRequest_json)

	fmt.Println("- Server is running on port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}