package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultFilename = "problems.csv"
const defaultTimerLengthInSec = 30

type question struct {
	prompt string
	answer string
}

func main() {
	fn := getFilename()
	b, err := ioutil.ReadFile(fn)
	checkError(err)
	administerQuiz(string(b))
}

func getFilename() string {
	args := os.Args
	if len(args) < 2 {
		return defaultFilename
	}
	return args[1]
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func administerQuiz(quiz string) {
	questions := getQuestions(quiz)
	administerQuestions(questions)
}

func administerQuestions(questions []question) {
	score := 0
	timer := getTimer()
	answerCh := make(chan string)
	for _, q := range questions {
		ask(q)
		go scanForAnswer(answerCh)
		select {
		case <-timer.C:
			printEndOfQuizMessage(questions, score)
			return
		case ans := <-answerCh:
			if ans == q.answer {
				score++
			}
		}
	}
	printEndOfQuizMessage(questions, score)
}

func getTimerLength() int {
	args := os.Args
	if len(args) < 3 {
		return defaultTimerLengthInSec
	}
	i, err := strconv.Atoi(args[2])
	checkError(err)
	return i
}

func getTimer() *time.Timer {
	lengthInSec := getTimerLength()
	duration := time.Duration(lengthInSec) * time.Second
	return time.NewTimer(duration)
}

func printEndOfQuizMessage(questions []question, score int) {
	fmt.Printf("\nYou answered %v of %v questions correctly!\n", score, len(questions))
}

func getQuestions(quiz string) []question {
	q := csv.NewReader(strings.NewReader(quiz))
	questions := []question{}
	for {
		q, err := q.Read()
		if err == io.EOF {
			break
		}
		checkError(err)
		question := newQuestion(q)
		questions = append(questions, question)
	}
	return questions
}

func newQuestion(q []string) question {
	checkValid(q)
	return question{
		prompt: q[0],
		answer: q[1],
	}
}

func checkValid(q []string) {
	if len(q) != 2 {
		log.Fatal("Invalid question", q)
	}
}

func ask(q question) {
	fmt.Printf("What is %v?\n", q.prompt)
}

func scanForAnswer(c chan (string)) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	checkError(err)
	c <- strings.Trim(text, "\n")
}
