package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"math/rand"
	"sort"
	"bytes"
	"bufio"
	"regexp"
	"text/template"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

type Test struct {
	Question string
	Variants []string
	Level int
	Code string
	No int
	SourceNo int
}

type Tests []Test

func (slice Tests) Len() int {
    return len(slice)
}

func (slice Tests) Less(i, j int) bool {
    if slice[i].Level < slice[j].Level {
		return true
	} else if slice[i].Level > slice[j].Level {
		return false
	} else {
		return slice[i].SourceNo < slice[j].SourceNo
	}
}

func (slice Tests) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}


var (
	base = kingpin.Flag("base", "Questions base.").Required().String()
	variant = kingpin.Flag("variant", "Job variant.").Required().Int()
	answers = kingpin.Flag("answers", "Print answers.").Default("false").Bool()
	limit = kingpin.Flag("limit", "Questions total limit.").Default("10").Int()
)

func shuffle(arr []Test) {	
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
// common
	answerRegexp, err := regexp.Compile(`!!$`)
	if err != nil {
		panic(err)
	}
// we need re-execution
	rand.Seed(int64(*variant))

// read base
	if _, err := os.Stat(*base); os.IsNotExist(err) {
		panic(fmt.Sprintf("No such file: %s", *base))
	}

	content, err := ioutil.ReadFile(*base)
    if err != nil {
        panic(fmt.Sprintf("Error in readfile: %v", err))
    }

	var rawBase struct{ Title string; Questions []Test }

	if err := yaml.Unmarshal([]byte(content), &rawBase); err != nil {
		panic(fmt.Sprintf("Error in yaml unmarshall: %s", err))
	}

// fetch questions
	qBase := rawBase.Questions

// fetch at least one per level
	selected := make(map[*Test]bool)
	selectedLevels := make(map[int]bool)
	shuffle(qBase)
	for i, test := range qBase {
		qBase[i].SourceNo = i
		if !selectedLevels[test.Level] {
			selected[&qBase[i]] = true
			selectedLevels[test.Level] = true
		}
	}
// add random tests
	for len(selected) < *limit && len(selected) < len(qBase) {
		i := rand.Intn(len(qBase))
		for selected[&qBase[i]] {
			i = (i + 1) % len(qBase)
		}
		selected[&qBase[i]] = true
	}
// flatten
	selectedQuestions := make(Tests, 0, len(selected))
	for key := range selected {
		selectedQuestions = append(selectedQuestions, *key)
	}
// enumerate them
	sort.Sort(selectedQuestions)
	for key := range selectedQuestions {
		selectedQuestions[key].No = key + 1
	}
	
	if *answers {
		for _, raw := range selectedQuestions {
			fmt.Printf("%d:", raw.No);
			for i, variant := range raw.Variants {
				if answerRegexp.MatchString(variant) {
					fmt.Printf(" %d", i + 1)
				}
			}
			fmt.Println()
		}
	} else {
	
		for i, test := range selectedQuestions {
			for j, _ := range test.Variants {
				selectedQuestions[i].Variants[j] = answerRegexp.ReplaceAllString(selectedQuestions[i].Variants[j], "")
			}
		}

		tmpl, _ := template.ParseFiles("test.tpl")
		var out bytes.Buffer
		writer := bufio.NewWriter(&out)
		tmpl.Execute(writer, struct { Questions []Test; Title string }{ selectedQuestions, fmt.Sprintf("%s №%d", rawBase.Title, *variant) })
		writer.Flush()
	
		reg, err := regexp.Compile(`\n[^\S\n]*\n`)
		if err != nil {
			panic(err)
		}
	
		result := reg.ReplaceAllString(out.String(), "\n")
		
		lstReg, err := regexp.Compile(`lstlisting}[^\S\n]*\n[^\S\n]*`)
		if err != nil {
			panic(err)
		}
		
		fmt.Print(lstReg.ReplaceAllString(result, "lstlisting}\n"))
	}
}