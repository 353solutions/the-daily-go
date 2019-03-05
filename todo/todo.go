package todo

import (
	"time"

	"github.com/att/todo/db"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	yaml "gopkg.in/yaml.v2"
)

// Todo is a todo item
type Todo struct {
	Key     string
	Content string
	Due     time.Time
	Done    bool
}

// NewTodo returns a new Todo item
func NewTodo(content string, due time.Time) *Todo {
	return &Todo{
		Key:     xid.New().String(),
		Content: content,
		Due:     due,
		Done:    false,
	}
}

// Marshal will marshal t to a []byte
func (t *Todo) Marshal() ([]byte, error) {
	return yaml.Marshal(t)
}

// DB is a Todo database
type DB struct {
	db *db.DB
}

// NewDB returns new database
func NewDB(dsn string) (*DB, error) {
	db, err := db.New(dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "can't create db from %s", dsn)
	}

	return &DB{db}, nil
}

// Keys returns all the keys in the databases
func (db *DB) Keys() ([]string, error) {
	return db.db.Keys()
}

// Save saves data to key
func (db *DB) Save(t *Todo) error {
	data, err := t.Marshal()
	if err != nil {
		return errors.Wrapf(err, "can't marshal %v", t)
	}

	return db.db.Save(t.Key, data)
}

// Load loads data for key
func (db *DB) Load(key string) (*Todo, error) {
	data, err := db.db.Load(key)
	if err != nil {
		return nil, errors.Wrapf(err, "%q: can't load", key)
	}

	t := &Todo{}
	if err := yaml.Unmarshal(data, t); err != nil {
		return nil, errors.Wrapf(err, "%q: can't unmarshal", key)
	}

	return t, nil
}

// Delete deletes key
func (db *DB) Delete(key string) error {
	return db.db.Delete(key)
}
