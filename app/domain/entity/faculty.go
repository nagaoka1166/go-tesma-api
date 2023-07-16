package entity

import github.com/google/uuid

type Faculty struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}