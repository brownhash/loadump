package logs

import (
	"fmt"
	"log"
	"net/http"
)

// LogRequest - log http requests
func LogRequest(handler http.Handler) http.Handler {
	log.SetFlags(log.Ldate | log.Ltime)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter := fmt.Sprintf("%s %s", fmt.Sprintf(SuccessColor, r.Method), fmt.Sprintf(DebugColor, r.URL))
		log.Printf(formatter)

		handler.ServeHTTP(w, r)
	})
}
