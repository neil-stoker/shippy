package main

import (
	"gopkg.in/mgo.v2"
)

// // CreateClient -
// func CreateClient(uri string) (*mongo.Client, error) {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
// }

// CreateSession -
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
