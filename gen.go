package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"math/rand"
	"time"
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
	Level string
	Code string
}

type Tests []Test

func (slice Tests) Len() int {
    return len(slice)
}

func (slice Tests) Less(i, j int) bool {
    return slice[i].Level < slice[j].Level;
}

func (slice Tests) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}


var (
	base = kingpin.Flag("base", "Questions base.").Required().String()
	limit = kingpin.Flag("limit", "Questions total limit.").Default("10").Int()
)

func shuffle(arr []Test) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	if _, err := os.Stat(*base); os.IsNotExist(err) {
		panic(fmt.Sprintf("No such file: %s", *base))
	}

	content, err := ioutil.ReadFile(*base)
    if err != nil {
        panic(fmt.Sprintf("Error in readfile: %v", err))
    }

	var qBase []Test

	if err := yaml.Unmarshal([]byte(content), &qBase); err != nil {
		panic(fmt.Sprintf("Error in yaml unmarshall: %s", err))
	}

	selected := make(map[*Test]bool)
	selectedLevels := make(map[string]bool)
	shuffle(qBase)
	for i, test := range qBase {
		if !selectedLevels[test.Level] {
			selected[&qBase[i]] = true
			selectedLevels[test.Level] = true
		}
	}
	for len(selected) < *limit && len(selected) < len(qBase) {
		i := rand.Int() % len(qBase)
		for selected[&qBase[i]] {
			i = (i + 1) % len(qBase)
		}
		selected[&qBase[i]] = true
	}
	selectedQuestions := make(Tests, 0, len(selected))
	for key := range selected {
		selectedQuestions = append(selectedQuestions, *key)
	}
	sort.Sort(selectedQuestions)

    tmpl, _ := template.ParseFiles("test.tpl")
	var out bytes.Buffer
	writer := bufio.NewWriter(&out)
	tmpl.Execute(writer, struct { Questions []Test }{ selectedQuestions })
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