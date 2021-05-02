package mongo

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type SearchResult struct {
	bongo.DocumentBase
	FcmTokens []string
}

func SearchFcmByIds(ids []string) ([]string, error) {
	col := Conn.Collection("Rooms")
	rp := col.Find(nil)
	cursor := rp.Query.Select(bson.M{
		"_id": bson.M{
			"$elemMatch": bson.M{
				"$in": ids,
			},
		},
		"FcmTokens": 1,
	})
	// cursor = cursor.Sort()
	token := &SearchResult{}
	tokens := []string{}
	iter := cursor.Iter()
	for iter.Next(&token) {
		tokens = append(tokens, token.FcmTokens...)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return tokens, nil
}
