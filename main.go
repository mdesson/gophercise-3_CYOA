package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var templateText = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <hr>
    {{if .Options}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Arc}}">{{.text}}</a></li>
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

	fmt.Println(arcs["intro"].Title)

}
