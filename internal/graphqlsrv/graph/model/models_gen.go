// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ForksCount  int     `json:"forksCount"`
}

type Projects struct {
	Nodes []*Node `json:"nodes"`
}
