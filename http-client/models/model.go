package models

type Result struct {
	TimeMicroSeconds []int64 `json:"time_micro_seconds"`
	Protocol         string  `json:"protocol"`
}
