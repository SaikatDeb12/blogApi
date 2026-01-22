package storage

import (
	"blog/internal/model"
	"time"
)

var Blog=[]model.Blog{
	{
		ID: "1",
		Title: "Adventures of Sherlock Holmes",
		Author: "Arthur Conan Doyle",
		Body: "The Adventures of Sherlock Holmes is a collection of short stories by British writer Arthur Conan Doyle, first published on 14 October 1892. It contains the earliest short stories featuring the consulting detective Sherlock Holmes, which had been published in twelve monthly issues of The Strand Magazine from July 1891 to June 1892",
		CreatedAt: time.Now(),
	},
}
