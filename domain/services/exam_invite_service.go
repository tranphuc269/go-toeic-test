package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/infrastructure/data/repositories/persistence"
)

type IExamInviteService interface {
	AddUserToExam(ctx context.Context, request dtos.AddTakerToExam) error
	RemoveUserToExam(ctx context.Context, ExamID int, UserID int) error
}

type ExamInviteServiceImpl struct {
	eir persistence.IExamInviteRepository
}

func (eis ExamInviteServiceImpl) RemoveUserToExam(ctx context.Context, ExamID int, UserID int) error {
	//TODO implement me
	err := eis.eir.RemoveUserToExam(ctx, ExamID, UserID)
	return err
}

func (eis ExamInviteServiceImpl) AddUserToExam(ctx context.Context, request dtos.AddTakerToExam) error {
	//TODO implement me
	err := eis.eir.AddUserToExam(ctx, request.ToListTakerEntity())
	return err
}

func CreateExamInviteService(eir persistence.IExamInviteRepository) IExamInviteService {
	return &ExamInviteServiceImpl{eir: eir}
}
