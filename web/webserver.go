package main  
import (
	"fmt"
	"net/http"
)

func main() {
	userNum := map[string]int {
		"Tarima": 99,
		"Buga": 199,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "NARAKA: Bladepoint"); 
	})

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "GITHUB: "); 
	// })

	http.HandleFunc("/about", about);
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/user/"):] 
		point := userNum[name];  
		fmt.Fprintf(w, "%s: %d", name, point); 
	});

	http.ListenAndServe(":8080", nil); 
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GITHUB:"); 
}
