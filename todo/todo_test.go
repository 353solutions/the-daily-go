package todo

import (
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	content := "Feed cat"
	due := time.Now()

	td := NewTodo(content, due)
	if td.Due != due {
		t.Fatalf("bad due: got %v, expected %v", td.Due, due)
	}

	if td.Content != content {
		t.Fatalf("bad content: got %v, expected %v", td.Content, content)
	}

	if len(td.Key) == 0 {
		t.Fatalf("empty key")
	}

	if td.Done {
		t.Fatalf("new todo is done")
	}

}
