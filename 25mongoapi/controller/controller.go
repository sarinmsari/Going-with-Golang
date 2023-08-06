package controller

import (
	"MongoAPI/model"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: change password to <password>
const connectionString = "mongodb+srv://doopstech:doops@cluster0.6vl9klf.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const collectionName = "Watchlist"

// MOST IMPORTANT: reference of mongodb collection
var collection *mongo.Collection

// Connect to MongoDB
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance created")
}

// MONGODG Helper functions

// insert one document into mongodb
func InsertOneMovie(movie model.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// update one document in mongodb
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Matched ", updateResult.MatchedCount, " documents and updated ", updateResult.ModifiedCount, " documents.")
}

// delete one document in mongodb
func DeleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{'_id':id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Movie deleted with delete count ",deleteResult.DeletedCount)
}

// delete all document in mongodb	
func deleteAllMovie() int64{
	deleteResult, err := collection.DeleteMany(context.Background(),bson.D{{}}, nil)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Movie deleted with delete count ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount;
}

// get all document in mongodb
func GetAllMovie() []primitive.M{
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []premitive.M

	cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies;
}



// Actual Controllers
func GetMyAllMovies(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies, err := GetAllMovie()
	if err != nil {
		log.Fatal(err);
	}
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE") 

	params:= mux.Vars(r)
	count:=DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(count)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE") 

	count:=deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}