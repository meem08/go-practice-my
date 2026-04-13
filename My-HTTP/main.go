package main

import(
	"net/http"
	"encoding/json"
)
type User struct {
		ID   int
		Name string
	}

// func main() {
// 	users := []User{
// 		{ID: 1, Name: "Bola"},
// 		{ID:  2, Name: "Fola"},
// 	}
	
// 	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		err := json.NewEncoder(w).Encode(users)
// 		if err != nil {
// 			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
// 		}
// 	})
// 	err:= http.ListenAndServe(":3000", nil)
// 	if err != nil {
// 		fmt.Println("Server error: ", err)
// 	}
	
	
// }
type Message struct {
		Text string `json:"text"`
	}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Text: "Hello Blessing"}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":3000", nil)
}
