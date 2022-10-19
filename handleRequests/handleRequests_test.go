package handleRequests

import (
	"testing"
)

var activeDatabase = make(map[string]interface{})

func TestSet(t *testing.T) {
	got := Set("test", 10, activeDatabase)
	want := 10

	if activeDatabase["test"] == nil {
		t.Errorf("Can not set in database")
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestGet(t *testing.T){
	want := Set("test", 10, activeDatabase)
	got := Get("test", activeDatabase)

	if activeDatabase["test"] == nil {
		t.Errorf("Can not set in database")
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	
}

func TestDel(t *testing.T){
	Set("test", 10, activeDatabase)
	Del("test", activeDatabase)

	if activeDatabase["test"] != nil {
		t.Errorf("Can not delete the value")
	}
}