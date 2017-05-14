package common

import "testing"

func TestToGlobalID(t *testing.T) {
	// generate a global ID
	id := MarshalGlobalID(&GlobalID{"Project", "1"})

	// make sure we didn't get an empty string back
	if id == "" {
		t.Error("Received empty string when computing global id")
	}
}

func TestFromGlobalID(t *testing.T) {
	// generate a global ID and get the GlobalID object back
	id, err := UnmarshalGlobalID(MarshalGlobalID(&GlobalID{"Project", "1"}))
	if err != nil {
		t.Error(err)
		return
	}

	// make sure we got the right type value
	if id.Type != "Project" {
		// the test failed
		t.Error("Did not receive the correct type when unmarshalling global ID")
	}
	// make sure we got the right ID value
	if id.ID != "1" {
		// the test failed
		t.Error("Did not receive the correct ID when unmarshalling global ID")
	}
}
