package service

import (
	"context"
	"errors"
	"database/sql"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type QuestionService interface {
	FindAll(ctx context.Context) ([]model.GetQuestionResponse, error)
	Create(ctx context.Context, request model.CreateQuestionRequest, createddBy int) (model.GetQuestionResponse, error)
	Delete(ctx context.Context, questionId int) error
	Update(ctx context.Context, request model.UpdateQuestionRequest, questionId int) (model.GetQuestionResponse, error)
	FindByUserId(ctx context.Context, userId int) ([]model.GetQuestionResponse, error)
}

type questionService struct {
	QuestionRepository repository.QuestionRepository
	DB               *sql.DB
}

func NewQuestionService(questionRepository *repository.QuestionRepository, db *sql.DB) QuestionService {
	return &questionService{
		QuestionRepository: *questionRepository,
		DB:               db,
	}
}

func (service *questionService) Create(ctx context.Context, request model.CreateQuestionRequest, userId int) (model.GetQuestionResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetQuestionResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	newQuestion := entity.Questions{
		UserId:				userId,
		ModuleId:  		request.ModuleId,
		Title:  			request.Title,
		Tags:      		request.Tags,
		Description: 	request.Description,
		CreatedAt:   	utils.TimeNow(),
		UpdatedAt:   	utils.TimeNow(),
	}

	question, err := service.QuestionRepository.Create(ctx, tx, newQuestion)
	if err != nil {
		return model.GetQuestionResponse{}, err
	}

	return utils.ToQuestionResponse(question), nil
}


func (service *questionService) FindAll(ctx context.Context) ([]model.GetQuestionResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetQuestionResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.QuestionRepository.FindAll(ctx, tx)
	if err != nil {
		return []model.GetQuestionResponse{}, err
	}

	var courseResponses []model.GetQuestionResponse
	for _, question := range courses {
		courseResponses = append(courseResponses, utils.ToQuestionResponse(question))
	}

	return courseResponses, nil
}


func (service *questionService) Delete(ctx context.Context, questionId int) error {
	userId := 11; // ini nanti akan diubah pakai data auth user-id dari middleware auth
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)
	
	getQuestions, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return err
	}

	if getQuestions.UserId != userId {
		return errors.New("access not allowed")
	}
	
	err = service.QuestionRepository.Delete(ctx, tx, questionId)
	if err != nil {
		return err
	}

	return nil
}

func (service *questionService) Update(ctx context.Context, request model.UpdateQuestionRequest, questionId int) (model.GetQuestionResponse, error) {
	userId := 10; // ini nanti akan diubah pakai data auth user-id dari middleware auth
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetQuestionResponse{}, err
	}
	defer utils.CommitOrRollback(tx)
	
	getQuestions, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return model.GetQuestionResponse{}, err
	}

	if getQuestions.UserId != userId {
		return model.GetQuestionResponse{}, errors.New("access not allowed")
	}

	newQuestion := entity.Questions{
		UserId:				getQuestions.UserId,
		ModuleId:  		request.ModuleId,
		Title:  			request.Title,
		Tags:      		request.Tags,
		Description: 	request.Description,
		CreatedAt:  	getQuestions.CreatedAt,
		UpdatedAt:   	utils.TimeNow(),
	}

	_, err = service.QuestionRepository.Update(ctx, tx, newQuestion, questionId)
	if err != nil {
		return model.GetQuestionResponse{}, err
	}

	getQuestionsUpdate, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return model.GetQuestionResponse{}, err
	}

	return utils.ToQuestionResponse(getQuestionsUpdate), nil
}


func (service *questionService) FindByUserId(ctx context.Context, userId int) ([]model.GetQuestionResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetQuestionResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.QuestionRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		return []model.GetQuestionResponse{}, err
	}

	var courseResponses []model.GetQuestionResponse
	for _, question := range courses {
		courseResponses = append(courseResponses, utils.ToQuestionResponse(question))
	}

	return courseResponses, nil
}