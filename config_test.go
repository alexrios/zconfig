package zconfig

import (
	"os"
	"testing"
)

type Dummy struct {
	Outer struct {
		InnerA string `yaml:"innera"`
		InnerB string `yaml:"innerb"`
	} `yaml:"outer"`
}

func TestLoad(t *testing.T) {
	t.Run("single yaml", func(t *testing.T) {
		dummy := &Dummy{}
		err := Load(dummy, "Dummy.yaml")
		if err != nil {
			t.FailNow()
		}
		if dummy.Outer.InnerA != "hello" {
			t.Fatal("InnerA should be ", "hello")
		}
		if dummy.Outer.InnerB != "world" {
			t.Fatal("InnerB should be ", "world")
		}
	})
	t.Run("multiple yamls", func(t *testing.T) {
		dummy := &Dummy{}
		err := Load(dummy, "Dummy1.yaml", "Dummy2.yaml")
		if err != nil {
			t.FailNow()
		}
		if dummy.Outer.InnerA != "hello" {
			t.Fatal("InnerA should be ", "hello")
		}
		if dummy.Outer.InnerB != "world" {
			t.Fatal("InnerB should be ", "world")
		}
	})
	t.Run("multiple yamls overrided by env vars", func(t *testing.T) {
		dummy := &Dummy{}
		err := os.Setenv("OUTER_INNERA", "bye bye")
		if err != nil {
			t.Fatal(err)
		}
		err = os.Setenv("OUTER_INNERB", "universe")
		if err != nil {
			t.Fatal(err)
		}
		err = Load(dummy, "Dummy1.yaml", "Dummy2.yaml")
		if err != nil {
			t.FailNow()
		}
		if dummy.Outer.InnerA != "bye bye" {
			t.Fatal("InnerA should be", "bye bye", "but got", dummy.Outer.InnerA)
		}
		if dummy.Outer.InnerB != "universe" {
			t.Fatal("InnerB should be", "universe", "but got", dummy.Outer.InnerB)
		}
	})
}
