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
