package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
curl -H "Content-Type: application/json" -d '{"name":"New Job"}' http://localhost:9000/jobs
*/

func main() {
	router := NewRouter()
	print("\033[H\033[2J")
	fmt.Println("\nWebserver running on port 9000\n")
	fmt.Println("The IP addresses are: ")
	fmt.Println("127.0.0.1")
	ipaddys()
	fmt.Println("\n")
	fmt.Println("Example: http://127.0.0.1:9000/jobs")
	fmt.Println("\n\nTo add records via curl:")
	fmt.Println("curl -H \"Content-Type: application/json\" -d '{\"name\":\"Call Stephen Burke\"}' http://localhost:9000/jobs")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func ipaddys() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Error obtaining addresses: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Runni!\n")
}

func JobIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jobs); err != nil {
		panic(err)
	}
}

func JobShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var jobId int
	var err error
	if jobId, err = strconv.Atoi(vars["jobId"]); err != nil {
		panic(err)
	}
	job := RepoFindJob(jobId)
	if job.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(job); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func JobCreate(w http.ResponseWriter, r *http.Request) {
	var job Job
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &job); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // un-processable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateJob(job)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

var currentId int

var jobs Jobs

func init() {
	RepoCreateJob(Job{Name: "Call Stephen Burke"})

}

func RepoFindJob(id int) Job {
	for _, t := range jobs {
		if t.Id == id {
			return t
		}
	}
	return Job{}
}

func RepoCreateJob(t Job) Job {
	currentId += 1
	t.Id = currentId
	jobs = append(jobs, t)
	return t
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"JobIndex",
		"GET",
		"/jobs",
		JobIndex,
	},
	Route{
		"JobCreate",
		"POST",
		"/jobs",
		JobCreate,
	},
	Route{
		"JobShow",
		"GET",
		"/jobs/{jobId}",
		JobShow,
	},
}

type Job struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time 	`json:"due"`
}

type Jobs []Job
