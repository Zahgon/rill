// Package mockapi provides a very basic mock API for examples and demos.
// It's intentionally kept public to enable running and experimenting with examples in the Go Playground.
// The implementation is naive and uses full scan for all operations.
package mockapi

import (
	"context"
	"sync"
	"time"
)

type User struct {
	ID         int
	Name       string
	Age        int
	Department string
	IsActive   bool
}

// don't use pointers here, to make sure that raw data is not accessible from outside
var departments []string
var users []User

var mu sync.RWMutex

func init() {
	const usersCount = 100

	var adjs = []string{"Big", "Small", "Fast", "Slow", "Smart", "Happy", "Sad", "Funny", "Serious", "Angry"}
	var nouns = []string{"Dog", "Cat", "Bird", "Fish", "Mouse", "Elephant", "Lion", "Tiger", "Bear", "Wolf"}

	mu.Lock()
	defer mu.Unlock()

	departments = []string{"HR", "IT", "Finance", "Marketing", "Sales", "Support", "Engineering", "Management"}

	// Generate users
	// Use deterministic values for all fields to make examples reproducible
	users = make([]User, 0, usersCount)

	for i := 1; i <= usersCount; i++ {
		user := User{
			ID:         i,
			Name:       adjs[hash(i, "name1")%len(adjs)] + " " + nouns[hash(i, "name2")%len(nouns)], // adj + noun
			Age:        hash(i, "age")%20 + 30,                                                      // 20-50
			Department: departments[hash(i, "dep")%len(departments)],                                // one of
			IsActive:   hash(i, "active")%100 < 60,                                                  // 60%
		}

		users = append(users, user)
	}
}

func GetDepartments() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetUser returns a user by ID.
func GetUser(ctx context.Context, id int) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUsers returns a list of users by IDs.
// If a user is not found, nil is returned in the corresponding position.
func GetUsers(ctx context.Context, ids []int) ([]*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type UserQuery struct {
	Department string
	Page       int
}

// ListUsers returns a paginated list of users optionally filtered by department.
func ListUsers(ctx context.Context, query *UserQuery) ([]*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveUser saves a user.
func SaveUser(ctx context.Context, user *User) error { _ = "STUB: not implemented"; return nil }

func getUserIndex(id int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func hash(input ...any) int { _ = "STUB: not implemented"; return 0 }

func randomSleep(ctx context.Context, max time.Duration) { _ = "STUB: not implemented"; return }
