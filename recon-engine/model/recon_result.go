package model

type ReconResult struct {
	Domain  string
	IPs     []string
	Headers map[string]string
	Server  string
	Notes   []string
}
