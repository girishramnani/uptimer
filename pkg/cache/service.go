package cache

import "sync"

type ServiceCache struct {
	mutex sync.RWMutex
	store map[string]int
}

func NewServiceCache() *ServiceCache {
	return &ServiceCache{
		mutex: sync.RWMutex{},
		store: make(map[string]int),
	}
}

func (sc *ServiceCache) ToReadableMap() map[string]string {
	rmap := make(map[string]string)
	sc.mutex.RLock()
	for k, v := range sc.store {
		rmap[k] = codeToString(v)
	}
	sc.mutex.RUnlock()
	return rmap
}

func (sc *ServiceCache) Set(url string, status_code int) {
	sc.mutex.Lock()
	sc.store[url] = status_code
	sc.mutex.Unlock()
}

func (sc *ServiceCache) Get(url string) int {
	sc.mutex.RLock()
	val := sc.store[url]
	sc.mutex.RUnlock()
	return val
}

func (sc *ServiceCache) GetString(url string) string {
	code := sc.Get(url)
	return codeToString(code)

}
func codeToString(code int) string {
	if code >= 200 && code < 500 {
		return "Available"
	} else if code >= 500 {
		return "Server Error"
	}
	// in case its 0
	return "Not reachable"
}
