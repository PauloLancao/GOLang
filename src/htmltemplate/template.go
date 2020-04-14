package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type user struct {
	Login string
}

type item struct {
	Number    int
	User      user
	Title     string
	CreatedAt time.Time
}

type template1 struct {
	TotalCount int
	Items      []item
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// https://learning.oreilly.com/library/view/the-go-programming/9780134190570/ebook_split_040.html
func main() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	result := template1{5, []item{{100, user{"Testing"}, "Manager", time.Date(1976, 3, 3, 0, 0, 0, 0, time.UTC)}}}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
