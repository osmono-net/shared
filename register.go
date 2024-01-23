package shared

type RegisterAgentBody struct {
	HwId    string
	Name    string
	Version WindowsVersion
	SiteId  string
}

type WindowsVersion struct {
	CurrentVersion string
	ProductName    string
	MajorVersion   uint64
	MinorVersion   uint64
	CurrentBuild   string
}
