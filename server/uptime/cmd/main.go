package main
import (
"log"
"net/http"
)

func function_register_domain(domain string) {
	
}
func function_update_domain(domain string) {
	
}


func controller_register_domain(w http.ResponseWriter, r *http.Request) {
	// domain := r.FormValue("domain")
	function_register_domain(domain)
	w.Write([]byte("Domain registered"))
}


func controller_update_domain(w http.ResponseWriter, r *http.Request) {
	// domain := r.FormValue("domain")
	function_update_domain(domain)
	w.Write([]byte("Domain updated"))
}





func main() {
	http.HandleFunc("POST /register_domain/{$}", controller_register_domain)
	http.HandleFunc("POST /update_domain/{$id}", controller_update_domain)
	log.Fatal(http.ListenAndServe(":8080", nil))
}