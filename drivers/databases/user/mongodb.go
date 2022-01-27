package user

import (
	"context"
	"disspace/business/user"
	"disspace/helpers/encryption"
	"disspace/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBUserRepository struct {
	Conn *mongo.Database
}

func NewMongoDBUserRepository(conn *mongo.Database) user.Repository {
	return &MongoDBUserRepository{
		Conn: conn,
	}
}

func (repository *MongoDBUserRepository) Register(ctx context.Context, data *user.UserDomain) (user.UserDomain, error) {
	newUser := UserFromDomain(*data)
	checkVar := []User{}
	filter := bson.D{{Key: "username", Value: newUser.Username}}
	check, errCheck := repository.Conn.Collection("users").Find(ctx, filter)
	if errCheck != nil {
		panic(errCheck)
	}
	if errCheck = check.All(ctx, &checkVar); errCheck != nil {
		panic(errCheck)
	}
	if len(checkVar) != 0 {
		return user.UserDomain{}, messages.ErrUsernameAlreadyExist
	}

	encryptedPass, _ := encryption.HashPassword(newUser.Password)
	newUser.Password = encryptedPass
	_, err := repository.Conn.Collection("users").InsertOne(ctx, newUser)

	newProfileUser := UserProfile{
		Username:    newUser.Username,
		ProfilePict: "https://firebasestorage.googleapis.com/v0/b/disspace-76973.appspot.com/o/user_profile_img%2Fprofile_default.jpg?alt=media&token=226ac15c-ebb6-4635-a708-2f923fd96808",
		Bio:         " ",
		Following:   []string{"0"},
		Followers:   []string{"0"},
		Threads:     []string{"0"},
		Reputation:  1,
	}
	_, errProfile := repository.Conn.Collection("user_profile").InsertOne(ctx, newProfileUser)
	if err != nil {
		return user.UserDomain{}, err
	}
	if errProfile != nil {
		return user.UserDomain{}, err
	}
	return user.UserDomain{}, nil
}

func (repository *MongoDBUserRepository) UserProfileGetByUsername(ctx context.Context, username string) (user.UserProfileDomain, error) {
	result := UserProfile{}
	filter := bson.D{{Key: "username", Value: username}}
	err := repository.Conn.Collection("user_profile").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return user.UserProfileDomain{}, err
	}
	// fmt.Println(username)

	return result.UserProfileToDomain(), nil
}

func (repository *MongoDBUserRepository) GetUserByID(ctx context.Context, id string) (user.UserDomain, error) {
	result := User{}
	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return user.UserDomain{}, errorConvert
	}
	filter := bson.D{{Key: "_id", Value: convert}}
	err := repository.Conn.Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return user.UserDomain{}, messages.ErrDataNotFound
	}
	// fmt.Println(result)
	return result.UserToDomain(), nil
}

func (repository *MongoDBUserRepository) Login(ctx context.Context, username string, password string) (user.UserDomain, error) {
	result := User{}
	filter := bson.D{{Key: "username", Value: username}}
	err := repository.Conn.Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return user.UserDomain{}, err
	}
	// decodedPass := encryption.Decode(result.Password)
	isMatch := encryption.CheckPasswordHash(password, result.Password)
	// fmt.Println(result.Password)
	if isMatch != true {
		return user.UserDomain{}, messages.ErrInvalidCredentials
	}
	return result.UserToDomain(), nil
}

func (repository *MongoDBUserRepository) InsertSession(ctx context.Context, dataSession user.UserSessionDomain) error {
	newSession := SessionFromDomain(dataSession)
	_, err := repository.Conn.Collection("session").InsertOne(ctx, newSession)
	if err != nil {
		return err
	}

	return nil
}

func (repository *MongoDBUserRepository) ConfirmAuthorization(ctx context.Context, session user.UserSessionDomain) (user.UserSessionDomain, error) {
	checkSession := SessionFromDomain(session)
	result := UserSession{}
	filter := bson.D{{Key: "token", Value: checkSession.Token}}
	err := repository.Conn.Collection("session").FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return user.UserSessionDomain{}, messages.ErrSessionNotFound
	}

	return result.SessionToDomain(), nil
}

func (repository *MongoDBUserRepository) CheckingSession(ctx context.Context, username string) error {
	result := UserSession{}
	filter := bson.D{{Key: "username", Value: username}}
	repository.Conn.Collection("session").FindOne(ctx, filter).Decode(&result)

	if result.Username == username {
		return messages.ErrAlreadyLoggedIn
	}

	return nil
}

func (repository *MongoDBUserRepository) UpdateUserProfile(ctx context.Context, username string, data user.UserProfileDomain) error {
	userProfile := UserProfileFromDomain(data)

	update := bson.D{{Key: "$set", Value: userProfile}}
	filter := bson.D{{Key: "username", Value: username}}
	err := repository.Conn.Collection("user_profile").FindOneAndUpdate(ctx, filter, update)
	if err.Err() != nil {
		return messages.ErrUpdateFailed
	}
	return nil
}

func (repository *MongoDBUserRepository) GetUserByUsername(ctx context.Context, username string) (user.UserDomain, error) {
	result := User{}

	filter := bson.D{{Key: "username", Value: username}}
	err := repository.Conn.Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return user.UserDomain{}, err
	}
	// fmt.Println(result)
	return result.UserToDomain(), nil
}

func (repository *MongoDBUserRepository) UpdateUserInfo(ctx context.Context, username string, data user.UserDomain) error {
	userProfile := UserFromDomain(data)

	update := bson.D{{Key: "$set", Value: userProfile}}
	filter := bson.D{{Key: "username", Value: username}}
	err := repository.Conn.Collection("users").FindOneAndUpdate(ctx, filter, update)
	if err.Err() != nil {
		return messages.ErrUpdateFailed
	}
	return nil
}

func (repository *MongoDBUserRepository) DeleteSession(ctx context.Context, dataSession user.UserSessionDomain) error {

	_, err := repository.Conn.Collection("session").DeleteOne(ctx, bson.D{{Key: "token", Value: dataSession.Token}})

	if err != nil {
		return err
	}
	return nil
}

func (repository *MongoDBUserRepository) GetModerators(ctx context.Context, idCategory string) ([]user.UserProfileDomain, error) {
	result := []user.UserProfileDomain{}
	// filter := bson.D{{Key: "kategori_id", Value: idCategory}}
	// cursor, err := repository.Conn.Collection("moderators").Find(ctx, filter)

	// if err = cursor.All(ctx, &result); err != nil {
	// 	return []user.UserProfileDomain{}, err
	// }
	// fmt.Println(result)
	return result, nil

}

func (repository *MongoDBUserRepository) GetAllUserProfile(ctx context.Context) ([]user.UserProfileDomain, error) {
	result := []user.UserProfileDomain{}
	cursor, err := repository.Conn.Collection("user_profile").Find(ctx, bson.M{})
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return []user.UserProfileDomain{}, err
	}
	return result, nil
}

func (repository *MongoDBUserRepository) Search(ctx context.Context, q string, sort string) ([]user.UserProfileDomain, error) {
	var result []user.UserProfileDomain
	if sort == "" {
		sort = "_id"
	}
	sorting := bson.M{sort: -1}
	opts := options.Find().SetSort(sorting)

	// Index for user_profile collection
	model := mongo.IndexModel{Keys: bson.D{{Key: "username", Value: "text"}}}
	_, errIndex := repository.Conn.Collection("user_profile").Indexes().CreateOne(ctx, model)
	if errIndex != nil {
		return []user.UserProfileDomain{}, errIndex
	}

	// Search Users
	filter := bson.D{}
	if q != "" {
		filter = bson.D{{Key: "username", Value: primitive.Regex{Pattern: q, Options: "i"}}}
	}

	cursor, err := repository.Conn.Collection("user_profile").Find(ctx, filter, opts)
	if err != nil {
		return []user.UserProfileDomain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []user.UserProfileDomain{}, err
	}
	return result, nil
}
