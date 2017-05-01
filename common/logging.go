package common

import (
	"encoding/json"
	"io"

	"github.com/nautilus/events"
)

type logWriter struct {
	events.EventBroker
	Label string
}

func (w *logWriter) Write(payload []byte) (int, error) {
	// summarize the log with the label
	input, err := json.Marshal(LogPayload{
		Label:   w.Label,
		Payload: string(payload),
	})
	// if something went wrong
	if err != nil {
		// we didn't write anything
		return 0, err
	}

	// attempt to publish the payload with the appropriate label
	err = w.Publish("log", &events.Action{
		Type:    ActionLogAction,
		Payload: string(input),
	})
	// if something went wrong
	if err != nil {
		// we didn't write anything
		return 0, err
	}

	// otherwise we published the full message
	return len(payload), nil
}

// LogWriter takes an event broker and returns a writer that
// publishes whatever is written to the logging service
func LogWriter(broker events.EventBroker, label string) (io.Writer, error) {
	// wrap the broker in a logWriter
	return &logWriter{
		EventBroker: broker,
		Label:       label,
	}, nil
}
