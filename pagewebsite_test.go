package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const (
	testPageWebsiteGet  string = "test_data/pagewebsite_get.json"
	testPageWebsiteList string = "test_data/pagewebsite_list.json"
)

func TestPageWebsiteService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPageWebsiteGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PageWebsite, 1)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name            string
		file            string
		id              int
		opts            []Option
		wantPageWebsite *PageWebsite
		wantErr         error
	}{
		{"Valid response", testPageWebsiteGet, 777777, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 777777, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 777777, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PageWebsites.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPageWebsite) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPageWebsite)
			}
		})
	}
}

func TestPageWebsiteService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPageWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PageWebsite, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name             string
		file             string
		ids              []int
		opts             []Option
		wantPageWebsites []*PageWebsite
		wantErr          error
	}{
		{"Valid response", testPageWebsiteList, []int{1111}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1111}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1111}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PageWebsites.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPageWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPageWebsites)
			}
		})
	}
}

func TestPageWebsiteService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPageWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PageWebsite, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name             string
		file             string
		opts             []Option
		wantPageWebsites []*PageWebsite
		wantErr          error
	}{
		{"Valid response", testPageWebsiteList, []Option{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PageWebsites.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPageWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPageWebsites)
			}
		})
	}
}

func TestPageWebsiteService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(-99999)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.PageWebsites.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPageWebsiteService_Fields(t *testing.T) {
	var tests = []struct {
		name       string
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"Asterisk", `["*"]`, []string{"*"}, nil},
		{"Empty response", "", nil, errInvalidJSON},
		{"No results", "[]", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			fields, err := c.PageWebsites.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
