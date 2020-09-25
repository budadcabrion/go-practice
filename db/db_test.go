package db

import (
	"testing"
	"fmt"
)

func init() {
	InitDB("test", true)

	t1 := Thing{0, "Toaster", "Appliance"}
	t2 := Thing{0, "Toe", "Appendage"}
	t3 := Thing{0, "Hamster", "Animal"}

	InsertThing(t1)
	InsertThing(t2)
	InsertThing(t3)
	fmt.Println("init done")
}

func TestListThings(t *testing.T) {
	ts := ListThings()
	fmt.Println(ts)

	if len(ts) != 3 {
		t.Error("len(ts) != 3")
	} else {
		if ts[0].Id != 1 {
			t.Error("ts[0].Id != 1")
		}
		if ts[1].Name != "Toe" {
			t.Error("ts[1].Name != Toe")
		}
		if ts[2].Type != "Animal" {
			t.Error("ts[2].Type != Animal")
		}
	}
}

func TestGetThing(t *testing.T) {
	thing, err := GetThing(3)
	fmt.Println(thing, err)

	if err != nil {
		t.Error(err)
	}
	if thing.Name != "Hamster" {
		t.Error("thing.Name != Hamster")
	}
}

func TestGetThing2(t *testing.T) {
	thing, err := GetThing(4)
	fmt.Println(thing, err)

	if err == nil {
		t.Error("err == nil")
	}
	if thing.Name != "" {
		t.Error("thing.Name != ")
	}
}