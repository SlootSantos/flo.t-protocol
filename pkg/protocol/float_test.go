package procotol

import (
	"fmt"
	"testing"
)

var testCasesUnitBench = []struct {
	name string
	h    *Headers
	want string
}{
	{
		"simple",
		&Headers{"Simple": "Value"},
		"Simple:Value\n",
	},
	{
		"increased",
		&Headers{"Simple": "Value", "Additional": "more"},
		"Simple:Value\nAdditional:more\n",
	},
	{
		"loaded",
		&Headers{"Simple": "Value", "Additional": "more", "Hevy": "load", "top": "one", "topper": "more"},
		"Simple:Value\nAdditional:more\nHevy:load\ntop:one\ntopper:more\n",
	},
}

func TestHeaders_String(t *testing.T) {
	for _, tt := range testCasesUnitBench {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.String(); got != tt.want {
				t.Errorf("Headers.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHeaderToString(b *testing.B) {
	for _, tt := range testCasesUnitBench {
		b.Run(fmt.Sprintf("header_%s", tt.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.h.String()
			}
		})

	}
}

func TestMessage_String(t *testing.T) {
	type fields struct {
		Version  string
		Resource string
		Headers  Headers
		Body     []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Version:  tt.fields.Version,
				Resource: tt.fields.Resource,
				Headers:  tt.fields.Headers,
				Body:     tt.fields.Body,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Message.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
