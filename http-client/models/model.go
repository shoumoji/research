package models

type Result struct {
	Count            int64   `json:"count"`
	TimeMicroSeconds []int64 `json:"time_micro_seconds"`
	Protocol         string  `json:"protocol"`
}
