package merging_meeting_times

import "sort"

type Meeting struct {
	StartTime int
	EndTime   int
}

func MergeRanges(meetings []Meeting) []Meeting {
	if len(meetings) <= 1 {
		return meetings
	}

	// O(n) Sort the meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].StartTime < meetings[j].StartTime
	})

	mergedMeetings := []Meeting{meetings[0]}

	// O(n)
	for i := 1; i < len(meetings); i++ {
		currentMeeting := meetings[i]
		lastMergedMeeting := &mergedMeetings[len(mergedMeetings)-1]

		if currentMeeting.StartTime <= lastMergedMeeting.EndTime {
			lastMergedMeeting.EndTime = max(lastMergedMeeting.EndTime, currentMeeting.EndTime)
		} else {
			mergedMeetings = append(mergedMeetings, currentMeeting)
		}
	}

	return mergedMeetings
}
