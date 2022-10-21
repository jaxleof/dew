package link

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
)

func SubmitCode(contest string, index string, code []byte) {
	Login()
	sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	sp.Prefix = "AC!!! "
	sp.Start()
	sp.FinalMSG = "submission complete\n"
	res, err := me.R().SetFormData(map[string]string{
		"csrf_token":            csrf,
		"action":                "submitSolutionFormSubmitted",
		"contestId":             contest,
		"submittedProblemIndex": index,
		"programTypeId":         "61",
		"source":                string(code),
		"tabSize":               "4",
	}).Post(fmt.Sprintf("https://codeforces.com/contest/%s/submit?csrf_token=%s", contest, csrf))
	if err != nil {
		log.Fatal(err)
	}
	sp.Stop()
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res.Body()))
	if err != nil {
		log.Fatal(err)
	}
	errText := doc.Find(".error.for__source").First().Text()
	if errText == "You have submitted exactly the same code before" {
		log.Fatal("You have submitted exactly the same code before")
	}
	log.Fatal(res.StatusCode())
	if err != nil {
		log.Fatal(err)
	}
}
