package routers

import (
	"net/http"

	"example.com/apiEnronMail/service"
)

func Document(w http.ResponseWriter, r *http.Request) {

	service.Get(w)

}
