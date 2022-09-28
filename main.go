package main

import (
	"event-challenge/events"
)

func main() {
	evts := events.GetEvents()

	// Create your in-memory projection here.
	// (No need to use an external data store.)

	for _, event := range evts {
		switch x := event.(type) {
		}
	}
}
