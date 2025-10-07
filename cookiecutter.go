package cookiecutter

import (
	"fmt"
	"net/http"
)

// Get returns the raw string value of the cookie header
func Get(request *http.Request) string {
	return request.Header.Get("Cookie")
}

// GetValue returns the value of the given cookie name
func GetValue(
	request *http.Request,
	cookieName string,
) string {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

// GetNameValueMap returns the name-value map for all the items in the cookie
func GetNameValueMap(request *http.Request) map[string]string {
	nameValueMap := make(map[string]string)
	for _, cookie := range request.Cookies() {
		nameValueMap[cookie.Name] = cookie.Value
	}
	return nameValueMap
}

// Add adds a new cookie with the name and value
func Add(
	request *http.Request,
	cookieName string,
	cookieValue string,
) {
	request.AddCookie(&http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	})
}

// Remove removes the named cookie if it exists
func Remove(
	request *http.Request,
	cookieName string,
) {
	var cookies []*http.Cookie

	for _, cookie := range request.Cookies() {
		if cookie.Name != cookieName {
			cookies = append(cookies, cookie)
		}
	}

	request.Header.Del("Cookie")

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}
}

// IsNameAvailable checks if the cookie is available
func IsNameAvailable(
	request *http.Request,
	cookieName string,
) bool {
	value := GetValue(request, cookieName)
	return value != ""
}

// NameHasValue checks if the given cookie name has the given cookie value
func NameHasValue(
	request *http.Request,
	cookieName string,
	cookieValue string,
) bool {
	value := GetValue(request, cookieName)
	if value == "" {
		return false
	}
	return value == cookieValue
}

// PrettyPrint prints the name and value tags to the stdout
func PrettyPrint(request *http.Request) {
	for _, cookie := range request.Cookies() {
		fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
	}
}
