package main

import "time"

type commit struct {
	AuthorEmail    string
	AuthorName     string
	CommitterEmail string
	CommitterName  string
	ID             string
	Message        string
	Parent         string
	Timestamp      time.Time
	Tree           string
}

type DBTreeEntryType string

const (
	TREE     DBTreeEntryType = "tree"
	DATABASE                 = "db"
	LICENCE                  = "licence"
)

type dbTree struct {
	ID      string
	Entries []dbTreeEntry
}
type dbTreeEntry struct {
	AType         DBTreeEntryType
	Last_Modified time.Time
	Licence       string
	Sha256        string
	Size          int
	Name          string
}

const STORAGEDIR = "/Users/jc/tmp/dioapistorage"