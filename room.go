package main

type room struct {
	// foward is a channel that holds incoming messages
	// that should be fowarded to hte other clients.
	foward chan []byte
}
