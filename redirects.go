package redirects

import (
	"log"
	"net/http"
	"os"
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

func Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &redirectsList)
	if err != nil {
		return err
	}

	return nil
}

func Run(w http.ResponseWriter, req *http.Request) bool {
	for _, redirect := range redirectsList.Redirects {

		// Replace "*" with "(.*)" to capture the wildcard parts
		pattern := redirect.From
		if !strings.Contains(pattern, "(.*)") {
			pattern = strings.ReplaceAll(pattern, "*", "([^/]*)")
		}
		// Compile the regex
		re, err := regexp.Compile("^" + pattern + "$")
		if err != nil {
			log.Printf("Error compiling regex: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}

		// Check if the request path matches the regex
		matches := re.FindStringSubmatch(req.URL.Path)
		if matches != nil {
			// Replace "$1", "$2", etc. in the redirect.To with the captured parts
			to := redirect.To
			for i, match := range matches[1:] {
				to = strings.Replace(to, "$"+strconv.Itoa(i+1), match, -1)
			}

			// Remove trailing slash
			to = strings.TrimSuffix(to, "/")

			http.Redirect(w, req, to, redirect.Status)
			return true
		}
	}

	// No matching redirect was found
	return false
}
