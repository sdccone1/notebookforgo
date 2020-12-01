package test

import (
	"../demo"
	"testing"
)

func TestMyType(t *testing.T) {
	t.Run("", func(t *testing.T) {
		s := &demo.Student1{
			Person: demo.Person{
				"Bob",
			},
		}
		s.Student1Daily()
	})
}
