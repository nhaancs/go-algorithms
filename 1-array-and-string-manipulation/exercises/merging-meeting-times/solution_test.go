package merging_meeting_times_test

import (
	merging_meeting_times "github.com/nhaancs/go-algorithms/1-array-and-string-manipulation/exercises/merging-meeting-times"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeRanges(t *testing.T) {
	testCases := []struct {
		name     string
		input    []merging_meeting_times.Meeting
		expected []merging_meeting_times.Meeting
	}{
		{
			"meetings overlap",
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 3},
				{StartTime: 2, EndTime: 4},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 4},
			},
		},
		{
			"meetings touch",
			[]merging_meeting_times.Meeting{
				{StartTime: 5, EndTime: 6},
				{StartTime: 6, EndTime: 8},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 5, EndTime: 8},
			},
		},
		{
			"meeting contains other meeting",
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 8},
				{StartTime: 2, EndTime: 5},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 8},
			},
		},
		{
			"meetings stay separate",
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 3},
				{StartTime: 4, EndTime: 8},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 3},
				{StartTime: 4, EndTime: 8},
			},
		},
		{
			"multiple merged meetings",
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 4},
				{StartTime: 2, EndTime: 5},
				{StartTime: 5, EndTime: 8},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 8},
			},
		},
		{
			"meetings not sorted",
			[]merging_meeting_times.Meeting{
				{StartTime: 5, EndTime: 8},
				{StartTime: 1, EndTime: 4},
				{StartTime: 6, EndTime: 8},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 4},
				{StartTime: 5, EndTime: 8},
			},
		},
		{
			"one long meeting contains smaller meetings",
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 10},
				{StartTime: 2, EndTime: 5},
				{StartTime: 6, EndTime: 8},
				{StartTime: 9, EndTime: 10},
				{StartTime: 10, EndTime: 12},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 12},
			},
		},
		{
			"sample input",
			[]merging_meeting_times.Meeting{
				{StartTime: 0, EndTime: 1},
				{StartTime: 3, EndTime: 5},
				{StartTime: 4, EndTime: 8},
				{StartTime: 10, EndTime: 12},
				{StartTime: 9, EndTime: 10},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 0, EndTime: 1},
				{StartTime: 3, EndTime: 8},
				{StartTime: 9, EndTime: 12},
			},
		},
		{
			"empty input",
			[]merging_meeting_times.Meeting{},
			[]merging_meeting_times.Meeting{},
		},
		{
			"one meeting",
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 2},
			},
			[]merging_meeting_times.Meeting{
				{StartTime: 1, EndTime: 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := merging_meeting_times.MergeRanges(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
