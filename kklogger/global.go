package kklogger

import "sync"

// severity identifies the sort of log: info, warning etc. The binding to flag.Value
// is handled in klog.go
type Severity int32 // sync/atomic int32

// These constants identify the log levels in order of increasing severity.
// A message written to a high-severity log file is also written to each
// lower-severity log file.
const (
	InfoLog Severity = iota
	WarningLog
	ErrorLog
	FatalLog
	NumSeverity = 4
)

var kkmap sync.Map
