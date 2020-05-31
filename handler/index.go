package handler

import (
	"dbwork/tpl"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//phrase := map[string]interface{}{}
		if err := tpl.Index.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}
