package ponzi_test

import (
	"testing"
	"time"

	"github.com/bketelsen/ponzi"
	"github.com/gopheracademy/material/content"
	"github.com/kr/pretty"
)

func TestGet(t *testing.T) {
	p := ponzi.New("http://127.0.0.1:8080", 5*time.Second, 2*time.Second)

	result := &content.CourseListResult{}
	err := p.Get(1, "Course", result)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(result)

	result = &content.CourseListResult{}
	err = p.Get(1, "Course", result)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(result)
}

func TestGetAll(t *testing.T) {
	p := ponzi.New("http://127.0.0.1:8080", 5*time.Second, 20*time.Second)

	result := &content.CourseListResult{}
	err := p.GetAll("Course", result)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(result.Data)
	result = &content.CourseListResult{}
	err = p.GetAll("Course", result)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(result.Data)
}
