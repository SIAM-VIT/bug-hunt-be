package models

type TestCases struct {
	ID         uint   `json:"id"`
	Input      string `json:"input"`
	Output     string `json:"output"`
	QuestionID uint   `json:"questionID"`
}

type Question struct {
	ID         uint        `json:"id"`
	Question   string      `json:"question"`
	Language string 	`json:"`
	TestCases  []TestCases `json:"testCaseId"`
	Set        uint        `json:"set"`
	Difficulty string      `json:"difficulty"`
}

