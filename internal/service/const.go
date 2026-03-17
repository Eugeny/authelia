package service

const (
	fmtLogServerListening = "Listening for %s connections on '%s' path '%s'"
)

const (
	logFieldService   = "service"
	logFieldFile      = "file"
	logFieldOP        = "op"
	logFieldFrequency = "frequency"

	serviceTypeServer    = "server"
	serviceTypeWatcher   = "watcher"
	serviceTypeSignal    = "signal"
	serviceTypeCollector = "collector"
)
