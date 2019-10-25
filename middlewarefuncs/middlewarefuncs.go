package middlewarefuncs

import (
	"fmt"
	"net/http"
)

func Index_middleware(w http.ResponseWriter,r *http.Request){
	fmt.Println("middleware index",r.RemoteAddr)
}