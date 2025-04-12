package main

type Documentation struct {
	Title     string
	paragraph string
}

var PageData = []Documentation{
	{
		Title:     "Pointers",
		paragraph: "Pointers point to points in your code when ... memory storage ... You get the idea ...",
	},
	{
		Title:     "Structs",
		paragraph: "Structs are structures which are used to provide custom types ... programming info ... You know what I mean ...",
	},
	{
		Title:     "For Loops",
		paragraph: "In Golang, loops are limited to only ... so programming stuff is done by doing things ... Yep.",
	},
}
