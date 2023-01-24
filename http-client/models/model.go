package models

type Result struct {
	Count               int64   `json:"count"`
	ConnectSuccessCount int64   `json:"connect_success_count"`
	TimeMicroSeconds    []int64 `json:"time_micro_seconds"`
	Protocol            string  `json:"protocol"`
}
