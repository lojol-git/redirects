package redirects

import (
	"net/http"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Redirects struct {
	Redirects []Redirect `yaml:"redirects"`
}

type Redirect struct {
	From   string `yaml:"from"`
	To     string `yaml:"to"`
	Status int    `yaml:"status"`
}

var redirects Redirects

func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&redirects)
	if err != nil {
		return err
	}

	return nil
}

func Run(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, redirect := range redirects.Redirects {
			matched, err := regexp.MatchString(redirect.From, r.URL.Path)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if matched {
				to := regexp.MustCompile(redirect.From).ReplaceAllString(r.URL.Path, redirect.To)
				http.Redirect(w, r, to, redirect.Status)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
