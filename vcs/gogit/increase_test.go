package gogit

import (
	"cdx/versioned"
	"testing"
)

func Test_increase(t *testing.T) {
	type args struct {
		latest string
		field  versioned.Field
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "increases patch",
			args: args{
				latest: "0.0.0",
				field:  versioned.Patch,
			},
			want:    "0.0.1",
			wantErr: false,
		},
		{
			name: "increases minor",
			args: args{
				latest: "0.0.0",
				field:  versioned.Minor,
			},
			want:    "0.1.0",
			wantErr: false,
		},
		{
			name: "increases major",
			args: args{
				latest: "0.0.0",
				field:  versioned.Major,
			},
			want:    "1.0.0",
			wantErr: false,
		},

		{
			name: "increases patch 2",
			args: args{
				latest: "0.0.1",
				field:  versioned.Patch,
			},
			want:    "0.0.2",
			wantErr: false,
		},
		{
			name: "increases minor 2",
			args: args{
				latest: "0.1.0",
				field:  versioned.Minor,
			},
			want:    "0.2.0",
			wantErr: false,
		},
		{
			name: "increases major 2",
			args: args{
				latest: "1.0.0",
				field:  versioned.Major,
			},
			want:    "2.0.0",
			wantErr: false,
		},

		{
			name: "increases patch",
			args: args{
				latest: "0.0.0",
				field:  versioned.Patch,
			},
			want:    "0.0.1",
			wantErr: false,
		},
		{
			name: "increases minor",
			args: args{
				latest: "0.0.0",
				field:  versioned.Minor,
			},
			want:    "0.1.0",
			wantErr: false,
		},
		{
			name: "increases major",
			args: args{
				latest: "0.0.0",
				field:  versioned.Major,
			},
			want:    "1.0.0",
			wantErr: false,
		},

		{
			name: "resets patch count minor",
			args: args{
				latest: "0.1.1",
				field:  versioned.Minor,
			},
			want:    "0.2.0",
			wantErr: false,
		},
		{
			name: "resets patch count major",
			args: args{
				latest: "1.0.1",
				field:  versioned.Major,
			},
			want:    "2.0.0",
			wantErr: false,
		},
		{
			name: "resets minor count major",
			args: args{
				latest: "1.1.0",
				field:  versioned.Major,
			},
			want:    "2.0.0",
			wantErr: false,
		},
		{
			name: "resets all count major",
			args: args{
				latest: "1.1.1",
				field:  versioned.Major,
			},
			want:    "2.0.0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := increase(tt.args.latest, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("increase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("increase() got = %v, want %v", got, tt.want)
			}
		})
	}
}
