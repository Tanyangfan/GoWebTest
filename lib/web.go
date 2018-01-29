package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var serverLog string

func handler(w http.ResponseWriter, r *http.Request) []byte {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	serverLog = fmt.Sprintf("%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		serverLog += fmt.Sprintf("Header[%q] = %q\n", k, v)
	}
	serverLog += fmt.Sprintf("Host = %q\n", r.Host)
	serverLog += fmt.Sprintf("RemoteAddr = %q\n", r.RemoteAddr)

	parseForm(w, r)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	serverLog += fmt.Sprintf("body = %q", body)

	LogD("http request", serverLog)

	return body
}

func parseForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		serverLog += fmt.Sprintf("Form[%q] = %q\n", k, v)
	}
}

func parseJSON(body []byte, clinetRequest interface{}) {
	err := json.Unmarshal(body, clinetRequest)
	if err != nil {
		LogD("parseJSON", fmt.Sprintf("err=%s", err))
	}
}

func createResponseJSON(code int, msg string, data interface{}) []byte {
	var response BaseResponse
	response.Code = code
	response.Msg = msg
	response.Data = data

	responseJSON, err := json.Marshal(response)
	if err != nil {
		LogD("createResponseJson", fmt.Sprintf("err=%s", err))
	}

	return responseJSON
}

//	注册用户
func registUser(body []byte) []byte {

	registRequest := RegistUserRequest{}
	parseJSON(body, &registRequest)

	var id int
	DbConnect.Get(&id, `SELECT id FROM test WHERE name = $1`, registRequest.Name)
	LogD("query user before inser", fmt.Sprintf("id=%d", id))
	if id != 0 {
		return createResponseJSON(http.StatusInternalServerError, "用户已存在", "{}")
	}

	DbConnect.Exec("INSERT INTO test (name) values ($1)", registRequest.Name)

	registResponse := RegistResponse{}
	DbConnect.Get(&registResponse, `SELECT * FROM test WHERE name = $1`, registRequest.Name)

	return createResponseJSON(http.StatusOK, "ok", registResponse)
}

func queryUser(body []byte) []byte {

	var userRequest QueryUserRequest
	parseJSON(body, &userRequest)

	LogD("queryUser", fmt.Sprintf("requestUser.ID=%d", userRequest.ID))
	userResponse := UserResponse{}
	DbConnect.Get(&userResponse, `SELECT * FROM test WHERE id = $1`, userRequest.ID)
	LogD("queryUser", fmt.Sprintf("responseUser=%s", userResponse.Name))
	if userResponse.ID == 0 {
		return createResponseJSON(http.StatusInternalServerError, "该用户不存在", "{}")
	}

	return createResponseJSON(http.StatusOK, "ok", userResponse)
}

func greet(body []byte) []byte {
	LogD("greet", fmt.Sprintf("Hello World!"))

	return createResponseJSON(http.StatusOK, "Hello World", "{}")
}

func check(args func(rsp []byte) []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := handler(w, r)
		w.Write(args(body))
	}
}

// Web ..
func Web() {
	InitDbConnect()
	http.HandleFunc("/greet", check(greet))
	http.HandleFunc("/registUser", check(registUser))
	http.HandleFunc("/queryUser", check(queryUser))
	http.ListenAndServe(":8080", nil)
}
