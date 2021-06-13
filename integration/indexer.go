package integration

import "github.com/mrobinsn/go-newznab/newznab"

type Indexer interface {
	Search(string) ([]newznab.NZB, error)
	TestConnection() error
}
