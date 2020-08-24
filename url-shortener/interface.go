package urlshortener

type Controller interface {
	Create(targetURL string) (*Redirect, error)
	Lookup(shortURL string) (*Redirect, error)
}
