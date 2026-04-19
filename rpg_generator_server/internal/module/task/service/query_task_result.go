package service

func (s *TaskService) QueryTaskResult(taskId uint) (*QueryTaskResultResp, error) {
	taskResult, err := s.TaskResultDao.GetTaskResultById(taskId)
	if err != nil {
		return nil, err
	}
	var resp QueryTaskResultResp
	resp.TaskId = taskId
	resp.Status = taskResult.Status
	if taskResult.Status == "completed" {
		resp.ResultData = taskResult.Dialogues
	}

	return &resp, nil

}
