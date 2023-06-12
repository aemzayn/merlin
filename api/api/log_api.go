// Copyright 2020 The Merlin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/caraml-dev/merlin/log"
	"github.com/caraml-dev/merlin/service"
)

var decoder = schema.NewDecoder()

// LogController controls logs API.
type LogController struct {
	*AppContext
}

// ReadLog parses log requests and fetches logs.
func (l *LogController) ReadLog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var errorMsg string
	var query service.LogQuery
	err := decoder.Decode(&query, r.URL.Query())
	if err != nil {
		errorMsg = fmt.Sprintf("Error while parsing query string: %v", err)
		log.Errorf(errorMsg)
		BadRequest(errorMsg).WriteTo(w)
		return
	}

	logLineCh := make(chan string)
	stopCh := make(chan struct{})

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(stopCh)
				return
			case logLine := <-logLineCh:
				// Write to the ResponseWriter
				_, err := w.Write(([]byte(logLine)))
				if err != nil {
					InternalServerError(err.Error()).WriteTo(w)
					return
				}

				// Send the response over network
				// although it's not guaranteed to reach client if it sits behind proxy
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			}
		}
	}()

	if err := l.LogService.StreamLogs(ctx, logLineCh, stopCh, &query); err != nil {
		errorMsg = fmt.Sprintf("Error while streaming logs: %v", err)
		log.Errorf(errorMsg)
		InternalServerError(errorMsg).WriteTo(w)
		return
	}
}
