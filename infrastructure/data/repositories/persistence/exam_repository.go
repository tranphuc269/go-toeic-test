package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"fmt"
)

type IExamRepository interface {
	CreateExam(context.Context, *entities.Exam) error
	FindExamById(context.Context, uint) (*entities.Exam, error)
	FindAllExams(ctx context.Context) ([]*entities.Exam, error)
	FindExamsByCreatorId(context.Context, uint) ([]*entities.Exam, error)
	FindExamsByTaskerId(context.Context, uint) ([]*entities.Exam, error)
}

type ExamRepository struct {
}

func (er ExamRepository) FindAllExams(ctx context.Context) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Order("created_at").Find(&exams)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindAllExams] failed to find exams from rdb"),
			OriginalError: nil,
		}
	}
	return exams, nil
}

func (er ExamRepository) CreateExam(ctx context.Context, exam *entities.Exam) error {
	//TODO implement me
	db := repositories.GetConn()
	result := db.Create(exam)
	fmt.Printf("result.Error : %s\n", result.Error)
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.CreateExam] fail to insert Exam to Database"),
			OriginalError: result.Error,
		}
	}
	return nil
}

func (er ExamRepository) FindExamById(ctx context.Context, u uint) (*entities.Exam, error) {
	//TODO implement me
	panic("implement me")
}

func (er ExamRepository) FindExamsByCreatorId(ctx context.Context, u uint) ([]*entities.Exam, error) {
	//TODO implement me
	panic("implement me")
}

func (er ExamRepository) FindExamsByTaskerId(ctx context.Context, u uint) ([]*entities.Exam, error) {
	//TODO implement me
	panic("implement me")
}

func CreateExamRepository() IExamRepository {
	return &ExamRepository{}
}
