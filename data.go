package main

type Documentation struct {
	Title string
	Links string
}

// Temp data before using db
var PageData = []Documentation{
	{
		Title: "Basic Go",
		Links: "/basicGo",
	},
	{
		Title: "Go File Structure",
		Links: "/goStructure",
	},
	{
		Title: "Go Libraries",
		Links: "/goLibraries",
	},
	{
		Title: "Things",
		Links: "/things",
	},
}
