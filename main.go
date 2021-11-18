package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var beVerbs []string
var thisThat []string
var prepositions []string

func init() {
	beVerbs = []string{"be", "is", "am", "are", "was", "were", "being", "been"}
	thisThat = []string{"it", "this", "that", "there"}
	prepositions = []string{"it", "this", "that", "there",
		"about", "above", "across", "after", "against", "along", "among", "around",
		"at", "before", "behind", "below", "beneath", "beside", "between", "beyond",
		"by", "down", "during", "far", "from", "in", "inside", "into", "like", "near",
		"of", "off", "on", "onto", "out", "outside", "over", "past", "since", "through",
		"throughout", "till", "to", "toward", "under", "underneath", "until", "up", "upon",
		"with", "within", "without"}
}

func main() {

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	// colorWhite := "\033[37m"

	files := os.Args[1:]
	if len(files) == 0 {
		// Nothing
	} else {
		for _, arg := range files {
			data, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			text := string(data)

			replace(&text, beVerbs, colorRed)
			replace(&text, prepositions, colorCyan)
			replace(&text, thisThat, colorGreen)

			fmt.Println(text)
		}
	}
}

func replace(t *string, list []string, color string) {
	colorReset := "\033[0m"
	for _, word := range list {
		term := " " + word + " "
		*t = strings.ReplaceAll(*t, term, fmt.Sprint(color, term, colorReset))
	}
}
