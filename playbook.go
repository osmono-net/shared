package shared

type PlaybookCommandType int

const (
	ScriptLiteral PlaybookCommandType = iota
	URL
	LocalDiskPath
)

type PlaybookCommand struct {
	Type       PlaybookCommandType
	ScriptData string
}
