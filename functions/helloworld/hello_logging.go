// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START functions_log_helloworld]

// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// HelloLogging logs messages.
func HelloLogging(w http.ResponseWriter, r *http.Request) {
	log.WithField("function_name", "HelloLogging").Info("received a new request in HelloLogging")
	fmt.Println("This is stdout")
}

// [END functions_log_helloworld]
