package models

type NodeInfo struct {
	Name        string
	Description string
	ForksCount  int
}

type GetNodeResponse struct {
	Projects map[string][]NodeInfo
}
