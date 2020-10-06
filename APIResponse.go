package main

// APIResponse - interface for any type of API Response
type APIResponse interface {
	APIResource() string
	Get() error
}
