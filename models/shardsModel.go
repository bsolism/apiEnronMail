package models

type Shard struct {
	total      int
	successful int
	skipped    int
	failed     int
}
