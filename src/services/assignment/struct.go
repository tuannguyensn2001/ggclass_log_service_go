package assignment

type createAssignmentLogInput struct {
	AssignmentId int    `json:"assignmentId"`
	Action       string `json:"action"`
	UserId       int    `json:"userId"`
}
