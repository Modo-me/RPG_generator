package service

type QueryTaskResultResp struct {
	Id         uint        `json:"id"`
	TaskId     uint        `json:"task_id"`
	Status     string      `json:"status"`
	ResultData interface{} `json:"result_data,omitempty"`
}
