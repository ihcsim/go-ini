package ini

import "testing"

func TestParse(t *testing.T) {
	var tests = []struct {
		input    string
		expected map[string]map[string]string
	}{
		{input: `[section1]
var1=val1
[section2]
var2=val2`,
			expected: map[string]map[string]string{
				"section1": map[string]string{"var1": "val1"},
				"section2": map[string]string{"var2": "val2"},
			},
		},
		{input: `[section1]
var1=val1
[section2]
var1=val1`,
			expected: map[string]map[string]string{
				"section1": map[string]string{"var1": "val1"},
				"section2": map[string]string{"var1": "val1"},
			},
		},
		{input: `[section1]
var1=val1
var2=val2

[section2]
var1=val1
var2=val2`,
			expected: map[string]map[string]string{
				"section1": map[string]string{"var1": "val1", "var2": "val2"},
				"section2": map[string]string{"var1": "val1", "var2": "val2"},
			},
		},
		{input: `[section1]

var1=val1

var2=val2

[section2]

var1=val1

var2=val2

`,
			expected: map[string]map[string]string{
				"section1": map[string]string{"var1": "val1", "var2": "val2"},
				"section2": map[string]string{"var1": "val1", "var2": "val2"},
			},
		},
		{input: `
# first comment
[section1]
var1=val1

# second comment
[section2]
var1=val1`,
			expected: map[string]map[string]string{
				"section1": map[string]string{"var1": "val1"},
				"section2": map[string]string{"var1": "val1"},
			},
		},

		{input: `[section1]
var1=val1 # first comment
var2=val2 # second comment

[section2]
var3=val3 ; third comment
var4=val4 ; fourth comment`,
			expected: map[string]map[string]string{
				"section1": map[string]string{"var1": "val1", "var2": "val2"},
				"section2": map[string]string{"var3": "val3", "var4": "val4"},
			},
		},
		{input: "", expected: nil},
	}

	for _, test := range tests {
		actual := Parse(test.input)
		for section, expectedBody := range test.expected {
			actualBody, exist := actual[section]
			if !exist {
				t.Error("Missing expected section: ", section)
			}

			for k, v := range expectedBody {
				if v != actualBody[k] {
					t.Errorf("Mismatched property/value pair. Bad input:%q\nExpected value of %s.%s to be %q, but got %q", test.input, section, k, v, actualBody[k])
				}
			}
		}
	}
}
