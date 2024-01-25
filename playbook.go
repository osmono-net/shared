package shared

type PlaybookCommandType int

const (
	ScriptLiteral PlaybookCommandType = iota
	URL
	LocalDiskPath
)

func (cType PlaybookCommandType) String() string {
	switch cType {
	case ScriptLiteral:
		return "Script Literal"
	case URL:
		return "URL"
	case LocalDiskPath:
		return "Local Disk Path"
	default:
		return "Unknown"
	}
}

type PlaybookCommand struct {
	Type       PlaybookCommandType
	ScriptData string
}
