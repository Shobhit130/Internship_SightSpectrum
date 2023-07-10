package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Products []Product // Add a slice of products owned by the user
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Inventory   int
}

var (
	users []User
)

func main() {
	// Create some sample users
	users = []User{
		{
			ID:       1,
			Username: "user1",
			Password: hashPassword("password1"),
			Email:    "abc@gmail.com",
			Products: []Product{
				{ID: 1, Name: "Product 1", Description: "Description for Product 1", Price: 9.99, Inventory: 10},
				{ID: 2, Name: "Product 2", Description: "Description for Product 2", Price: 19.99, Inventory: 5},
			},
		},
		{
			ID:       2,
			Username: "user2",
			Password: hashPassword("password2"),
			Email:    "xyz@gmail.com",
			Products: []Product{
				{ID: 3, Name: "Product 3", Description: "Description for Product 3", Price: 14.99, Inventory: 8},
			},
		},
	}

	// Initialize Gin router
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/static", "./static")

	// Initialize session middleware
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   30,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("mysession", store))

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")

	// Routes accessible without authentication
	router.GET("/", homeHandler)
	router.GET("/login", loginFormHandler)
	router.POST("/login", loginHandler)
	router.GET("/logout", logoutHandler)
	router.GET("/signup", signupFormHandler)
	router.POST("/signup", signupHandler)

	// Group routes that require authentication
	authRoutes := router.Group("/auth")
	authRoutes.Use(authMiddleware())
	{
		authRoutes.GET("/products", getProductsHandler)
		authRoutes.GET("/products/:id", getProductHandler)
		authRoutes.GET("/products/create", createProductFormHandler)
		authRoutes.POST("/products/create", createProductHandler)
		authRoutes.GET("/products/:id/edit", updateProductFormHandler)
		authRoutes.POST("/products/:id/edit", updateProductHandler)
		authRoutes.POST("/products/:id/delete", deleteProductHandler)
	}

	// Start the server
	log.Fatal(router.Run(":8080"))
}

// Middleware to check if the user is authenticated
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is logged in
		if !isLoggedIn(c) {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// User is authenticated, proceed to the next middleware or handler
		c.Next()
	}
}

// Check if the user is logged in
func isLoggedIn(c *gin.Context) bool {
	session := sessions.Default(c)
	userID := session.Get("userID")
	return userID != nil && userID.(int) > 0
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusSeeOther, "/login")
}

func signupFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func signupHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	// Retrieve any other necessary form fields

	// Check if the username or email already exists
	_, err := findUserByUsername(username)
	if err == nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"Error": "Username already exists"})
		return
	}

	_, err = findUserByEmail(email)
	if err == nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"Error": "Email already exists"})
		return
	}

	// Create a new user with the provided details
	user := User{
		ID:       generateUserID(),
		Username: username,
		Password: hashPassword(password),
		Email:    email,
		// Initialize the user's products slice
		Products: []Product{},
		// Set any other necessary fields
	}

	// Add the new user to the users slice
	users = append(users, user)

	// Redirect the user to the login page
	c.Redirect(http.StatusSeeOther, "/login")
}

func findUserByEmail(email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

func generateUserID() int {
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].ID + 1
}

func homeHandler(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/auth/products")
}

func loginFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := findUserByUsername(username)
	if err != nil || !verifyPassword(user.Password, password) {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"Error": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth/products")
}

func getProductsHandler(c *gin.Context) {
	// Retrieve the currently logged-in user
	user := getCurrentUser(c)

	c.HTML(http.StatusOK, "products.html", gin.H{"Products": user.Products})
}

func getProductHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Invalid product ID"})
		return
	}

	// Retrieve the currently logged-in user
	user := getCurrentUser(c)

	product, err := findProductByID(productID, user.Products)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"Error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "product_details.html", gin.H{"Product": product})
}

func createProductFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create_product.html", nil)
}

func createProductHandler(c *gin.Context) {
	// Retrieve the currently logged-in user
	user := getCurrentUser(c)

	name := c.PostForm("name")
	description := c.PostForm("description")
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Invalid price"})
		return
	}

	inventory, err := strconv.Atoi(c.PostForm("inventory"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Invalid inventory"})
		return
	}

	product := Product{
		ID:          generateProductID(),
		Name:        name,
		Description: description,
		Price:       price,
		Inventory:   inventory,
	}

	// Add the product to the user's products slice
	for i, u := range users {
		if u.ID == user.ID {
			users[i].Products = append(users[i].Products, product)
			break
		}
	}

	c.Redirect(http.StatusSeeOther, "/auth/products")
}

func updateProductFormHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Invalid product ID"})
		return
	}

	// Retrieve the currently logged-in user
	user := getCurrentUser(c)

	product, err := findProductByID(productID, user.Products)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"Error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "update_product.html", gin.H{"Product": product})
}

func updateProductHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Invalid product ID"})
		return
	}

	// Retrieve the currently logged-in user
	user := getCurrentUser(c)

	productIndex := -1
	for i, product := range user.Products {
		if product.ID == productID {
			productIndex = i
			break
		}
	}

	if productIndex == -1 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"Error": "Product not found"})
		return
	}

	// Update the fields of the product directly in the user's products slice
	user.Products[productIndex].Name = c.PostForm("name")
	user.Products[productIndex].Description = c.PostForm("description")
	user.Products[productIndex].Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	user.Products[productIndex].Inventory, _ = strconv.Atoi(c.PostForm("inventory"))

	c.Redirect(http.StatusSeeOther, "/auth/products")
}

func deleteProductHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Invalid product ID"})
		return
	}

	// Retrieve the currently logged-in user
	user := getCurrentUser(c)

	// Find the index of the user within the 'users' slice
	userIndex := -1
	for i, u := range users {
		if u.ID == user.ID {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"Error": "User not found"})
		return
	}

	productIndex := -1
	products := users[userIndex].Products
	for i, product := range products {
		if product.ID == productID {
			productIndex = i
			break
		}
	}

	if productIndex == -1 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"Error": "Product not found"})
		return
	}

	// Remove the product from the user's products slice using append
	users[userIndex].Products = append(products[:productIndex], products[productIndex+1:]...)

	c.Redirect(http.StatusSeeOther, "/auth/products")
}

func findUserByUsername(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

func verifyPassword(savedPassword, providedPassword string) bool {
	providedPasswordHash := hashPassword(providedPassword)
	return savedPassword == providedPasswordHash
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func findProductByID(id int, products []Product) (*Product, error) {
	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("Product not found")
}

func generateProductID() int {
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].ID + 1
}

func getCurrentUser(c *gin.Context) *User {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		return nil
	}

	userIDInt := userID.(int)
	for _, user := range users {
		if user.ID == userIDInt {
			return &user
		}
	}

	return nil
}
