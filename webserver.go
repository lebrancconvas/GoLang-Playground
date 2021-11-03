package main 
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "NARAKA: Bladepoint"); 
	})

http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GITHUB: "); 
})

	http.ListenAndServe(":8080", nil); 
}