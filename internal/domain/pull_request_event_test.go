package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPullRequestEvent(t *testing.T) {
	type args struct {
		eventType PullRequestEventType
		timestamp time.Time
	}

	tests := []struct {
		name string
		args args
		want *PullRequestEvent
	}{
		{
			name: "Create a new pull request event",
			args: args{
				eventType: PullRequestEventTypeCommit,
				timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: &PullRequestEvent{
				eventType: PullRequestEventTypeCommit,
				timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPullRequestEvent(tt.args.eventType, tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullRequestEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PullRequestEvent_EventType(t *testing.T) {
	type fields struct {
		eventType PullRequestEventType
		timestamp time.Time
	}

	tests := []struct {
		name   string
		fields fields
		want   PullRequestEventType
	}{
		{
			name: "PullRequestEvent_EventType",
			fields: fields{
				eventType: PullRequestEventTypeCommit,
				timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: PullRequestEventTypeCommit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &PullRequestEvent{
				eventType: tt.fields.eventType,
				timestamp: tt.fields.timestamp,
			}
			if got := e.EventType(); got != tt.want {
				t.Errorf("PullRequestEvent.EventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PullRequestEvent_Timestamp(t *testing.T) {
	type fields struct {
		eventType PullRequestEventType
		timestamp time.Time
	}

	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "PullRequestEvent_Timestamp",
			fields: fields{
				eventType: PullRequestEventTypeCommit,
				timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &PullRequestEvent{
				eventType: tt.fields.eventType,
				timestamp: tt.fields.timestamp,
			}
			if got := e.Timestamp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullRequestEvent.Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
