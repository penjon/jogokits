package logs

type Level int8
const (
	ERROR Level = -1
	DEBUG Level = 0
	INFO Level = 1
	FATAL Level = 2
	WARN Level = 3
)