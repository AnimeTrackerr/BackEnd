package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Title struct {
	English string `bson:"english"` 
	Romaji string `bson:"romaji"` 
}

type Tag struct {
	Name string `bson:"name"`
}

type Date struct {	
	Year int `bson:"year"`
	Month int `bson:"month"`
	Day int `bson:"day"`
}

type CoverImage struct {	
	ExtraLarge string `bson:"extraLarge"`
	Large string `bson:"large"`
	Medium string `bson:"medium"`
	Color string `bson:"color"`
}

type VoiceActor struct {
	 Name string `bson:"name"`
    Image string `bson:"image"`	
}

type Character struct {	
	Role string `bson:"role"`
	VoiceActors []VoiceActor `bson:"voiceActors"`
	Name string `bson:"name"`
	Image string `bson:"image"`
}

type Studio struct {	
	Name string `bson:"name"`
}

type Media struct {
	ID				primitive.ObjectID	`bson:"_id"`
	AniID			int					`bson:"id"`
	Title 			Title 				`bson:"title"`
	Synomyms		[]string 			`bson:"synonyms"`
	Type 			string 				`bson:"type"`
	Episodes 		int 				`bson:"episodes"`
	MeanScoreAni	int 				`bson:"meanScoreAni"`
	Genres			[]string 			`bson:"genres"`
	Tags 			[]Tag 				`bson:"tags"`
	Status 			string 				`bson:"status"`
    IsAdult 		bool 				`bson:"isAdult"`
    Description 	string 				`bson:"description"`
    Season 			string 				`bson:"season"`
    StartDate 		Date 				`bson:"startDate"`
    EndDate 		Date 				`bson:"endDate"`
    BannerImage 	string 				`bson:"bannerImage"`
	CoverImage		CoverImage 			`bson:"coverImage"`
	Characters		[]Character 		`bson:"characters"`
	Studios			[]Studio 			`bson:"studios"`
}