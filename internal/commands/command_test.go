package commands

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "no command",
			input: "",
			want:  nil,
		}, {
			name:  "no command, just whitespace",
			input: "   ",
			want:  nil,
		}, {
			name:  "simple, no arguments",
			input: "command",
			want:  []string{"command"},
		}, {
			name:  "command with simple argument",
			input: "command argument",
			want:  []string{"command", "argument"},
		}, {
			name:  "command with two simple argument",
			input: "command argument argument2",
			want:  []string{"command", "argument", "argument2"},
		}, {
			name:  "command with two simple arguments and more whitespace inbetween",
			input: "command   argument  argument2  ",
			want:  []string{"command", "argument", "argument2"},
		}, {
			name:  "command with one simple argument and a string",
			input: "command argument \"argument2 is long\"",
			want:  []string{"command", "argument", "argument2 is long"},
		}, {
			name:  "command with one simple argument and single arguments containing escaped quotes",
			input: "command argument \\\"argument2 argument3\\\"",
			want:  []string{"command", "argument", "\"argument2", "argument3\""},
		}, {
			name:  "command with one simple argument and a string containg an escaped quote",
			input: "command argument \"argument2 is \\\" long\"",
			want:  []string{"command", "argument", "argument2 is \" long"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
