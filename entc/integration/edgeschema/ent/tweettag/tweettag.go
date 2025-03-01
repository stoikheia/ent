// Code generated by ent, DO NOT EDIT.

package tweettag

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tweettag type in the database.
	Label = "tweet_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddedAt holds the string denoting the added_at field in the database.
	FieldAddedAt = "added_at"
	// FieldTagID holds the string denoting the tag_id field in the database.
	FieldTagID = "tag_id"
	// FieldTweetID holds the string denoting the tweet_id field in the database.
	FieldTweetID = "tweet_id"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeTweet holds the string denoting the tweet edge name in mutations.
	EdgeTweet = "tweet"
	// Table holds the table name of the tweettag in the database.
	Table = "tweet_tags"
	// TagTable is the table that holds the tag relation/edge.
	TagTable = "tweet_tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "tag_id"
	// TweetTable is the table that holds the tweet relation/edge.
	TweetTable = "tweet_tags"
	// TweetInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetInverseTable = "tweets"
	// TweetColumn is the table column denoting the tweet relation/edge.
	TweetColumn = "tweet_id"
)

// Columns holds all SQL columns for tweettag fields.
var Columns = []string{
	FieldID,
	FieldAddedAt,
	FieldTagID,
	FieldTweetID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAddedAt holds the default value on creation for the "added_at" field.
	DefaultAddedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
