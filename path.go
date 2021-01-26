package eldis

import (
	"net/url"
	"path"
)

const (
	// methodLogin URL метода Login
	methodLogin = "/api/v2/users/login"
	// methodLogout URL метода Logout
	methodLogout = "/api/v2/users/logout"
	// methodListForDevelopment URL метода ListForDevelopment
	methodListForDevelopment = "/api/v2/tv/listForDevelopment"
	// methodUOMList URL метода UOM/List
	methodUOMList = "/api/v2/uom/list"
	// methodDataNormalized URL метода Data/Normalized
	methodDataNormalized = "/api/v2/data/normalized"
	// methodRawData URL метода Data/RawData
	methodRawData = "/api/v2/data/rawData"
)

// join возвращает полный URI метода API
func join(baseURL, method string) (string, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return method, err
	}

	u.Path = path.Join(u.Path, method)

	return u.String(), nil
}
