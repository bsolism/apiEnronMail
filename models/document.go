package models

import (
	"time"
)

type Document struct {
	took      int
	timed_out bool
	_shards   shard
	hits      hits
}

type shard struct {
	total      int
	successful int
	skipped    int
	failed     int
}
type source struct {
	Date    time.Time
	From    string
	Message string
	To      string
}

type hit struct {
	_index    string
	_type     string
	_id       string
	_score    int
	timestamp time.Time
	_source   source
}
type total struct {
	value int
}
type hits struct {
	total     total
	max_score int
	hits      []hit
}
