package helper

type HTTPResponse struct {
	Message    HTTPMessage `json:"message"`
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type HTTPMessage struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

type Meta struct {
}

type Pagination struct {
	CurrentPage        int
	CurrentPageElement int
	TotalPages         int
	TotalElement       int
	CursorStart        string
	CursorEnd          string
}

type Param struct {
	Page   int
	Limit  int
	SortBy []string
}
