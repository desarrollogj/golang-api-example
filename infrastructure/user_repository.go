package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/libs/database"
	"github.com/desarrollogj/golang-api-example/libs/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository represents the methods to be implemented by users repositories
type UserRepository interface {
	FindAllActive() ([]domain.User, error)
	FindActiveByReference(reference string) (domain.User, error)
	FindByReference(reference string) (domain.User, error)
	SearchActive(input domain.UserSearchInput) (domain.UserSearchOutput, error)
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(reference string) (domain.User, error)
}

// mongoUserRepository is the MongoDB implementation of UserRepository
type mongoUserRepository struct {
	config domain.MongoRepositoryConfiguration
	mapper UserMongoRepositoryMapper
}

// NewMongoUserRepository creates a new mongoUserRepository
func NewMongoUserRepository(config domain.MongoRepositoryConfiguration, mapper UserMongoRepositoryMapper) mongoUserRepository {
	return mongoUserRepository{
		config: config,
		mapper: mapper,
	}
}

func (r mongoUserRepository) FindAllActive() ([]domain.User, error) {
	client := database.Mongo.Client
	collection := client.Database(r.config.Database).Collection(r.config.UsersCollection)

	users := []MongoUser{}
	cur, err := collection.Find(context.TODO(), bson.D{{Key: "is_active", Value: true}})
	if err != nil {
		errMsg := "unexpected error when find all users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return []domain.User{}, errors.New(errMsg)
	}

	err = cur.All(context.TODO(), &users)
	if err != nil {
		errMsg := "unexpected error when find all users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return []domain.User{}, errors.New(errMsg)
	}

	return r.mapper.MapRepositoryListToDomainList(users), nil
}

func (r mongoUserRepository) FindActiveByReference(reference string) (domain.User, error) {
	user, err := r.findByReference(reference, true)
	if err != nil {
		return domain.User{}, err
	}

	return r.mapper.MapRepositoryToDomain(user), nil
}

func (r mongoUserRepository) FindByReference(reference string) (domain.User, error) {
	user, err := r.findByReference(reference, false)
	if err != nil {
		return domain.User{}, err
	}

	return r.mapper.MapRepositoryToDomain(user), nil
}

func (r mongoUserRepository) findByReference(reference string, onlyActives bool) (MongoUser, error) {
	client := database.Mongo.Client
	collection := client.Database(r.config.Database).Collection(r.config.UsersCollection)

	user := MongoUser{}
	filter := bson.D{{Key: "reference", Value: reference}}
	if onlyActives {
		filter = append(filter, bson.E{Key: "is_active", Value: true})
	}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return MongoUser{}, nil
		}
		errMsg := "unexpected error when find user by its reference"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return MongoUser{}, errors.New(errMsg)
	}

	return user, nil
}

func (r mongoUserRepository) SearchActive(input domain.UserSearchInput) (domain.UserSearchOutput, error) {
	client := database.Mongo.Client
	collection := client.Database(r.config.Database).Collection(r.config.UsersCollection)

	filters := bson.D{{Key: "is_active", Value: true}}
	if len(input.FirstName) > 0 {
		filter := fmt.Sprintf("^%s", input.FirstName)
		filters = append(filters, bson.E{Key: "first_name", Value: primitive.Regex{Pattern: filter, Options: "i"}})
	}
	if len(input.LastName) > 0 {
		filter := fmt.Sprintf("^%s", input.LastName)
		filters = append(filters, bson.E{Key: "last_name", Value: primitive.Regex{Pattern: filter, Options: "i"}})
	}
	if len(input.Email) > 0 {
		filter := fmt.Sprintf("^%s", input.LastName)
		filters = append(filters, bson.E{Key: "email", Value: primitive.Regex{Pattern: filter, Options: "i"}})
	}
	limit := int64(input.PageSize)
	skip := int64((input.Page * input.PageSize) - input.PageSize)
	paging := options.FindOptions{Limit: &limit, Skip: &skip}

	users := []MongoUser{}
	cur, err := collection.Find(context.TODO(), filters, &paging)
	if err != nil {
		errMsg := "unexpected error when search active users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.UserSearchOutput{}, errors.New(errMsg)
	}

	err = cur.All(context.TODO(), &users)
	if err != nil {
		errMsg := "unexpected error when search active users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.UserSearchOutput{}, errors.New(errMsg)
	}

	total, err := collection.CountDocuments(context.TODO(), filters)
	if err != nil {
		errMsg := "unexpected error when search active users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.UserSearchOutput{}, errors.New(errMsg)
	}

	return r.mapper.MapRepositorySearchActiveToOutput(users, total, input.Page, input.PageSize), nil
}

func (r mongoUserRepository) Create(user domain.User) (domain.User, error) {
	client := database.Mongo.Client
	collection := client.Database(r.config.Database).Collection(r.config.UsersCollection)

	mongoUser := r.mapper.MapDomainToRepository(user)
	mongoUser.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), mongoUser)
	if err != nil {
		errMsg := "unexpected error when create the user"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.New(errMsg)
	}

	return user, nil
}

func (r mongoUserRepository) Update(user domain.User) (domain.User, error) {
	client := database.Mongo.Client
	collection := client.Database(r.config.Database).Collection(r.config.UsersCollection)

	// Find document to update
	currentUser, err := r.findByReference(user.Reference, false)
	if err != nil {
		return domain.User{}, err
	} else if len(currentUser.Reference) == 0 {
		return domain.User{}, errors.New("user to update was not found")
	}

	// Update document
	updatedUser := r.mapper.MapDomainToRepository(user)
	updatedUser.ID = currentUser.ID
	result, err := collection.ReplaceOne(context.TODO(), bson.D{{Key: "reference", Value: user.Reference}}, updatedUser)
	if err != nil {
		errMsg := "unexpected error when update the user"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.New(errMsg)
	}

	if result.MatchedCount != 1 || result.ModifiedCount != 1 {
		errMsg := "unexpected error when update the user. Matched update elements was not one"
		logger.AppLog.Error().Int64("matchedCount", result.MatchedCount).Int64("modifiedCount", result.ModifiedCount).Msg(errMsg)
		return domain.User{}, errors.New(errMsg)
	}

	return user, nil
}

func (r mongoUserRepository) Delete(reference string) (domain.User, error) {
	client := database.Mongo.Client
	collection := client.Database(r.config.Database).Collection(r.config.UsersCollection)

	// Find document to update
	currentUser, err := r.findByReference(reference, true)
	if err != nil {
		return domain.User{}, err
	} else if len(currentUser.Reference) == 0 {
		return domain.User{}, errors.New("user to delete was not found")
	}

	// Mark document as deleted
	currentUser.IsActive = false
	currentUser.UpdatedDate = time.Now().UTC()
	result, err := collection.ReplaceOne(context.TODO(), bson.D{{Key: "reference", Value: reference}}, currentUser)
	if err != nil {
		errMsg := "unexpected error when mark the user as deleted"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.New(errMsg)
	}

	if result.MatchedCount != 1 || result.ModifiedCount != 1 {
		errMsg := "unexpected error when mark the user as deleted. Matched update elements was not one"
		logger.AppLog.Error().Int64("matchedCount", result.MatchedCount).Int64("modifiedCount", result.ModifiedCount).Msg(errMsg)
		return domain.User{}, errors.New(errMsg)
	}

	return r.mapper.MapRepositoryToDomain(currentUser), nil
}
