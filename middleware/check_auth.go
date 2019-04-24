package middleware

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/orgmatileg/go-blog-rest-api-example/config"
	"github.com/orgmatileg/go-blog-rest-api-example/helper"
)

// CheckAuth func Middleware
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bypassedURLs := []string{
			"auth:POST",
			"contact-us:POST",
			"post:GET",
			"tags:GET",
			"settings:GET",
		}

		// regex untuk hapus /v1/
		re := regexp.MustCompile("/v[0-9]{1}/")
		URLRequest := re.ReplaceAllString(r.RequestURI, "")

		for _, uri := range bypassedURLs {

			// Cek method dan URL yang diperbolehkan akses tanpa token
			urlAndMethod := strings.Split(uri, ":")

			// Hilangkan semua string setelah -> /
			re = regexp.MustCompile("/.*")
			URLRequestWithRegex := re.ReplaceAllString(URLRequest, "")

			// Jika url sama dengan url yang terdaftar & method yang di allow adalah ALL
			// berarti semua method diperbolehkan akses tanpa token
			if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "ALL" {
				next.ServeHTTP(w, r)
				return
			}
			if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "POST" {
				next.ServeHTTP(w, r)
				return
			}
			if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "GET" {
				next.ServeHTTP(w, r)
				return
			}
			if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "PUT" {
				next.ServeHTTP(w, r)
				return
			}
			if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "DELETE" {
				next.ServeHTTP(w, r)
				return
			}
		}

		res := helper.Response{}
		res.Err = errors.New("Unauthorized")
		res.Body.StatusCode = 401
		res.Body.StatusMessage = "Error"
		res.Body.Href = r.RequestURI

		tokenFromRequest := r.Header.Get("Authorization")

		if tokenFromRequest == "" {
			res.Err = errors.New("Unauthorized: Please send your credential!")
			res.ServeJSON(w, r)
			return
		}

		if tokenFromRequest != "" {
			jwtToken := &helper.Token{
				Key:         config.GetJWTKey(),
				TokenString: tokenFromRequest,
			}

			if ok := jwtToken.IsValidToken(); !ok {
				res.Err = errors.New("Invalid or malformed token")
				res.ServeJSON(w, r)
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}
		}

	})
}
