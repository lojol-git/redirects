// File: redirects/redirects.go

package redirects

import (
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

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

var redirectsList Redirects

func (r *Redirects) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redirects) Run(w http.ResponseWriter, req *http.Request) bool {
	for _, redirect := range redirectsList.Redirects {
		var match bool
		var err error

		if strings.Contains(redirect.From, "*") {
			// Wildcard matching
			match, err = path.Match(redirect.From, req.URL.Path)
			if match {
				to := strings.Replace(req.URL.Path, strings.TrimSuffix(redirect.From, "*"), redirect.To, 1)
				to = strings.TrimSuffix(to, "/") // Remove trailing slash
				log.Printf("Redirecting from %s to %s", req.URL.Path, to)
				http.Redirect(w, req, to, redirect.Status)
				return true
			}
		} else {
			// Exact matching
			re, err := regexp.Compile("^" + redirect.From + "$")
			if err != nil {
				log.Printf("Error compiling regex: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return false
			}
			match = re.MatchString(req.URL.Path)
			if match {
				matches := re.FindStringSubmatch(req.URL.Path)
				to := redirect.To
				for i, match := range matches[1:] {
					to = strings.Replace(to, "$"+strconv.Itoa(i+1), match, -1)
				}
				to = strings.TrimSuffix(to, "/") // Remove trailing slash
				log.Printf("Redirecting from %s to %s", req.URL.Path, to)
				http.Redirect(w, req, to, redirect.Status)
				return true
			}
		}

		if err != nil {
			log.Printf("Error matching redirect: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}
	}
	return false
}
