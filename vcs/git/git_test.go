package git

import (
	"cdx/versioned"
	"errors"
	"reflect"
	"testing"
)

func TestNewGit(t *testing.T) {
	type args struct {
		app   string
		field versioned.Field
		push  bool
		r     repoF
	}
	tests := []struct {
		name    string
		args    args
		want    *Git
		wantErr bool
	}{
		{
			name: "passed through parameters",
			args: args{
				app:   "app",
				field: 1,
				push:  true,
				r: func() (Repository, error) {
					return &FakeGitRepo{}, nil
				},
			},
			want: &Git{
				app:   "app",
				field: 1,
				r:     &FakeGitRepo{},
				push:  true,
			},
			wantErr: false,
		},
		{
			name: "repo func returning error, passes through err",
			args: args{
				app:   "",
				field: 0,
				push:  false,
				r: func() (Repository, error) {
					return nil, errors.New("something went wrong")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.app, tt.args.field, tt.args.push, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGit_Ready(t *testing.T) {
	type fields struct {
		app   string
		field versioned.Field
		r     Repository
		push  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ready returns true when push is disabled",
			fields: fields{
				app:   "",
				field: 0,
				r:     nil,
				push:  false,
			},
			want: true,
		},
		{
			name: "ready returns true when push is disabled and on origin/master",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{isOnMaster: true},
				push:  true,
			},
			want: true,
		},
		{
			name: "ready returns true when push is enabled and on origin/master",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{isOnMaster: true},
				push:  true,
			},
			want: true,
		},
		{
			name: "ready returns false when push is enabled and not origin/master",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{isOnMaster: false},
				push:  true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Git{
				app:   tt.fields.app,
				field: tt.fields.field,
				r:     tt.fields.r,
				push:  tt.fields.push,
			}
			if got := g.Ready(); got != tt.want {
				t.Errorf("Ready() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGit_Release(t *testing.T) {
	type fields struct {
		app   string
		field versioned.Field
		r     Repository
		push  bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "release passes through fields to release",
			fields: fields{
				app:   "app",
				field: 1,
				r:     &FakeGitRepo{},
				push:  false,
			},
			wantErr: false,
		},
		{
			name: "release passes through repo Increment error",
			fields: fields{
				app:   "app",
				field: 1,
				r:     &FakeGitRepo{incrementTagErr: errors.New("something went wrong")},
				push:  false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Git{
				app:   tt.fields.app,
				field: tt.fields.field,
				r:     tt.fields.r,
				push:  tt.fields.push,
			}
			err := g.Release()
			if app, field := g.r.(*FakeGitRepo).passedIncrementTag(); app != tt.fields.app || field != tt.fields.field {
				t.Errorf("Release() passed in: %s & %v, want %s & %v", tt.fields.app, tt.fields.field, app, field)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Release() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGit_Promote(t *testing.T) {
	type fields struct {
		app   string
		field versioned.Field
		r     Repository
		push  bool
	}
	tests := []struct {
		name    string
		stage   string
		fields  fields
		wantErr bool
	}{
		{
			name: "promote passes through fields to promote",
			fields: fields{
				app:   "app",
				field: 1,
				r:     &FakeGitRepo{},
				push:  false,
			},
			stage:   "stage",
			wantErr: false,
		},
		{
			name: "promote passes through repo Promote error",
			fields: fields{
				app:   "app",
				field: 1,
				r:     &FakeGitRepo{promoteErr: errors.New("something went wrong")},
				push:  false,
			},
			stage:   "stage",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Git{
				app:   tt.fields.app,
				field: tt.fields.field,
				r:     tt.fields.r,
				push:  tt.fields.push,
			}
			err := g.Promote(tt.stage)
			if app, stage := g.r.(*FakeGitRepo).passedPromote(); app != tt.fields.app || stage != tt.stage {
				t.Errorf("Release() passed in: %s & %v, want %s & %v", tt.fields.app, tt.stage, app, stage)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Release() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGit_Distribute(t *testing.T) {
	type fields struct {
		app   string
		field versioned.Field
		r     Repository
		push  bool
	}
	tests := []struct {
		name     string
		fields   fields
		wantPush bool
		wantErr  bool
	}{
		{
			name: "does not call push if push flag is false",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{},
				push:  false,
			},
			wantPush: false,
			wantErr:  false,
		},
		{
			name: "pushes if push flag is true",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{},
				push:  true,
			},
			wantPush: true,
			wantErr:  false,
		},
		{
			name: "does not return push error if push flag is false",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{pushTagsErr: errors.New("something went wrong")},
				push:  false,
			},
			wantPush: false,
			wantErr:  false,
		},
		{
			name: "passes through push error",
			fields: fields{
				app:   "",
				field: 0,
				r:     &FakeGitRepo{pushTagsErr: errors.New("something went wrong")},
				push:  true,
			},
			wantPush: true,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Git{
				app:   tt.fields.app,
				field: tt.fields.field,
				r:     tt.fields.r,
				push:  tt.fields.push,
			}
			err := g.Distribute()
			if tt.wantPush != g.r.(*FakeGitRepo).pushed {
				t.Errorf("Release() got push %t, want %t", g.r.(*FakeGitRepo).pushed, tt.wantPush)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Distribute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
