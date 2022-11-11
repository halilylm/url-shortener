package domain

// Url is the domain
type Url struct {
	ID       string `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

type URLRepository interface {
	Insert(url *Url) (*Url, error)
	GetByID(id string) (*Url, error)
}

type URLUsecase interface {
	Generate(url *Url) (*Url, error)
	Redirect(id string) (*Url, error)
}
