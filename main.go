package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// StoryArc Contains data for one page of choose your own adventure novel
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// StoryHandler HttpHandler for Story
type StoryHandler struct {
	Arcs     map[string]StoryArc
	Template string
}

func (handler StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arcKey := r.URL.Path[1:]

	if arcKey == "" {
		arcKey = "intro"
	}

	arc, ok := handler.Arcs[arcKey]
	if !ok {
		fmt.Fprintf(w, "<p>Invalid URL: Story arc named %v not found</p>", arcKey)
		return
	}
	parsedTemplate := template.Must(template.New(arcKey).Parse(handler.Template))
	if err := parsedTemplate.Execute(w, arc); err != nil {
		log.Fatal(err)
	}
}

var templateText = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Story}}
    <p>{{.}}</p>
    {{end}}
    <hr>
    {{if .Options}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
    </ul>
    {{else}}
    <h3>The End!</h3>
    {{end}}
  </body>
</html>`

func getStoryArcs() map[string]StoryArc {
	var storyArcs map[string]StoryArc

	jsonText, err := ioutil.ReadFile("./story.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonText, &storyArcs)

	return storyArcs
}

func main() {

	arcs := getStoryArcs()
	handler := StoryHandler{arcs, templateText}
	log.Fatal(http.ListenAndServe("localhost:8000", handler))

}
