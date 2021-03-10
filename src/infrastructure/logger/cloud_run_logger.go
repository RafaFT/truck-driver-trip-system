/*
This logger is suitable for Cloud Run.

This struct does not make use of Google's Cloud Logging API, because the library
fails to add resource information on the Log Entry that allows it to be easily
associated with cloud run Request Logs.

The API also does not make very clear how to trigger Error Reporting from
ordinary Log Entry's.

Useful links:
https://cloud.google.com/run/docs/logging#viewing-logs-cloud-logging
https://cloud.google.com/logging/docs/agent/configuration#special-fields
https://cloud.google.com/error-reporting/docs/formatting-error-messages

Issue tracker:
https://issuetracker.google.com/issues/172694908
*/

package logger

import (
	"encoding/json"
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type logEntry struct {
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"`
	Type     string `json:"@type,omitempty"`
}

func NewCloudRunLogger(GCPProject, GCPTrace string) usecase.Logger {
	var logger logEntry
	if GCPProject != "" && GCPTrace != "" {
		logger.Trace = fmt.Sprintf("projects/%s/traces/%s", GCPProject, GCPTrace)
	}

	return logger
}

func (e logEntry) String() string {
	b, _ := json.Marshal(e)

	return string(b)
}

func (e logEntry) Debug(msg string) {
	e.Severity = "DEBUG"
	e.Message = msg
	fmt.Println(e)
}

func (e logEntry) Info(msg string) {
	e.Severity = "INFO"
	e.Message = msg
	fmt.Println(e)
}

func (e logEntry) Warning(msg string) {
	e.Severity = "WARNING"
	e.Message = msg
	fmt.Println(e)
}

func (e logEntry) Error(msg string) {
	e.Severity = "ERROR"
	e.Message = msg
	// e.Type = "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent"
	fmt.Println(e)
}
