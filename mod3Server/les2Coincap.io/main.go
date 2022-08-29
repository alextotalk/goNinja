package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

type loggerRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_, err := fmt.Fprintf(l.logger, "[%s], %s, %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	if err != nil {
		return nil, err
	}

	return l.next.RoundTrip(r)
}

func main() {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "some_token",
		MaxAge: 300,
	}

	cookies := []*http.Cookie{cookie}
	u := url.URL{Host: "https://academy.golang-ninja.com/"}

	jar, err := cookiejar.New(nil)
	jar.SetCookies(&u, cookies)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("\nCookies", req.Cookies())
			fmt.Println(req.Response.Header.Get("Location"))
			fmt.Println("REDIRECT")

			return nil
		},

		Transport: &loggerRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},

		Timeout: time.Second * 10,

		Jar: jar,
	}

	resp, err := client.Get("https://academy.golang-ninja.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	fmt.Println("Response status", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	//	client := &http.Client{}
	//	fmt.Sprintf("%v", client)
	//	fmt.Printf("%v \n", client)
	//	fmt.Printf("%#v", client)

}
