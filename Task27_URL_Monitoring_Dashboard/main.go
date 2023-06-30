package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// represents the status and response time of a URL
type URLStatus struct {
	URL          string
	Status       string
	//time.Duration is a built-in type in Go that represents a duration or elapsed time
	ResponseTime time.Duration
}

var (
	urlList      = []string{"https://example.com", "https://google.com","https://invalidurl.com"} // Add your URLs here
	//urlStatusMap variable is a map that stores the status of each URL. The URL serves as the key, and the URLStatus struct represents the value.
	urlStatusMap = make(map[string]*URLStatus)
	//a reader/writer mutual exclusion lock
	//sync.RWMutex provides a mechanism for synchronizing access to shared resources by multiple goroutines. It allows multiple readers to acquire the lock simultaneously, but only one writer can acquire the lock exclusively
	//creating a new instance of sync.RWMutex with its zero value, the zero value of sync.RWMutex is a fully usable and unblocked lock
	mutex        = sync.RWMutex{}
)

//mutex is used to protect concurrent access to the urlStatusMap map in the monitorURLs and checkURL functions

//Gin is a popular web framework for building web applications and APIs in the Go programming language
func main() {
	// creating a new instance of the Gin router
	router := gin.Default()

	//The LoadHTMLGlob function in Gin is specifically designed to load HTML templates. It looks for files with the .html extension by default and treats them as HTML templates. 
	//templates/*" means all files in the "templates" directory will be considered (known as glob pattern)
	router.LoadHTMLGlob("templates/*")

	//The first argument specifies the URL pattern at which the static files will be accessible. In this case, any request that starts with "/static" will be matched.
	//The second argument "static" is the file path to the directory where the static files are located. It is set to "static", which means Gin expects to find the static files in a directory named "static"

	//When a request is made to a URL matching the "/static" prefix, Gin will look for the corresponding file in the "static" directory and serve it as a static file response. For example, if we have a file named "styles.css" in the "static" directory, it can be accessed using the URL "/static/styles.css".
	router.Static("/static", "static")

	// setting up a route handler for the root URL ("/") using the HTTP GET method in the Gin web framework.
	router.GET("/", func(c *gin.Context) {
		//c.HTML is used to render an HTML response
		//setting the response status code (StatusOK) (representing a successful response)
		//specifying the HTML template to render (index.html)
		//and passing data to the template (nil)
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// setting up a route handler for the "/status" URL path using the HTTP GET method in the Gin web framework
	router.GET("/status", func(c *gin.Context) {
		//acquire a read lock from the mutex
		//This read lock allows multiple goroutines to read from urlStatusMap concurrently without blocking each other
		//acquiring a reader lock helps prevent conflicts with the goroutine that updates the urlStatusMap every 5 seconds in the monitorURLs function
		mutex.RLock()
		//After reading from the map and creating the JSON response releasing the reader lock
		defer mutex.RUnlock()

		//an empty slice statuses of type []*URLStatus to store the status information
		var statuses []*URLStatus
		for _, url := range urlList {
			statuses = append(statuses, urlStatusMap[url])
		}

		//send an HTTP response with a JSON representation of the statuses slice
		//The response will have a status code of 200 (OK) and the JSON data containing the URL statuses
		c.JSON(http.StatusOK, statuses)
	})

	//start the monitorURLs function as a separate goroutine
	//it will run concurrently with the rest of the program
	go monitorURLs()

	//starting the HTTP server and listening on port 8080 for incoming requests
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

//responsible for periodically checking the status of URLs specified in the urlList slice and updating the urlStatusMap accordingly
//The monitorURLs function iterates over the urlList slice, which contains the URLs to be monitored
//The monitorURLs function runs indefinitely in an infinite loop
func monitorURLs() {
	for {
		//sync.WaitGroup is a struct that keeps track of the number of goroutines that are currently active
		var wg sync.WaitGroup
		//run a loop for all the URLs in the urlList, it will run for every 5 seconds
		for _, url := range urlList {
			//Before starting a goroutine call wg.Add(1) to increment the count of active goroutines
			//This indicates that a new goroutine is about to start and needs to be waited for
			//basically we start a new go routine for every URL
			wg.Add(1)
			//For each URL in the urlList, a new goroutine is started
			go func(u string) {
				//decrement the active goroutine count when the URL monitoring is complete
				defer wg.Done()
				checkURL(u)
			}(url)
		}
		time.Sleep(5 * time.Second) // Check URLs every 5 seconds
	}
}

//responsible for checking the status of a given URL and updating the urlStatusMap
//the map will contain the key value as that specific URL and value will be the
//structure URLStaus containing each URL's url, status and response time in milliseconds
func checkURL(url string) {
	//recording the current time as the start time for measuring the response time
	startTime := time.Now()
	//an HTTP GET request to the specified UR
	_, err := http.Get(url)
	//calculating the duration of the request by subtracting the start time from the current time
	duration := time.Since(startTime)

	//locking the mutex using mutex.Lock() to ensure exclusive access to the urlStatusMap while updating its content
	//ensures that the JSON response in the /status endpoint handler is sent only when the map is updated completely for every URL
	mutex.Lock()
	defer mutex.Unlock()

	//If there is an error 
	if err != nil {
		urlStatusMap[url] = &URLStatus{
			URL:          url,
			Status:       "Down",
			ResponseTime: 0,
		}
	//If there is no error
	} else {
		urlStatusMap[url] = &URLStatus{
			URL:          url,
			Status:       "Up",
			//This modification ensures that the ResponseTime field contains the duration in milliseconds rather than the default time.Duration format, which represents the duration with nanosecond precision.
			ResponseTime: time.Duration(duration.Milliseconds()),
		}
	}
}
