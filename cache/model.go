package cache

// Character model for storing and accessing information from the cache
type Character struct {
	Uid     string   `json:"uid"`
	Species []string `json:"species"`
}

// Interface that all library functions have to impliment
// in order to talk to cache
type libInterface = func(name string) (string, []string, error)
