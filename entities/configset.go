package entities

type AppType int

const (
	Rest AppType = iota
	GraphQL
)

type ConfigSet struct {
	PackageName string
	AppName     string
	AppType     AppType
}
