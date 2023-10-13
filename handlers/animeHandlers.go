package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AnimeTrackerr/v2/backend/DB"
	"github.com/AnimeTrackerr/v2/backend/models"

	"github.com/AnimeTrackerr/v2/backend/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ColCount int = 0

func SetCollection(URI string) {
	DB_conn := DB.ConnectDB(URI)
	collection = DB.GetCollection(DB_conn.Client ,"animeDB", "animeCollection")
	count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	ColCount = int(count)
}

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To AnimeTracker v2.0 :)")
}

func GetAnime(w http.ResponseWriter, r *http.Request) {
	animeID, err := strconv.Atoi(mux.Vars(r)["id"])	

	// handle conversion errors
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error : Bad Request in offset and limit query params\n")
		return
	}

	// fetch anime
	var result models.Media
	err = collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: animeID}}).Decode(&result)

	// check if ID exists
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error : No document was found with the id %d\n", animeID)
		return
	}
	if err != nil {
		panic(err)
	}

	// return json encoded response
	json.NewEncoder(w).Encode(result)
}	

func GetAllAnime(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var (
		pageLimit int = 25
		offSet int = 0
		err1 error
		err2 error
	)

	// input validation - (spell check, types handling and invalid params)
	if !q.Has("offset") && !q.Has("limit") && len(q) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error 400: Bad Request\n")
		return
	}

	// handle string conversion errors
	if q.Has("offset") {offSet, err1 = strconv.Atoi(q.Get("offset"))}
	if q.Has("limit") {pageLimit, err2 = strconv.Atoi(q.Get("limit"))}
	
	// handle conversion errors
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error : Bad Request in offset and limit query params\n")
		return
	}

	// restrict query params to a specific range
	offSet = utils.RestrictNumber(offSet, 0, ColCount)
	pageLimit = utils.RestrictNumber(pageLimit, 0, 50)

	// fetch list of anime
	var result models.Page

	cursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSkip(int64(offSet)).SetLimit(int64(pageLimit)))
	if err != nil {
		panic(err)
	}

	err = cursor.All(context.TODO(), &result.AnimeList)
	if err != nil {
		panic(err)
	}

	result.PageInfo = models.PageInfo{
		HasNext : true,
		DocsCount : len(result.AnimeList),
	}

	// return json encoded response
	json.NewEncoder(w).Encode(result)
}

func GetRandomAnime(w http.ResponseWriter, r *http.Request) {
	// aggregate to get random sample
	aggregate := []bson.M{{"$sample": bson.M{"size": 1}}}
	var animes []models.Media

	cursor, err := collection.Aggregate(context.TODO(), aggregate)
    if err!=nil {
        panic(err)
    }

	// decode list of docs into var animes
	cursor.All(context.TODO(), &animes)

	// return json encoded response
	json.NewEncoder(w).Encode(animes)
}

func SearchAnime(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("term")

	if len(query) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request Error : Empty Query\n")
		return
	}
	
	var animes []models.Media
	pipeline := bson.A{
		bson.D{
			{Key: "$search",
				Value: bson.D{
					{Key: "index", Value: "animeSearchIndex"},
					{Key: "autocomplete",
						Value: bson.D{
							{Key: "query", Value: query},
							{Key: "path", Value: "synonyms"},
						},
					},
				},
			},
		},
		bson.D{{Key: "$limit", Value: 7}},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "title", Value: 1},
					{Key: "meanScoreAni", Value: 1},
					{Key: "startDate", Value: 1},
					{Key: "endDate", Value: 1},
					{Key: "status", Value: 1},
					{Key: "type", Value: 1},
					{Key: "episodes", Value: 1},
				},
			},
		},
	}
	
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}

	// decode list of docs into var animes
	cursor.All(context.TODO(), &animes)

	// return json encoded response
	json.NewEncoder(w).Encode(animes)
}