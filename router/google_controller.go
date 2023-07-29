package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/andrewarrow/feedback/network"
	"github.com/andrewarrow/feedback/util"
)

func handleGoogle(c *Context, second, third string) {
	if second == "" && third == "" && c.Method == "POST" {
		handleGoogleRedirect(c)
		return
	}
	c.NotFound = true
}

func writeGoogleFile(file string) string {
	// https://developers.google.com/identity/gsi/web/guides/verify-google-id-token
	// cache-control: public, max-age=18702, must-revalidate, no-transform
	jsonString, code := network.GetTo("https://www.googleapis.com/oauth2/v3/certs", "")
	if code == 200 {
		ioutil.WriteFile(file, []byte(jsonString), 0644)
		return jsonString
	}
	return ""
}
func getGoogleCerts() string {
	file := "/certs/google_jwt_oauth.json"
	fileInfo, err := os.Stat(file)
	if err != nil {
		return writeGoogleFile(file)
	}
	lastModified := fileInfo.ModTime().Unix()
	if time.Now().Unix()-lastModified > 86400 {
		return writeGoogleFile(file)
	}
	b, _ := ioutil.ReadFile(file)
	return string(b)
}

func handleGoogleRedirect(c *Context) {
	googleCerts := getGoogleCerts()
	c.ReadFormValuesIntoParams("credential")
	credential := c.Params["credential"].(string)
	jwksJSON := json.RawMessage(googleCerts)
	returnPath := "/sessions/new"

	jwks, err := keyfunc.NewJSON(jwksJSON)
	if err != nil {
		SetFlash(c, err.Error())
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}
	token, err := jwt.Parse(credential, jwks.Keyfunc)
	if err != nil {
		SetFlash(c, err.Error())
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	if !token.Valid {
		SetFlash(c, "token not valid")
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Params["email"] = claims["email"]
	c.Params["username"] = claims["email"]
	c.Params["first_name"] = claims["given_name"]
	c.Params["last_name"] = claims["family_name"]
	c.Params["password"] = "google"
	c.ValidateCreate("user")
	m := c.Insert("user")
	if m != "" {
		delete(c.Params, "password")
		c.Update("user", "where email=", claims["email"])
	}
	row := c.One("user", "where email=$1", claims["email"])

	guid := util.PseudoUuid()
	c.Params = map[string]any{"guid": guid, "user_id": row["id"].(int64)}
	c.Insert("cookie_token")
	SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))

	returnPath = "/"
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}
