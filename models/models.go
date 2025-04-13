package models

type Topic struct {
	Title     string
	Paragraph string
}

// Temp data before using db
var Topics = []Topic{
	{
		Title:     "Pointers",
		Paragraph: "Pointers point to points in your code when ... memory storage ... You get the idea ...",
	},
	{
		Title:     "Structs",
		Paragraph: "Structs are structures which are used to provide custom types ... programming info ... You know what I mean ...",
	},
	{
		Title:     "For Loops",
		Paragraph: "In Golang, loops are limited to only ... so programming stuff is done by doing things ... Yep.",
	},
	{
		Title:     "Arrays",
		Paragraph: "In Golang, to make an array of you must declare its length... this is useful because reasons and stuff ... I think it's pretty obvious and the fact that you didn't know is sort of embaressing ... panguines ... programming things ...",
	},
	{
		Title:     "Slices",
		Paragraph: "In Golang, to make an array of uncertain length... this is useful because reasons ... I think it's pretty obvious and the fact that you didn't know is sort of embaressing ...",
	},
}
