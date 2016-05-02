package util

import (
	"github.com/momazia/GoProject/datastore"
	"github.com/momazia/GoProject/log"
	"github.com/momazia/GoProject/memcache"
	"github.com/momazia/GoProject/session"
	"net/http"
)

// Stores the user in memcache and data store
func SaveUser(req *http.Request, u datastore.User) {
	// Storing into memcache
	err := memcache.Store(u.Email, u, req)
	log.LogErrorWithMsg("Cannot store the user into memcache:", err)
	// Storing into datastore
	err = datastore.Store(req, datastore.KIND_USER, u)
	log.LogErrorWithMsg("Cannot store the user into datastore:", err)
}

// Get's user's information from memcache, if it does not exists, it will look into datastore.
func GetUser(req *http.Request) datastore.User {
	// Getting user's email from session
	email := session.GetUser(req).Email

	// Getting the data from memcache
	var u datastore.User
	err := memcache.Retrieve(email, req, &u)
	if err != nil {
		log.Println("Cannot get the user from memcache", err)
		// Getting the user from datastore
		u, err = datastore.Retrieve(req, datastore.KIND_USER, email)
		log.LogError(err)
	}
	return u
}
