package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/pawanmkr/mockint/graph/model"
	"github.com/pawanmkr/mockint/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client     *mongo.Client
	db         string
	collection string
}

var database string
var collection string

func Connect() *DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	var uri = os.Getenv("MONGO_URI")
	database = os.Getenv("DB")
	collection = os.Getenv("COLLECTION")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return &DB{
		client:     client,
		db:         database,
		collection: collection,
	}
}

func (db *DB) ScheduleInterview(input model.InterviewInput) *model.Interview {
	collection := db.client.Database(db.db).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserted, err := collection.InsertOne(ctx, bson.M{
		"duration":    input.Duration,
		"name":        input.Name,
		"time":        input.Time,
		"skills":      input.Skills,
		"guestType":   input.GuestType,
		"note":        input.Note,
		"difficulty":  input.Difficulty,
		"guest":       input.Guest,
		"booked":      input.Booked,
		"joinUrl":     input.JoinURL,
		"meetingCode": input.MeetingCode,
	})
	logError(err)

	return &model.Interview{
		ID:          inserted.InsertedID.(primitive.ObjectID).Hex(),
		Duration:    input.Duration,
		Time:        input.Time,
		Name:        input.Name,
		Skills:      input.Skills,
		Difficulty:  input.Difficulty,
		GuestType:   input.GuestType,
		Note:        input.Note,
		Booked:      input.Booked,
		JoinURL:     input.JoinURL,
		MeetingCode: input.MeetingCode,
	}
}

func (db *DB) GetMeetingById(id *string) *model.Interview {
	collection := db.client.Database(db.db).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var res model.Interview
	_id, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(ctx, filter).Decode(&res)
	logError(err)

	return &res
}

func (db *DB) GetAllMeetings() []*model.Interview {
	collection := db.client.Database(db.db).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var res []*model.Interview

	cursor, err := collection.Find(ctx, bson.M{})
	logError(err)

	if err := cursor.All(context.TODO(), &res); err != nil {
		panic(err)
	}

	return res
}

func (db *DB) UpdateMeeting(id string, input model.InterviewInput) *model.Interview {
	collection := db.client.Database(db.db).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var res model.Interview
	meet := bson.M{}

	if input.Difficulty != "" {
		meet["difficulty"] = input.Difficulty
	}

	if input.Duration != 0 {
		meet["duration"] = input.Duration
	}

	if input.Time != "" {
		meet["time"] = input.Time
	}

	if input.Name != "" {
		meet["name"] = input.Name
	}

	if input.Skills != "" {
		meet["skills"] = input.Skills
	}

	if input.GuestType != "" {
		meet["guestType"] = input.GuestType
	}

	if input.Note != "" {
		meet["note"] = input.Note
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": meet}

	err := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&res)
	logError(err)

	return &res
}

func (db *DB) BookMeeting(input model.BookInterview) *model.Interview {
	collection := db.client.Database(db.db).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var doc *model.Interview

	// chekcing if the interview is scheduled first or not
	_id, _ := primitive.ObjectIDFromHex(input.InterviewID)
	filter := bson.M{"_id": _id}

	err := collection.FindOne(ctx, filter, options.FindOne()).Decode(&doc)
	logError(err)

	// if the input is already available then return
	for _, guest := range doc.Guest {
		if guest.Email == input.Email {
			return nil
		}
	}

	sDnt := doc.Time
	layout := "2006-01-02T15:04:05.0000000-07:00"

	t, err := time.Parse(layout, sDnt)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	newTime := t.Add(time.Duration(doc.Duration) * time.Minute)
	eDnt := newTime.Format(layout)

	// creating a meeting on teams first, if not created then return
	body, _ := json.Marshal(map[string]string{
		"startDateTime": sDnt,
		"endDateTime":   eDnt,
		"subject":       "trying to generate a meet link using api",
	})

	// fmt.Println(sDnt)
	// fmt.Println(eDnt)

	var meeting *services.Meeting
	meeting, err = services.CreateMeeting(body)
	if err != nil {
		return nil
	}

	// update fields
	update := bson.M{
		"$push": bson.M{
			"guest": bson.M{
				"name":     input.Name,
				"whatsapp": input.Email,
			},
		},
		"$set": bson.M{
			"booked":      true,
			"joinUrl":     meeting.JoinUrl,
			"meetingCode": meeting.MeetingCode,
		},
	}

	services.SendMail(input.Email, meeting.JoinUrl)

	err = collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&doc)
	logError(err)

	return doc
}

func (db *DB) CancelMeeting(id string) *model.DeleteResponse {
	collection := db.client.Database(db.db).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	_, err := collection.DeleteOne(ctx, filter)
	logError(err)

	return &model.DeleteResponse{
		DeleteInterviewID: id,
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
