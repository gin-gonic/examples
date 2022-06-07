package main

import "time"

// User map, included user1 and user2
// In a production environment, it is recommended to use a database instead
var users = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
}

// Session map stores the users sessions
// In a production environment, it is recommended to use a database or cache instead
var sessions = map[string]session{}

// Each session contains the username and expire time
type session struct {
	username   string
	expireTime time.Time
}

// Determine if the session has expired
func (s session) Expired() bool {
	return s.expireTime.Before(time.Now())
}
