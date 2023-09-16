package mocks

import "github.com/jmccallum-rbi/golearnings/crud/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
}
