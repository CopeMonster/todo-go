package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-go/models"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"userID"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Done        bool               `bson:"done"`
}

func toModel(todo *models.Todo) *Todo {
	uid, _ := primitive.ObjectIDFromHex(todo.UserID)

	return &Todo{
		UserID:      uid,
		Title:       todo.Title,
		Description: todo.Description,
		Done:        todo.Done,
	}
}

func toTodo(todo *Todo) *models.Todo {
	return &models.Todo{
		ID:          todo.ID.Hex(),
		UserID:      todo.UserID.Hex(),
		Title:       todo.Title,
		Description: todo.Description,
		Done:        todo.Done,
	}
}

func toTodos(todos []*Todo) []*models.Todo {
	out := make([]*models.Todo, len(todos))

	for i, t := range todos {
		out[i] = toTodo(t)
	}

	return out
}

type TodoRepository struct {
	db *mongo.Collection
}

func NewTodoRepository(db *mongo.Database, collection string) *TodoRepository {
	return &TodoRepository{
		db: db.Collection(collection),
	}
}

func (r *TodoRepository) GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)
	var result *Todo

	err := r.db.FindOne(ctx, bson.M{
		"_id":    objID,
		"userID": uID,
	}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return toTodo(result), nil
}

func (r *TodoRepository) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	cur, err := r.db.Find(ctx, bson.M{
		"userID": uID,
	})

	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Todo, 0)

	for cur.Next(ctx) {
		todo := new(Todo)

		err := cur.Decode(todo)

		if err != nil {
			return nil, err
		}

		out = append(out, todo)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toTodos(out), nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, user *models.User, td *models.Todo) error {
	td.UserID = user.ID

	model := toModel(td)

	res, err := r.db.InsertOne(ctx, model)

	if err != nil {
		return err
	}

	td.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, user *models.User, id string, td *models.Todo) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	update := bson.M{
		"$set": bson.M{
			"title":       td.Title,
			"description": td.Description,
			"done":        td.Done,
		},
	}

	_, err := r.db.UpdateOne(ctx, bson.M{
		"_id":    objID,
		"userID": uID,
	}, update)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{
		"_id":    objID,
		"userID": uID,
	})

	return err
}
