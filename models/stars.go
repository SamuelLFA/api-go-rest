package models

type Star struct {
	Id    int    `json:id`
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Stars []Star
