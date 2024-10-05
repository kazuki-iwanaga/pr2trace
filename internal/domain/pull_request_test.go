package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPullRequest(t *testing.T) {
	type args struct {
		id       PullRequestID
		metadata PullRequestMetadata
		events   []*PullRequestEvent
	}

	tests := []struct {
		name string
		args args
		want *PullRequest
	}{
		{
			name: "NewPullRequest",
			args: args{
				id:       PullRequestID{owner: "owner", repo: "repo", number: 1},
				metadata: PullRequestMetadata{title: "title"},
				events: []*PullRequestEvent{
					{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				},
			},
			want: &PullRequest{
				id:       PullRequestID{owner: "owner", repo: "repo", number: 1},
				metadata: PullRequestMetadata{title: "title"},
				events: []*PullRequestEvent{
					{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPullRequest(tt.args.id, tt.args.metadata, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPullRequest_ID(t *testing.T) {
	type fields struct {
		id       PullRequestID
		metadata PullRequestMetadata
		events   []*PullRequestEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   PullRequestID
	}{
		{
			name: "PullRequest_ID",
			fields: fields{
				id: PullRequestID{
					owner:  "owner",
					repo:   "repo",
					number: 1,
				},
				metadata: PullRequestMetadata{
					title: "title",
				},
				events: []*PullRequestEvent{
					{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				},
			},
			want: PullRequestID{
				owner:  "owner",
				repo:   "repo",
				number: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PullRequest{
				id:       tt.fields.id,
				metadata: tt.fields.metadata,
				events:   tt.fields.events,
			}
			if got := pr.ID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullRequest.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPullRequest_Metadata(t *testing.T) {
	type fields struct {
		id       PullRequestID
		metadata PullRequestMetadata
		events   []*PullRequestEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   PullRequestMetadata
	}{
		{
			name: "PullRequest_Metadata",
			fields: fields{
				id: PullRequestID{
					owner:  "owner",
					repo:   "repo",
					number: 1,
				},
				metadata: PullRequestMetadata{
					title: "title",
				},
				events: []*PullRequestEvent{
					{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				},
			},
			want: PullRequestMetadata{
				title: "title",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PullRequest{
				id:       tt.fields.id,
				metadata: tt.fields.metadata,
				events:   tt.fields.events,
			}
			if got := pr.Metadata(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullRequest.Metadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPullRequest_Events(t *testing.T) {
	type fields struct {
		id       PullRequestID
		metadata PullRequestMetadata
		events   []*PullRequestEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   []*PullRequestEvent
	}{
		{
			name: "PullRequest_Events",
			fields: fields{
				id: PullRequestID{
					owner:  "owner",
					repo:   "repo",
					number: 1,
				},
				metadata: PullRequestMetadata{
					title: "title",
				},
				events: []*PullRequestEvent{
					{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				},
			},
			want: []*PullRequestEvent{
				{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PullRequest{
				id:       tt.fields.id,
				metadata: tt.fields.metadata,
				events:   tt.fields.events,
			}
			if got := pr.Events(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullRequest.Events() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPullRequest_SelectEvent(t *testing.T) { // nolint: funlen
	type fields struct {
		id       PullRequestID
		metadata PullRequestMetadata
		events   []*PullRequestEvent
	}

	type args struct {
		eventType PullRequestEventType
		method    PullRequestEventSelectMethod
	}

	commonFields := fields{
		id:       PullRequestID{owner: "owner", repo: "repo", number: 1},
		metadata: PullRequestMetadata{title: "title"},
		events: []*PullRequestEvent{
			{eventType: PullRequestEventTypeCommit, timestamp: time.Date(2020, 12, 29, 0, 0, 0, 0, time.UTC)},
			{eventType: PullRequestEventTypeCommit, timestamp: time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)},
			{eventType: PullRequestEventTypeOpen, timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			{eventType: PullRequestEventTypeCommit, timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
			{eventType: PullRequestEventTypeMerge, timestamp: time.Date(2021, 1, 20, 0, 0, 0, 0, time.UTC)},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PullRequestEvent
		wantErr bool
	}{
		{
			name:   "PullRequest_SelectEvent_First_Success",
			fields: commonFields,
			args: args{
				eventType: PullRequestEventTypeCommit,
				method:    PullRequestEventSelectMethodFirst,
			},
			want: &PullRequestEvent{
				eventType: PullRequestEventTypeCommit,
				timestamp: time.Date(2020, 12, 29, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name:   "PullRequest_SelectEvent_Last_Success",
			fields: commonFields,
			args: args{
				eventType: PullRequestEventTypeCommit,
				method:    PullRequestEventSelectMethodLast,
			},
			want: &PullRequestEvent{
				eventType: PullRequestEventTypeCommit,
				timestamp: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name:   "PullRequest_SelectEvent_NotFoundErr",
			fields: commonFields,
			args: args{
				eventType: PullRequestEventTypeApprove,
				method:    PullRequestEventSelectMethodFirst,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PullRequest{
				id:       tt.fields.id,
				metadata: tt.fields.metadata,
				events:   tt.fields.events,
			}
			got, err := pr.SelectEvent(tt.args.eventType, tt.args.method)

			if (err != nil) != tt.wantErr {
				t.Errorf("PullRequest.SelectEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullRequest.SelectEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
