package cache

import "testing"

func TestCreateCache(t *testing.T) {
	// Checks to see if the global Cache variable is for some reason set
	if Cache != nil {
		t.Error("Global Cache variable was set before being initalized")
	}

	// Creates cache
	Create()

	// Checks to see if the global Cache variable was set
	if Cache == nil {
		t.Error("Create() failed, expected Cache to have initialized")
	}
}
