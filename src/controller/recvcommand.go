package controller

import (
	"net/http"
	"fmt"
)

func Execute(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	sql := req.Form.Get("sql")
	fmt.Println(sql)
	w.Write([]byte(sql))
}
