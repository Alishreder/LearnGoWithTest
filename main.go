package main

import (
	. "LearnGoWithTest/dependecyInjection"
	. "LearnGoWithTest/mocking"
	"log"
	"net/http"
	"os"
	"time"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// DIModule - run code written for dependency injection chapter
func DIModule() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

// MockingModule - run code written for dependency injection chapter
func MockingModule() {
	sleeper := &ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func main() {
	// DIModule()
	MockingModule()
}
