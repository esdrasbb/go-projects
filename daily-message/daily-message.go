package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/esdras.barreto/verses"
)

type Config struct {
	Username string
	Password string
	URL      string
}

var c Config

func doRequest() ([]byte, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	req.SetBasicAuth(c.Username, c.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func getRandomNumber(limit int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(limit) + 1
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div><form action=\"daily-message\"><input type=\"submit\" value=\"Pegue sua mensagem\" /></form></div>")
}

func handlerDaily(w http.ResponseWriter, r *http.Request) {
	body, err := doRequest()
	if err != nil {
		fmt.Fprintf(w, "<div>Error! <br> %s</div>", err)
	} else {
		v := verses.DailyMessage{}
		err := json.Unmarshal(body, &v)
		if err != nil {
			fmt.Fprintf(w, "<div>Error parse! <br> %s</div>", err)
		} else {
			fmt.Fprintf(w, "<div>%s</div>", v.Response.Texts[getRandomNumber(36)])
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		os.Exit(1)
	}

	c = Config{}

	err = json.Unmarshal(dat, &c)
	if err != nil {
		os.Exit(1)
	}

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/daily-message", handlerDaily)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
