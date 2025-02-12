package db

type AppUrl struct {
	Id      string
	SrcUrl  string `db:"src_url"`
	UrlCode string `db:"url_code"`
}
