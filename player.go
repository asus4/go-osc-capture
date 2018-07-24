package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

// const
var isNumber = regexp.MustCompile(`^[-+]?[0-9]+$`)

func csv2osc(csv []string) (float32, *osc.Message, error) {

	if len(csv) < 2 {
		return 0, nil, errors.New("lack of csv cols")
	}

	time, _ := strconv.ParseFloat(csv[0], 32)
	msg := osc.NewMessage(csv[1])

	args := csv[2:]
	for _, arg := range args {
		if isNumber.Match([]byte(arg)) {
			// Int
			msg.Append(strconv.ParseInt(arg, 10, 32))
		} else {
			n, err := strconv.ParseFloat(arg, 32)
			if err == nil {
				// Float
				msg.Append(float32(n))
			} else {
				// String
				msg.Append(arg)
			}
		}
	}
	return float32(time), msg, nil
}

func playRoutine(client *osc.Client, reader *csv.Reader, isFinished chan bool) error {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()

	for {
		elapsed := float32(time.Now().Sub(startTime)) / float32(time.Second)

		cols, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}

		now, msg, err := csv2osc(cols)
		if err != nil {
			continue
		}
		waitTime := now - elapsed
		if waitTime > 0 {
			time.Sleep(time.Duration(waitTime * float32(time.Second)))
		}
		client.Send(msg)
		log.Println(msg)
	}
	isFinished <- true

	return nil
}

func player(addr string, port uint, path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	log.Println("player start")

	isFinished := make(chan bool)
	client := osc.NewClient(addr, int(port))
	go playRoutine(client, reader, isFinished)

	// Wait for finish
	_ = <-isFinished
	log.Println("player finished")

	return nil
}
