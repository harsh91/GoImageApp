package handler

import (
	"github.com/momazia/GoProject/session"
	"log"
	"net/http"
)

type APlusTemplate struct {
	Header Header
	Data   interface{}
}

type Header struct {
	IsLoggedIn bool
}

func GetAPlusTemplateHeader(req *http.Request, data interface{}) APlusTemplate {
	log.Println("Logged Email:")
	log.Println(session.GetUser(req).Email)
	return APlusTemplate{
		Header: Header{
			IsLoggedIn: session.GetUser(req).Email != "",
		},
		Data: data,
	}
}
