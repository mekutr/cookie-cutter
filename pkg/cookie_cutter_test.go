package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testCookieValue = "sessionId=xyz789; _csrf=edc456tgh; __Secure-UserToken=bc9m7tyG4VgXsabcA; trackingId=123abc456def; rememberMe=true; AuthToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c;"

func TestGet(t *testing.T) {
	testRequest := getTestRequest()
	actual := Get(testRequest)
	assert.Equal(t, testCookieValue, actual)
}

func TestGetValue(t *testing.T) {
	testRequest := getTestRequest()
	actual := GetValue(testRequest, "sessionId")
	assert.Equal(t, "xyz789", actual)
}

func TestGetNameValueMap(t *testing.T) {
	testRequest := getTestRequest()

	actual := GetNameValueMap(testRequest)
	expected := map[string]string{
		"sessionId":          "xyz789",
		"_csrf":              "edc456tgh",
		"__Secure-UserToken": "bc9m7tyG4VgXsabcA",
		"trackingId":         "123abc456def",
		"rememberMe":         "true",
		"AuthToken":          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}

	assert.Equal(t, expected, actual)
}

func TestAdd(t *testing.T) {
	cookieName := "testCookieName"
	cookieValue := "testCookieValue"
	testRequest := getTestRequest()

	Add(testRequest, cookieName, cookieValue)

	actual := GetValue(testRequest, cookieName)
	assert.Equal(t, cookieValue, actual)
}

func TestRemove(t *testing.T) {
	testRequest := getTestRequest()

	Remove(testRequest, "AuthToken")

	actual := GetNameValueMap(testRequest)
	expected := map[string]string{
		"sessionId":          "xyz789",
		"_csrf":              "edc456tgh",
		"__Secure-UserToken": "bc9m7tyG4VgXsabcA",
		"trackingId":         "123abc456def",
		"rememberMe":         "true",
	}

	assert.Equal(t, expected, actual)
}

func TestIsNameAvailable(t *testing.T) {
	testRequest := getTestRequest()
	actual := IsNameAvailable(testRequest, "trackingId")
	assert.True(t, actual)
}

func TestIsNameAvailable_NotAvailable(t *testing.T) {
	testRequest := getTestRequest()
	actual := IsNameAvailable(testRequest, "testCookieName")
	assert.False(t, actual)
}

func TestIsNameHasValue(t *testing.T) {
	testRequest := getTestRequest()
	actual := NameHasValue(testRequest, "_csrf", "edc456tgh")
	assert.True(t, actual)
}

func TestIsNameHasValue_ValueMisMatch(t *testing.T) {
	cookieName := "_csrf"
	testRequest := getTestRequest()

	actual := NameHasValue(testRequest, cookieName, "test")

	assert.True(t, IsNameAvailable(testRequest, cookieName))
	assert.False(t, actual)
}

func TestIsNameHasValue_NotAvailable(t *testing.T) {
	cookieName := "csrf"
	testRequest := getTestRequest()

	actual := NameHasValue(testRequest, cookieName, "edc456tgh")

	assert.False(t, IsNameAvailable(testRequest, cookieName))
	assert.False(t, actual)
}

func TestPrettyPrint(t *testing.T) {
	testRequest := getTestRequest()
	PrettyPrint(testRequest)
}

func getTestRequest() *http.Request {
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	request.Header.Set("Cookie", testCookieValue)
	return request
}
