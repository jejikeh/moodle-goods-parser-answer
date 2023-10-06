package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

type Questions []Question

type Question struct {
	AnswerText  string `json:"answerText"`
	RightAnswer string `json:"rightAnswer"`
}

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	subject := strings.ToLower(strings.Replace(strings.Split(req.RequestURI, "?")[0], "/", "", -1))

	answer := strings.Split(req.RequestURI, "=")[1]

	jsonFile, err := os.OpenFile("./data/"+subject+".json", os.O_RDONLY, 0644)
	if err != nil {
		io.WriteString(res, err.Error())
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var questions Questions

	err = json.Unmarshal(byteValue, &questions)
	if err != nil {
		io.WriteString(res, "server2: "+err.Error())
		return
	}

	is_answered := false
	for i, question := range questions {
		if strutil.Similarity(question.AnswerText, answer, metrics.NewHamming()) > 0.25 {
			sDec, _ := base64.StdEncoding.DecodeString(question.AnswerText)
			io.WriteString(res, strconv.Itoa(i)+" QUESTION: "+string(sDec)+"\n\n")
			io.WriteString(res, strconv.Itoa(i)+" answer: "+question.RightAnswer+"\n\n")
			is_answered = true
		}

	}

	if is_answered {
		return
	}

	io.WriteString(res, "жопа! придется использовать поиск\n\n")

	for i, question := range questions {
		if strutil.Similarity(question.AnswerText, answer, metrics.NewHamming()) > 0.10 {
			sDec, _ := base64.StdEncoding.DecodeString(question.AnswerText)
			io.WriteString(res, strconv.Itoa(i)+" QUESTION: "+string(sDec)+"\n\n")
			io.WriteString(res, strconv.Itoa(i)+" answer: "+question.RightAnswer+"\n\n")
			is_answered = true
		}

	}

	io.WriteString(res, "я текст, меня надо любить")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		queryParamDisplayHandler(res, req)
	})

	http.ListenAndServe(":8080", nil)
}
