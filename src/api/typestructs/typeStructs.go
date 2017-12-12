package typestructs

// author structs
type SearchAuthorID struct {
	ID string `json:"id"`
}

type GetAuthorData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BestBook  string `json:"best_book"`
	AliasName string `json:"alias_name"`
}

// book structs
type SearchBookID struct {
	ID string `json:"id"`
}

type GetBookData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AuthorName  string `json:"author_name"`
	PublishDate string `json:"publish_date"`
}
