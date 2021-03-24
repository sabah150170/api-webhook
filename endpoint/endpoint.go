package endpoint

import (
	f "Project/functions"
	c "Project/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)


//for endpoint's func
type ReqEndpoint interface{
	SendResponse(http.ResponseWriter, *http.Request, c.Configuration)
}
type Alert struct {}
type Whoami struct{}
type Index struct {}


func (Index) SendResponse(w http.ResponseWriter, r *http.Request){
	n := c.Names{"buse nur","sabah"}
	fmt.Fprintln(w,bytes.NewBuffer(f.ResJSON(n)))
}

func (Whoami) SendResponse(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,bytes.NewBuffer(f.ResJSON(f.ParseReq(r))))
}

//curl -i -X POST -H "content-type: application/json" -d '{"firstname": "buse", "lastname": "sabah"}' localhost:<given_port>/alert
func (Alert) SendResponse(w http.ResponseWriter, r *http.Request, conf c.Configuration){
	if f.CheckMethod(w,r){ // only POST method is accepted
		var n c.Names
		if f.CheckHeaderType(w,r){ // only application/json content-type is accepted
			err := json.NewDecoder(r.Body).Decode(&n)
			if f.ErrCheck(err,w, "Bad Request"){
				return
			}
		}

		f.WebhookReq(w,bytes.NewBuffer(f.ResJSON(n)).String(), conf) // send to JSON req. to webhook
	}
}