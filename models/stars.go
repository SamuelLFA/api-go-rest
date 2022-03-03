package models

type Star struct {
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Stars []Star
