package gomediacenter

// Helper function for converting between time units and C#'s time type Ticks.
// 10 000 ticks in 1 millisecond.

// TicksToMilliSeconds converts Ticks to milliseconds.
func TicksToMilliSeconds(ticks int64) int64 {
	return ticks / 10000
}

// TicksToNanoSeconds converts Ticks to nanoseconds.
func TicksToNanoSeconds(ticks int64) int64 {
	return ticks * 10
}

// NanoSecondsToTicks converts nanoseconds to ticks.
func NanoSecondsToTicks(ns int64) int64 {
	return ns / 10
}

// MilliSecondsToTicks converts milliseconds to ticks.
func MilliSecondsToTicks(us int64) int64 {
	return us * 10000
}
