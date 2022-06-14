package main
import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)


type Product struct {
	ProductId	int	`json:"productId"`
	Manufuacturer	string	`json:"manaufacurer"`
	Price		int	`json:"price"`
	ProductName	string	`json:"productname"`
}





type helloHandler struct {
	Message string
}

func (f *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(f.Message))
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi baby \n"))
}


func main() {
	http.Handle("/hello", &helloHandler{Message: "hello world \n"})
	http.HandleFunc("/hi",hiHandler)

	product := &Product{
		ProductId:	123,
		Manufuacturer:	"Intel",
		Price:		20,
		ProductName:	"Laptop",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(productJSON))



	http.ListenAndServe(":5000",nil)
}
