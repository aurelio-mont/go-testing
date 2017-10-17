package main

type Person struct {
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person