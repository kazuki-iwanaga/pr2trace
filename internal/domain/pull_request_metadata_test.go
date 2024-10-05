package domain

import (
	"reflect"
	"testing"
)

func TestNewPullRequestMetadata(t *testing.T) {
	type args struct {
		title string
	}

	tests := []struct {
		name string
		args args
		want *PullRequestMetadata
	}{
		{
			name: "NewPullRequestMetadata",
			args: args{
				title: "This is a test tile.",
			},
			want: &PullRequestMetadata{
				title: "This is a test tile.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPullRequestMetadata(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullRequestMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PullRequestMetadata_Title(t *testing.T) {
	type fields struct {
		title string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "PullRequestMetadata_Title",
			fields: fields{
				title: "This is a test tile.",
			},
			want: "This is a test tile.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prm := &PullRequestMetadata{
				title: tt.fields.title,
			}
			if got := prm.Title(); got != tt.want {
				t.Errorf("PullRequestMetadata.Title() = %v, want %v", got, tt.want)
			}
		})
	}
}
