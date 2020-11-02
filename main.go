// Copyright 2020 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/logging"
)

var fruits = []string{
	"banana",
	"apple",
	"strawberry",
	"grape",
	"pineapple",
	"peach",
	"orange",
	"cherry",
}

func main() {
	pid := os.Getenv("PROJECT_ID")
	if pid == "" {
		fmt.Println("Specify PROJECT_ID environment variable")
		os.Exit(1)
	}

	ctx := context.Background()
	cli, err := logging.NewClient(ctx, fmt.Sprintf("projects/%s", pid))
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cli.Close()
		if err != nil {
			panic(err)
		}
	}()

	logger := cli.Logger("test-log")
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	for {
		t := time.Now().In(tokyo)
		msg := fmt.Sprintf("local time is %v", t)

		p := map[string]interface{}{
			"message": msg,
			"fruit":   fruits[rand.Intn(len(fruits))],
		}

		s := logging.Info
		if t.Nanosecond()%7 == 0 {
			s = logging.Warning
		}
		e := logging.Entry{
			Severity: s,
			Payload:  p,
		}
		logger.Log(e)

		time.Sleep(10 * time.Millisecond)
	}
}
