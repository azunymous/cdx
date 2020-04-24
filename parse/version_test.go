package parse

import "testing"

func TestVersion(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parses semantic version with single digits",
			args: args{
				tag: "1.2.3",
			},
			want: "1.2.3",
		},
		{
			name: "parses semantic version with multiple digits",
			args: args{
				tag: "11.22.33",
			},
			want: "11.22.33",
		},
		{
			name: "parses semantic version with app name prefix",
			args: args{
				tag: "app-11.22.33",
			},
			want: "11.22.33",
		},
		{
			name: "parses semantic version with app name prefix and stage suffix",
			args: args{
				tag: "app-11.22.33-stage",
			},
			want: "11.22.33",
		},
		{
			name: "returns empty string when no semantic version found",
			args: args{
				tag: "app--suffix",
			},
			want: "",
		},
		{
			name: "returns empty string when malformed semantic version found",
			args: args{
				tag: "1..3",
			},
			want: "",
		},
		{
			name: "returns empty string when no digits in semantic version found",
			args: args{
				tag: "..",
			},
			want: "",
		},
		{
			name: "returns empty string when letters in semantic version found",
			args: args{
				tag: "1.2.c",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Version(tt.args.tag); got != tt.want {
				t.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}
