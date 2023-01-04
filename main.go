package main

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var (
	requiredUser     = []byte("user")
	requiredPassword = []byte("test")
)

func BascicAuth(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		//Get the basic auth credentials
		user, password, hasAuth := r.BasicAuth()

		if !hasAuth || subtle.ConstantTimeCompare(requiredUser, []byte(user)) != 1 || subtle.ConstantTimeCompare(requiredPassword, []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}

// Root path function added
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Root path requested\n")
	fmt.Fprint(w, "<h1>This is my website</h1>")
}

// /hello path function added
func getShop(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("/hello path requested\n")
	fmt.Fprint(w, "<h1>Access granted</h1>")
}

// Program main entry point
func main() {
	// Registering function as handler for specific path, in this case getRoot and getHello
	http.HandleFunc("/", getRoot)
	http.Handle("/shop", BascicAuth(http.HandlerFunc(getShop)))
	// Error handling in specific cases
	err := http.ListenAndServe("127.0.0.1:3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed succsefully.\n")
	} else if err != nil {
		fmt.Printf("server is not started %s\n", err)
		fmt.Println("Closing, thank you!")
		os.Exit(1)
	}
}

type Product struct {
	name        string
	id          int
	description string
	price       int
}

type User struct {
	username  string
	password  string
	id        string
	firstname string
	lastname  string
	balance   float32
}

// Products GRUD

func createProduct(products map[string]interface{}, product *Product) error {
	// insert product into the map
	products[product.name] = product
	return nil
}

func getProduct(products map[string]interface{}, name string) (*Product, error) {
	// retrive product from the map
	product, ok := products[name]
	if !ok {
		return nil, fmt.Errorf("Product not found")
	}
	return product.(*Product), nil
}

func updateProduct(products map[string]interface{}, product *Product) error {
	// update product in the map
	products[product.name] = product
	return nil
}

func deleteProduct(products map[string]interface{}, name string) error {
	// delete product from the map
	delete(products, name)
	return nil
}

// User CRUD

func createUser(users map[string]interface{}, user *User) error {
	// insert user into the map
	users[user.username] = users
	return nil
}

func getUser(users map[string]interface{}, name string) (*User, error) {
	// retrive user from the map
	user, ok := users[name]
	if !ok {
		return nil, fmt.Errorf("Product not found")
	}
	return user.(*User), nil
}

func updateUser(users map[string]interface{}, user *User) error {
	// update user in the map
	users[user.username] = user
	return nil
}

func deleteUser(users map[string]interface{}, name string) error {
	// delete user from the map
	delete(users, name)
	return nil
}
