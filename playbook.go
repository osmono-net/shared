package shared

type PlaybookCommandType int

const (
	ScriptLiteral PlaybookCommandType = iota
	URL
	LocalDiskPath
)
