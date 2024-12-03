package main

import (
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	tests := []struct {
		name        string
		re          *regexp.Regexp
		input       string
		expectMatch bool
	}{
		{
			name:        "Simple match",
			re:          regexp.MustCompile(`^hello$`),
			input:       "hello",
			expectMatch: true,
		},
		// ^ (Caret): This asserts the start of the string. It means that the pattern must match from the very beginning of the string.
		// $ (Dollar sign): This asserts the end of the string. It means that the pattern must match right up to the end of the string.
		{
			name:        "No match",
			re:          regexp.MustCompile(`^hello$`),
			input:       "world",
			expectMatch: false,
		},

		{
			name:        "Simple match only at end",
			re:          regexp.MustCompile(`hello$`),
			input:       "abc test hello",
			expectMatch: true,
		},
		{
			name:        "No match only at end",
			re:          regexp.MustCompile(`hello$`),
			input:       "abc test hello abc",
			expectMatch: false,
		},

		{
			name:        "Simple match only at start",
			re:          regexp.MustCompile(`^hello`),
			input:       "hello abc test",
			expectMatch: true,
		},
		{
			name:        "No match only at start",
			re:          regexp.MustCompile(`^hello`),
			input:       "test hello abc test",
			expectMatch: false,
		},

		{
			name:        "Match digits",
			re:          regexp.MustCompile(`^\d+$`),
			input:       "12345",
			expectMatch: true,
		},
		// +: Means "one or more" of the preceding \d (digit). So, it will match at least one digit, but it can also match many digits.

		{
			name:        "Does not match digits and chars",
			re:          regexp.MustCompile(`^\d+$`),
			input:       "123abc",
			expectMatch: false,
		},

		{
			name:        "Match email",
			re:          regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
			input:       "example@test.com",
			expectMatch: true,
		},
		{
			name:        "Invalid email",
			re:          regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
			input:       "example@test",
			expectMatch: false,
		},

		{
			name:        "Match alphanumeric",
			re:          regexp.MustCompile(`^[a-zA-Z0-9]+$`),
			input:       "abc123",
			expectMatch: true,
		},
		{
			name:        "Does not match alphanumeric",
			re:          regexp.MustCompile(`^[a-zA-Z0-9]+$`),
			input:       "abc 123",
			expectMatch: false,
		},

		{
			name:        "Match special characters",
			re:          regexp.MustCompile(`^[!@#$%^&*()]+$`),
			input:       "!@#$%^&*()",
			expectMatch: true,
		},
		{
			name:        "Does not match special characters",
			re:          regexp.MustCompile(`^[!@#$%^&*()]+$`),
			input:       "abc!@#",
			expectMatch: false,
		},

		{
			name:        "Valid hexadecimal",
			re:          regexp.MustCompile(`^#?([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`),
			input:       "#1a2b3c",
			expectMatch: true,
		},
		// |: The pipe | is a logical OR operator, meaning either the 6-digit option or the 3-digit option can match.
		{
			name:        "Invalid hexadecimal",
			re:          regexp.MustCompile(`^#?([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`),
			input:       "#1a2b3g",
			expectMatch: false,
		},

		{
			name:        "Valid date (YYYY-MM-DD)",
			re:          regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
			input:       "2023-12-31",
			expectMatch: true,
		},
		{
			name:        "Invalid date format",
			re:          regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
			input:       "31-12-2023",
			expectMatch: false,
		},

		{
			name:        "Valid phone number",
			re:          regexp.MustCompile(`^\+?\d{1,3}?[-.\s]?\(?\d{1,4}?\)?[-.\s]?\d{1,4}[-.\s]?\d{1,9}$`),
			input:       "+1 800-555 1234",
			expectMatch: true,
		},
		//The \s in regular expressions is a shorthand character class that matches any whitespace character. Specifically, \s will match:
		//* Spaces
		//* Tabs
		//* Newlines
		//* Carriage returns
		//* Form feeds
		{
			name:        "Invalid phone number",
			re:          regexp.MustCompile(`^\+?\d{1,3}?[-.\s]?\(?\d{1,4}?\)?[-.\s]?\d{1,4}[-.\s]?\d{1,9}$`),
			input:       "555-1234-ABC",
			expectMatch: false,
		},

		{
			name:        "Repeated word match",
			re:          regexp.MustCompile(`^(word)+$`),
			input:       "wordwordword",
			expectMatch: true,
		},
		{
			name:        "Partial word repetition",
			re:          regexp.MustCompile(`^(word)+$`),
			input:       "wordwordtest",
			expectMatch: false,
		},
		{
			name:        "Exact repeated word match (3 times)",
			re:          regexp.MustCompile(`^(word){3}$`), // Matches 'word' repeated exactly 3 times
			input:       "wordwordword",
			expectMatch: true,
		},
		{
			name:        "Exact repeated word match (3 times with spaces)",
			re:          regexp.MustCompile(`(word\s+){2}word`), // Matches 'word' repeated exactly 3 times with spaces between
			input:       "word word word",
			expectMatch: true,
		},
		{
			name:        "Too few repetitions",
			re:          regexp.MustCompile(`^(word){3}$`),
			input:       "wordword",
			expectMatch: false,
		},
		{
			name:        "Too many repetitions",
			re:          regexp.MustCompile(`^(word){3}$`),
			input:       "wordwordwordword",
			expectMatch: false,
		},
		{
			name:        "Different word",
			re:          regexp.MustCompile(`^(word){3}$`),
			input:       "wordtestwordword",
			expectMatch: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			match := test.re.MatchString(test.input)
			if match != test.expectMatch {
				t.Errorf("Name: %s .Pattern %q with input %q: expected match %v, got %v",
					test.name, test.re, test.input, test.expectMatch, match,
				)
			}
		})
	}
}
