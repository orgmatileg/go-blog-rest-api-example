package middleware

import (
	"net/http"
)

// CheckAuth func Middleware
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// // regex untuk hapus /v1/
		// re := regexp.MustCompile("/v[0-9]{1}/")
		// URLRequest := re.ReplaceAllString(c.Input.URL(), "")

		// for _, uri := range bypassedURLs {

		// 	// Cek method dan URL yang diperbolehkan akses tanpa token
		// 	urlAndMethod := strings.Split(uri, ":")

		// 	// Hilangkan semua string setelah -> /
		// 	re = regexp.MustCompile("/.*")
		// 	URLRequestWithRegex := re.ReplaceAllString(URLRequest, "")

		// 	// Sementara
		// 	if urlAndMethod[0] == URLRequest && urlAndMethod[1] == "POST" {
		// 		return
		// 	}

		// 	// Jika url sama dengan url yang terdaftar & method yang di allow adalah ALL
		// 	// berarti semua method diperbolehkan akses tanpa token
		// 	if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "ALL" {
		// 		return
		// 	}
		// 	if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "POST" {
		// 		return
		// 	}
		// 	if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "GET" {
		// 		return
		// 	}
		// 	if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "PUT" {
		// 		return
		// 	}
		// 	if urlAndMethod[0] == URLRequestWithRegex && urlAndMethod[1] == "DELETE" {
		// 		return
		// 	}
		// }

		// res := &nst.Response{
		// 	StatusCode:    401,
		// 	StatusMessage: "Error",
		// 	Description:   "Unauthorized",
		// 	Href:          c.Input.URI(),
		// }

		// if token := c.Request.URL.Query().Get("token"); token == "" {
		// 	if token = c.Request.Header.Get("Authorization"); token != "" {

		// 		re := regexp.MustCompile("(?i)bearer\\s+")
		// 		refinedToken := re.ReplaceAllString(token, "")
		// 		jwtToken := &helpers.Token{
		// 			Key:         JWTKey,
		// 			TokenString: refinedToken,
		// 		}
		// 		if ok := jwtToken.IsValidToken(); !ok {
		// 			res.Description = "Invalid or malformed Token"
		// 		} else {
		// 			if claims := jwtToken.GetPayload(); claims != nil {
		// 				for k, v := range claims {
		// 					switch v.(type) {
		// 					case string:
		// 						c.Request.Header.Add(k, v.(string))
		// 					case float32:
		// 						c.Request.Header.Add(k, fmt.Sprintf("%.f", v))
		// 					case float64:
		// 						c.Request.Header.Add(k, fmt.Sprintf("%.f", v))
		// 					case int:
		// 						c.Request.Header.Add(k, fmt.Sprintf("%d", v))

		// 					}
		// 				}
		// 			}
		// 			return
		// 		}
		// 	}
		// 	c.Output.SetStatus(401)
		// 	c.Output.JSON(res, false, false)
		// }
		// return

		next.ServeHTTP(w, r)
	})
}
