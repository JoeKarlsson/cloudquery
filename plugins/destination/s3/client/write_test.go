package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/s3/client/spec"
	"github.com/cloudquery/filetypes/v4"
	"github.com/google/go-cmp/cmp"
)

func TestSanitizeRawJsonMessage(t *testing.T) {
	testTable := []struct {
		initialArray []byte
		expected     []uint8
	}{
		{
			initialArray: []byte(`{"target": "localhost"}`),
			expected:     []uint8(`{"target":"localhost"}`),
		},
		{
			initialArray: []byte(`{"ta.rget*": "localhost**"}`),
			expected:     []uint8(`{"ta_rget_":"localhost**"}`),
		},
	}
	for _, test := range testTable {
		data := (json.RawMessage)(test.initialArray)
		bArray, err := sanitizeRawJsonMessage(data)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if diff := cmp.Diff(bArray, test.expected); diff != "" {
			t.Errorf("sanitizeRawJsonMessage() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestSanitizeJSONKeys(t *testing.T) {
	m := map[string]any{
		"foo": "bar",
		"bar": map[string]any{
			"foo-bar": "baz",
		},
		"foo:bar":     "baz",
		"foo-bar-baz": []any{"baz", map[string]any{"foo:bar": "baz"}},
		"string": map[string]string{
			"foo-bar": "baz",
		},
		"int": map[string]int{
			"foo-bar": 123,
		},
		"pointer": map[string]*string{
			"foo-bar": &[]string{"baz"}[0],
		},
	}
	sanitizeJSONKeysForObject(m)
	want := map[string]any{
		"foo": "bar",
		"bar": map[string]any{
			"foo_bar": "baz",
		},
		"foo_bar":     "baz",
		"foo_bar_baz": []any{"baz", map[string]any{"foo_bar": "baz"}},
		"string": map[string]string{
			"foo_bar": "baz",
		},
		"int": map[string]int{
			"foo_bar": 123,
		},
		"pointer": map[string]*string{
			"foo_bar": &[]string{"baz"}[0],
		},
	}
	if diff := cmp.Diff(want, m); diff != "" {
		t.Errorf("sanitizeJSONKeys() mismatch (-want +got):\n%s", diff)
	}
}

func TestReplacePathVariables(t *testing.T) {
	cases := []struct {
		inputPath    string
		uuid         string
		tableName    string
		expectedPath string
	}{
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}",
			uuid:         "",
			tableName:    "",
			expectedPath: "test/test",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.json",
			tableName:    "test-table",
			uuid:         "",
			expectedPath: "test/test/test-table/.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.json",
			tableName:    "test-table",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/test-table/FAKE-UUID.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.json",
			tableName:    "",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/FAKE-UUID.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.{{FORMAT}}",
			tableName:    "",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/FAKE-UUID.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/year={{YEAR}}/month={{MONTH}}/day={{DAY}}/hour={{HOUR}}/minute={{MINUTE}}/{{UUID}}.json",
			tableName:    "test-table",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/test-table/year=2021/month=03/day=05/hour=04/minute=01/FAKE-UUID.json",
		},
	}

	tm := time.Date(2021, 3, 5, 4, 1, 2, 3, time.UTC)
	for _, tc := range cases {
		c := &Client{
			spec: &spec.Spec{
				Path:     tc.inputPath,
				FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON},
			},
		}
		if diff := cmp.Diff(tc.expectedPath, c.spec.ReplacePathVariables(tc.tableName, tc.uuid, tm)); diff != "" {
			t.Errorf("unexpected Path Substitution (-want +got):\n%s", diff)
		}
	}
}
