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
	"runtime"
	"strconv"

	"github.com/rafaft/truck-driver-trip-system/usecase"
)

// usecase Logger implementation
type cloudRunLogger struct {
	trace string
}

type logEntry struct {
	Message        string          `json:"message"`
	Severity       string          `json:"severity"`
	SourceLocation *sourceLocation `json:"logging.googleapis.com/sourceLocation,omitempty"`
	Trace          string          `json:"logging.googleapis.com/trace,omitempty"`
	Type           string          `json:"@type,omitempty"`
}

type sourceLocation struct {
	File     string `json:"file"`
	Line     string `json:"line"`
	Function string `json:"function"`
}

func NewCloudRunLogger(GCPProject, GCPTrace string) usecase.Logger {
	var logger cloudRunLogger
	if GCPProject != "" && GCPTrace != "" {
		logger.trace = fmt.Sprintf("projects/%s/traces/%s", GCPProject, GCPTrace)
	}

	return logger
}

func (e logEntry) String() string {
	b, _ := json.Marshal(e)

	return string(b)
}

func (l cloudRunLogger) Debug(msg string) {
	fmt.Println(logEntry{
		Message:        msg,
		Severity:       "DEBUG",
		SourceLocation: l.getSourceLocationJSON(2),
		Trace:          l.trace,
	})
}

func (l cloudRunLogger) Info(msg string) {
	fmt.Println(logEntry{
		Message:        msg,
		Severity:       "INFO",
		SourceLocation: l.getSourceLocationJSON(2),
		Trace:          l.trace,
	})
}

func (l cloudRunLogger) Warning(msg string) {
	fmt.Println(logEntry{
		Message:        msg,
		Severity:       "WARNING",
		SourceLocation: l.getSourceLocationJSON(2),
		Trace:          l.trace,
	})
}

func (l cloudRunLogger) Error(msg string) {
	fmt.Println(logEntry{
		Message:        msg,
		Severity:       "DEBUG",
		SourceLocation: l.getSourceLocationJSON(2),
		Trace:          l.trace,
		// Type:           "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent",
	})
}

func (l cloudRunLogger) getSourceLocationJSON(skip int) *sourceLocation {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}

	sl := sourceLocation{
		File:     file,
		Line:     strconv.Itoa(line),
		Function: runtime.FuncForPC(pc).Name(),
	}

	return &sl
}
