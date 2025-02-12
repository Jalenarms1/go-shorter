package db

func (appUrl *AppUrl) Save() error {

	var exists bool
	err := DB.QueryRow("select exists(select 1 from AppUrl where url_code = ?)", appUrl.UrlCode).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = DB.Exec("insert into AppUrl (src_url, url_code) values (?, ?)", appUrl.SrcUrl, appUrl.UrlCode)

		return err

	} else {
		_, err = DB.Exec("update au set au.url_code = ? from AppUrl as au", appUrl.UrlCode)

		return err
	}

}

func GetAppUrl(urlCode string) (*AppUrl, error) {
	var appUrl AppUrl
	err := DB.QueryRow("select id, src_url, url_code from AppUrl where url_code = ?", urlCode).Scan(&appUrl.Id, &appUrl.SrcUrl, &appUrl.UrlCode)

	return &appUrl, err
}
