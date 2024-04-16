package main

type Book struct {
    ID               string `json:"id"`
    Title            string `json:"title"`
    Author           string `json:"author"`
    PublishedDate    string `json:"published_date"`
    OriginalLanguage string `json:"original_language"`
}

var books = []*Book{
    {
        ID:               "1",
        Title:            "7 habits of Highly Effective People",
        Author:           "Stephen Covey",
        PublishedDate:    "15/08/1989",
        OriginalLanguage: "English",
    },
}

func listBooks() []*Book {
    return books
}
