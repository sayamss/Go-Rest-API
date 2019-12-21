package main

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

// Library Details
type LibDetails struct {
	TotalBooks     int `json:"TotalBooks"`
	TotalAvaiBooks int `json:"TotalAvaiBooks"`
	TotalIssued    int `json:"TotalIssued"`
}

// Book Data Structure
type Book struct {
	Name      string `json:"Name"`
	Issues    int    `json:"Issues"`
	Available bool   `json:"Available"`
	Issuedto  string `json:"Issuedto"`
}

// Data for Books
var AllBooks = []Book{
	Book{Name: "Lean Startup", Issues: 20, Available: true, Issuedto: "sayam"},
	Book{Name: "The secret", Issues: 40, Available: false, Issuedto: "sayam"},
	Book{Name: "Steve Jobs", Issues: 200, Available: true, Issuedto: "sayam"},
	Book{Name: "Moby Dick", Issues: 10, Available: true, Issuedto: "sayam"},
	Book{Name: "Intro to Algorithms", Issues: 5, Available: false, Issuedto: "sayam"},
	Book{Name: "Game of Thrones", Issues: 90, Available: true, Issuedto: "sayam"},
	Book{Name: "The Snape", Issues: 40, Available: true, Issuedto: "sayam"},
	Book{Name: "Hamlet", Issues: 14, Available: false, Issuedto: "sayam"},
	Book{Name: "The Wit and Wisdom", Issues: 10, Available: true, Issuedto: "sayam"},
	Book{Name: "Infinite Game", Issues: 1, Available: false, Issuedto: "sayam"},
	Book{Name: "War and Peace", Issues: 0, Available: true, Issuedto: "sayam"},
	Book{Name: "Zero to One", Issues: 100, Available: false, Issuedto: "sayam"},
	Book{Name: "Madame Bovary", Issues: 5, Available: false, Issuedto: "sayam"},
	Book{Name: "Artificial Intelligence", Issues: 34, Available: true, Issuedto: "sayam"},
	Book{Name: "Why i Killed Gandhi", Issues: 22, Available: true, Issuedto: "sayam"},
}

// Home page
func homepage(w http.ResponseWriter, r *http.Request) {

	var totalIssued, totalAvailable int
	var totalBooks = len(AllBooks)

	for i, s := range AllBooks {
		if s.Available == true {
			totalAvailable++
		}
		totalIssued += s.Issues
		i++
	}

	details := LibDetails{
		TotalBooks:     totalBooks,
		TotalAvaiBooks: totalAvailable,
		TotalIssued:    totalIssued,
	}

	fmt.Fprintf(w, "<h1>Data For Sayam's Library </h1> <br></br>")
	json.NewEncoder(w).Encode(details)
}

// Return All the Available Books
func AllBooksAvailable(w http.ResponseWriter, r *http.Request) {

	var books []Book

	for i, s := range AllBooks {
		if s.Available == true {
			books = append(books, s)
			fmt.Println(i, s.Name)
		}
	}
	fmt.Println("Endpoint hit: Available Books")
	json.NewEncoder(w).Encode(books)
}

// Return True if book is available false if not
func BookAvailable(w http.ResponseWriter, r *http.Request) {

	bookName := r.URL.Query().Get("book")
	for i, s := range AllBooks {
		if s.Name == bookName {
			if s.Available == false {
				json.NewEncoder(w).Encode(false)
			} else {
				json.NewEncoder(w).Encode(true)
			}
			break
		}
		i++
	}
}

// Return the user Book is Issued To
func IssuedUser(w http.ResponseWriter, r *http.Request) {

	bookName := r.URL.Query().Get("book")

	for i, s := range AllBooks {
		if s.Name == bookName {
			json.NewEncoder(w).Encode(s.Issuedto)
			break
		}
		i++
	}
}

// Return Most Issued
func MostIssued(w http.ResponseWriter, r *http.Request) {

	sort.Slice(AllBooks, func(i, j int) bool {
		return AllBooks[i].Issues < AllBooks[j].Issues
	})

	fmt.Println("Endpoint hit: Most Issued Book")
	json.NewEncoder(w).Encode(AllBooks[len(AllBooks)-1])
}

// Return Top Trending
func TopTrending(w http.ResponseWriter, r *http.Request) {

	var trendingBook = AllBooks[0]
	fmt.Println("Endpoint hit: Top Trending Book")
	json.NewEncoder(w).Encode(trendingBook)

}

func handleRequest() {

	http.HandleFunc("/api/", homepage)
	http.HandleFunc("/api/booksAvailable", AllBooksAvailable)
	http.HandleFunc("/api/bookAvailable", BookAvailable)
	http.HandleFunc("/api/MostIssued", MostIssued)
	http.HandleFunc("/api/IssuedTo", IssuedUser)
	http.HandleFunc("/api/TopTrending", TopTrending)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}
