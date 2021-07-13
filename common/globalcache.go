package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	GlobalCache *cache.Cache
)

func InitGlobalCache() {
	GlobalCache = cache.New(5*time.Minute, 10*time.Minute)
}
