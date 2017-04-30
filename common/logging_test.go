package common

import (
	"testing"

	"encoding/json"

	"github.com/nautilus/events"
)

func TestLogging_writerPublishesToLogging(t *testing.T) {
	// a mock event broker we can test with
	broker := events.NewMockEventBroker()
	// the byte string we are going to write
	action := LogPayload{
		Label:   "log label",
		Payload: "hello world",
	}
	payload, err := json.Marshal(action)
	// if something went wrong
	if err != nil {
		//  the test failed
		t.Error(err)
		return
	}

	// create the writer appropriate for this broker
	writer, err := LogWriter(broker, action.Label)
	// if something went wrong
	if err != nil {
		//  the test failed
		t.Error(err)
		return
	}
	if writer == nil {
		t.Error("nil writer was returned")
		return
	}

	// we are expecting messages to be published on the log topioc
	broker.ExpectPublish("log", &events.Action{
		Type:    ActionLogAction,
		Payload: string(payload),
	})

	// attempt to write the byte string to the writer
	l, err := writer.Write([]byte(action.Payload))

	switch {
	// if ther was something wrong
	case err != nil:
		t.Error(err)
	// make sure we wrote everything
	case l != len(payload):
		t.Error("did not write the full bytestring")
	// try to close the broker to check for unsatisfied expectations
	case broker.Close() != nil:
		t.Error("Did not publish action on log topic")
	}
}
