package functions

import (
	c "Project/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//make json req. with parameters
func ResJSON(name c.Names)(name_json []byte){
	n := c.Names{name.Firstname,name.Lastname}
	name_json, err := json.Marshal(n)
	if err != nil{
		ResJSON(EmpytJSON())
	}
	return
}

//parse GET req. as firstname,lastname
func ParseReq(r *http.Request) (n c.Names){
	req := r.URL.String()  // /whoami?firstname?=buse&lastname=sabah
	res := strings.Split(req, "&lastname=") //res[0]=whoami?firstname=buse, res[1]=sabah

	if len(res) == 2{
		n.Lastname = res[1]
		n.Firstname = strings.SplitAfter(res[0],"=")[1]
	}else{
		n = EmpytJSON()
	}
	return
}

//make empty JSON
func EmpytJSON() c.Names{
	n := c.Names{"",""}
	return n
}

func CheckMethod(w http.ResponseWriter, r *http.Request) bool{
	if r.Method=="GET"{
		fmt.Fprintln(w, "This endpoint can not accept GET request")
		return false
	}
	return true
}

func CheckHeaderType(w http.ResponseWriter, r *http.Request) bool{
	var headerType string
	if headerType = r.Header.Get("Content-Type"); headerType != "application/json" {
		fmt.Fprintln(w, "Invalid Header Content Type")
		return false
	}
	return true
}

func ErrCheck(err error, w http.ResponseWriter, message string) bool{
	if err != nil {
		fmt.Fprintln(w, message)
		return true
	}
	return false
}

func WebhookReq(w http.ResponseWriter, data string, conf c.Configuration){
	fmt.Println(data)

	reqBody := strings.NewReader(data)
	req, err := http.NewRequest("POST", conf.Environment.DUMMY_WEBHOOK_URL, reqBody)

	ErrCheck(err,w,"Problem occurred at request process")
	req.Header.Add("Content-Type", "application/json")

	//request-response
	_, err = http.DefaultClient.Do(req)
	ErrCheck(err,w,"Problem occurred at response process")
}

