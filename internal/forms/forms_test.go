package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required field when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has fields when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for none-existant field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, bud did not get one")
	}

	postedData := url.Values{}
	postedData.Add("x", "some value")
	form = New(postedData)

	form.MinLength("x", 100)
	if form.Valid() {
		t.Error("form shows minlength of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("x", "abc123")
	form = New(postedData)

	form.MinLength("x", 1)
	if !form.Valid() {
		t.Error("form shows minlength of 1 met when it is")
	}

	isError = form.Errors.Get("x")
	if isError != "" {
		t.Error("should not have an error, bud got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmaild("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existant field")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@test.com")
	form = New(postedData)

	form.IsEmaild("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@test")
	form = New(postedData)

	form.IsEmaild("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}
