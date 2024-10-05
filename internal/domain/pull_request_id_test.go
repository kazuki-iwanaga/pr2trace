package domain

import (
	"reflect"
	"testing"
)

func TestNewPullRequestID(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		number int
	}

	tests := []struct {
		name    string
		args    args
		want    *PullRequestID
		wantErr bool
	}{
		{
			name: "Create a new pull request ID",
			args: args{
				owner:  "owner",
				repo:   "repo",
				number: 1,
			},
			want: &PullRequestID{
				owner:  "owner",
				repo:   "repo",
				number: 1,
			},
			wantErr: false,
		},
		{
			name: "Create a new pull request ID with invalid number",
			args: args{
				owner:  "owner",
				repo:   "repo",
				number: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPullRequestID(tt.args.owner, tt.args.repo, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPullRequestID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PullRequestID_Owner(t *testing.T) {
	type fields struct {
		owner  string
		repo   string
		number int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "PullRequestID_Owner",
			fields: fields{
				owner:  "owner",
				repo:   "repo",
				number: 1,
			},
			want: "owner",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prid := &PullRequestID{
				owner:  tt.fields.owner,
				repo:   tt.fields.repo,
				number: tt.fields.number,
			}
			if got := prid.Owner(); got != tt.want {
				t.Errorf("PullRequestID.Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PullRequestID_Repo(t *testing.T) {
	type fields struct {
		owner  string
		repo   string
		number int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "PullRequestID_Owner",
			fields: fields{
				owner:  "owner",
				repo:   "repo",
				number: 1,
			},
			want: "repo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prid := &PullRequestID{
				owner:  tt.fields.owner,
				repo:   tt.fields.repo,
				number: tt.fields.number,
			}
			if got := prid.Repo(); got != tt.want {
				t.Errorf("PullRequestID.Repo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PullRequestID_Number(t *testing.T) {
	type fields struct {
		owner  string
		repo   string
		number int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "PullRequestID_Owner",
			fields: fields{
				owner:  "owner",
				repo:   "repo",
				number: 1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prid := &PullRequestID{
				owner:  tt.fields.owner,
				repo:   tt.fields.repo,
				number: tt.fields.number,
			}
			if got := prid.Number(); got != tt.want {
				t.Errorf("PullRequestID.Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
