package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type MockHttpDoer struct {
	context string
}

func (m *MockHttpDoer) Do(req *http.Request) (*http.Response, error) {
	//body это текст который наш фейковый HttpDoer возвращает.
	bodyContent := fmt.Sprintf("<html><body><p>%s</p></body></html>", m.context)
	body := ioutil.NopCloser(strings.NewReader(bodyContent))
	return &http.Response{
		StatusCode: 200,
		Status:     http.StatusText(200),
		Body:       body,
	}, nil
}

func NewHttpDoerMock(context string) *MockHttpDoer {
	return &MockHttpDoer{
		context: context,
	}
}

func TestGetURL(t *testing.T) {

	testCases := map[string]struct {
		url         string
		content     string
		wantContent []byte
		wantErr     error
	}{
		"test-1": {
			url:         "local.test1",
			content:     "local test 1",
			wantContent: []byte("<html><body><p>local test 1</p></body></html>"),
			wantErr:     nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			client := NewClient().SetHttp(NewHttpDoerMock(tc.content))
			gotContent, gotErr := client.getURL(tc.url)
			diff := cmp.Diff(tc.wantContent, gotContent)
			if diff != "" {
				t.Fatalf(diff)
			}
			if gotErr != tc.wantErr {
				t.Fatalf("%s: expected err: %v, got err: %v", name, tc.wantErr, gotErr)
			}
		})
	}
}
