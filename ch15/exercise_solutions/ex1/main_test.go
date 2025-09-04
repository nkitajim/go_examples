package main

import (
	"bytes"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_parser(t *testing.T) {
	data := []struct {
		name   string
		in     []byte
		out    Input
		errMsg string
	}{
		{
			name:   "sucess",
			in:     []byte("test\n+\n1\n1"),
			out:    Input{Id: "test", Op: "+", Val1: 1, Val2: 1},
			errMsg: "",
		},
		{
			name:   "empty",
			in:     []byte(""),
			out:    Input{},
			errMsg: "Invalid input",
		},
		{
			name:   "emptyId",
			in:     []byte("\n+\n1\n1"),
			out:    Input{Id: "", Op: "+", Val1: 1, Val2: 1},
			errMsg: "",
		},
		{
			name:   "invalidVal1",
			in:     []byte("1\n+\none\n1"),
			out:    Input{},
			errMsg: `strconv.Atoi: parsing "one": invalid syntax`,
		},
		{
			name:   "invalidVal2",
			in:     []byte("1\n+\n1\none"),
			out:    Input{},
			errMsg: `strconv.Atoi: parsing "one": invalid syntax`,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			input, err := parser(d.in)
			if diff := cmp.Diff(input, d.out); diff != "" {
				t.Errorf("input error: %s", diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if diff := cmp.Diff(errMsg, d.errMsg); diff != "" {
				t.Errorf("err error: %s", diff)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	data := []struct {
		name string
		in   []byte
		out  Result
	}{
		{
			name: "plus",
			in:   []byte("test\n+\n1\n1"),
			out:  Result{Id: "test", Value: 2},
		},
		{
			name: "minus",
			in:   []byte("test\n-\n1\n1"),
			out:  Result{Id: "test", Value: 0},
		},
		{
			name: "mul",
			in:   []byte("test\n*\n1\n1"),
			out:  Result{Id: "test", Value: 1},
		},
		{
			name: "div",
			in:   []byte("test\n/\n1\n1"),
			out:  Result{Id: "test", Value: 1},
		},
		{
			name: "divError",
			in:   []byte("test\n/\n1\n0"),
			out:  Result{},
		},
		{
			name: "invalidOp",
			in:   []byte("test\n?\n1\n1"),
			out:  Result{},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			in := make(chan []byte, 1)
			out := make(chan Result, 1)
			in <- d.in
			close(in)

			DataProcessor(in, out)

			var got Result
			select {
			case got = <-out:
			default:
				got = Result{}
			}
			if diff := cmp.Diff(got, d.out); diff != "" {
				t.Errorf("out error: %s", diff)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	ch := make(chan Result, 1)
	ch <- Result{Id: "a", Value: 1}
	close(ch)

	var buf bytes.Buffer
	WriteData(ch, &buf)

	got := buf.String()
	want := "a:1\n"
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("error %s", diff)
	}
}

func TestNewController(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		out := make(chan []byte, 1)
		handler := NewController(out)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("hello"))
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		if "OK: 1" != string(body) {
			t.Errorf("Could not match body: %s", string(body))
		}
	})
	t.Run("Bad Input", func(t *testing.T) {
		out := make(chan []byte, 1)
		handler := NewController(out)

		badReader := &brokenReader{}
		req := httptest.NewRequest(http.MethodPost, "/", badReader)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		if "Bad Input" != string(body) {
			t.Errorf("Could not match body: %s", string(body))
		}
	})
	t.Run("Too Busy", func(t *testing.T) {
		out := make(chan []byte, 1)
		out <- []byte("already full")
		handler := NewController(out)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("hello"))
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		if "Too Busy: 1" != string(body) {
			t.Errorf("Could not match body: %s", string(body))
		}
	})
	t.Run("Race", func(t *testing.T) {
		out := make(chan []byte, 1)
		handler := NewController(out)

		for i := 1; i <= 1000; i++ {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("hello"))
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)
			res := rec.Result()
			defer res.Body.Close()

			select {
			case <-out:
			}

			body, _ := io.ReadAll(res.Body)
			if fmt.Sprintf("OK: %d", i) != string(body) {
				t.Errorf("Could not match body: %s, i: %d", string(body), i)
			}
		}
	})
}

type brokenReader struct{}

func (b *brokenReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}

func Test_getCounter(t *testing.T) {
	counter := getCounter()
	if counter != 5000 {
		t.Errorf("counter invalid. count = %d\n", counter)
	}
}

func Fuzz_parser(f *testing.F) {
	testcases := [][]byte{
		[]byte("test\n+\n1\n1"),
		[]byte("test\n"),
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		input, err := parser(in)
		if err != nil {
			t.Skip("Invalid data")
		}

		reIn := join(input)
		input2, err := parser(reIn)

		if diff := cmp.Diff(input, input2); diff != "" {
			t.Error(diff)
		}
	})
}

func join(i Input) []byte {
	r := fmt.Sprintf("%s\n%s\n%d\n%d", i.Id, i.Op, i.Val1, i.Val2)
	return []byte(r)
}
