package services

import (
	"github.com/siam-vit/bughunt-be/internal/database"
	"github.com/siam-vit/bughunt-be/internal/models"
)

func CreateQuestion(question models.Question) error {
	db := database.DB.Db

	_, err := db.Exec(`
		INSERT INTO questions ( question, set, language, difficulty)
		VALUES ($1, $2, $3, $4)`,
		question.Question, question.Set, question.Language, question.Difficulty)
	return err
}

func GetQuestionsByDifficulty(difficulty string, language string) ([]models.Question, error) {
	db := database.DB.Db
	var questionsMap = make(map[uint]*models.Question)
	var questions []models.Question

	rows, err := db.Query(`
		SELECT q.id, q.question, q.set, q.difficulty, q.language,
		       t.id, t.input, t.output, t.question_id
		FROM questions q
		LEFT JOIN test_cases t ON q.id = t.question_id
		WHERE q.difficulty = $1 AND q.language = $2`, difficulty, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var qID uint
		var qText string
		var set uint
		var lang string
		var diff string
		var testCaseID *uint
		var input, output *string
		var tQuestionID *uint

		err := rows.Scan(&qID, &qText, &set, &diff,&lang, &testCaseID, &input, &output, &tQuestionID)
		if err != nil {
			return nil, err
		}

		if _, exists := questionsMap[qID]; !exists {
			questionsMap[qID] = &models.Question{
				ID:         qID,
				Question:   qText,
				Language: lang,
				Set:        set,
				Difficulty: diff,
				TestCases:  []models.TestCases{},
			}
		}

		if testCaseID != nil && input != nil && output != nil && tQuestionID != nil {
			questionsMap[qID].TestCases = append(questionsMap[qID].TestCases, models.TestCases{
				ID:         *testCaseID,
				Input:      *input,
				Output:     *output,
				QuestionID: *tQuestionID,
			})
		}
	}

	for _, q := range questionsMap {
		questions = append(questions, *q)
	}

	return questions, nil
}

func CreateTestCase(testCase models.TestCases) error {
	db := database.DB.Db

	var newTestCaseID int
	err := db.QueryRow(`
		INSERT INTO test_cases (input, output, question_id)
		VALUES ($1, $2, $3) RETURNING id`,
		testCase.Input, testCase.Output, testCase.QuestionID).Scan(&newTestCaseID)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		UPDATE questions
		SET test_case_id = array_append(test_case_id, $1)
		WHERE id = $2`,
		newTestCaseID, testCase.QuestionID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTestCases() ([]models.TestCases, error) {
	db := database.DB.Db
	var testCases []models.TestCases

	rows, err := db.Query(`SELECT id, input, output, question_id FROM test_cases`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tc models.TestCases
		err := rows.Scan(&tc.ID, &tc.Input, &tc.Output, &tc.QuestionID)
		if err != nil {
			return nil, err
		}
		testCases = append(testCases, tc)
	}
	return testCases, nil
}
