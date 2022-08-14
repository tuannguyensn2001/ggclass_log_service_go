package assignment

import (
	"context"
	"ggclass_log_service/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, assignment models.Assignment) error {
	_, err := r.db.Database("ggclass").Collection(models.Assignment{}.CollectionName()).InsertOne(ctx, assignment)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]models.Assignment, error) {
	var result []models.Assignment

	cursor, err := r.db.Database("ggclass").Collection(models.Assignment{}.CollectionName()).Find(ctx, models.Assignment{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var item models.Assignment
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		result = append(result, item)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindByAssignmentId(ctx context.Context, assignmentId int, userId int) ([]models.Assignment, error) {
	var result []models.Assignment

	cursor, err := r.db.Database("ggclass").Collection(models.Assignment{}.CollectionName()).Find(ctx, models.Assignment{
		AssignmentId: assignmentId,
		UserId:       userId,
	})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var item models.Assignment
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		result = append(result, item)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
