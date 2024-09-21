# Merging meeting times

## Problem

**Your company built an in-house calendar tool called HiCal. You want to add a feature to see the times in a day when everyone is available.**

To do this, you’ll need to know when any team is having a meeting. In HiCal, a meeting is stored as objects with integer properties startTime and endTime.
These integers represent the number of 30-minute blocks past 9:00am.

For example:

```
{ startTime: 2, endTime: 3 }  // meeting from 10:00 – 10:30 am
{ startTime: 6, endTime: 9 }  // meeting from 12:00 – 1:30 pm
```

Write a function mergeRanges() that takes an array of multiple meeting time ranges and returns an array of condensed ranges.

For example, given:

```
[
    { startTime: 0,  endTime: 1 },
    { startTime: 3,  endTime: 5 },
    { startTime: 4,  endTime: 8 },
    { startTime: 10, endTime: 12 },
    { startTime: 9,  endTime: 10 },
]
```

your function would return:

```
[
    { startTime: 0, endTime: 1 },
    { startTime: 3, endTime: 8 },
    { startTime: 9, endTime: 12 },
]
```

**Do not assume the meetings are in order**. The meeting times are coming from multiple teams.

**Write a solution that's efficient even when we can't put a nice upper bound on the numbers representing our time ranges**. Here we've simplified our times down to the number of 30-minute slots past 9:00 am. But we want the function to work even for very large numbers, like Unix timestamps. In any case, the spirit of the challenge is to merge meetings where startTime and endTime don't have an upper bound.

## Analysis

What if we only had two ranges? Let's take:

```
[{ startTime: 1, endTime: 3 }, { startTime: 2, endTime: 4 }]
```

These meetings clearly overlap, so we should merge them to give:

```
[{ startTime: 1, endTime: 4 }]
```

But how did we know that these meetings overlap?

We could tell the meetings overlapped because the end time of the first one was after the start time of the second one! But our ideas of "first" and "second" are important here—this only works after we ensure that we treat the meeting that starts earlier as the "first" one.

Edge cases:
- The end time of the first meeting and the start time of the second meeting are equal. For example, `{ startTime: 1, endTime: 2 }` and `{ startTime: 2, endTime: 3 }`. These meetings should probably be merged into `{ startTime: 1, endTime: 3 }`, although they don't "overlap".
- The second meeting ends before the first meeting ends. For example, `{ startTime: 1, endTime: 5 }` and `{ startTime: 2, endTime: 3 }`. These meetings should be merged into `{ startTime: 1, endTime: 5 }`.

Here's a formal algorithm:
- We treat the meeting with earlier start time as "first," and the other as "second."
- If the end time of the first meeting is equal to or greater than the start time of the second meeting, we merge the two meetings into one time range. The resulting time range's start time is the first meeting's start, and its end time is the later of the two meetings' end times.
Else, we leave them separate.

So, we could compare every meeting to every other meeting in this way, merging them or leaving them separate.

Comparing all pairs of meetings would be `O(n^2)` time. We can do better!

We could sort our meetings in order of start time. Then, we could walk through the sorted meetings from left to right.

Sorting takes `O(n lg n)` time in the worst case. If we can then do the merging in one pass, that's another O(n) time, for O(n lg n) overall. That's not as good as O(n), but it's better than O(n^2).

## Solution

- [Go implementation](exercise.go)
- [Test cases](exercise_test.go)

## Complexity

O(n lg n) time and O(n) space.

Even though we only walk through our array of meetings once to merge them, we sort all the meetings first, giving us a runtime of O(n lg n). It's worth noting that if our input were sorted, we could skip the sort and do this in O(n) time!

We create a new array of merged meeting times. In the worst case, none of the meetings overlap, giving us an array identical to the input array. Thus we have a worst-case space cost of O(n).

## Bonus

1. What if we did have an upper bound on the input values? Could we improve our runtime? Would it cost us memory?
2. Could we do this "in place" on the input array and save some space? What are the pros and cons of doing this in place?

## What We Learned

This one arguably uses a greedy ↴ approach as well, except this time we had to sort the array first.

How did we figure that out?

We started off trying to solve the problem in one pass, and we noticed that it wouldn't work. We then noticed the reason it wouldn't work: to see if a given meeting can be merged, we have to look at all the other meetings! That's because the order of the meetings is random.

That's what got us thinking: what if the array were sorted? We saw that then a greedy approach would work. We had to spend O(nlgn) time on sorting the array, but it was better than our initial brute force approach, which cost us O(n^2) time!

