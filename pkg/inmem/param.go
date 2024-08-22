package inmem

type CacheInitParam struct {
	Size int
}

type Option func(param CacheInitParam) CacheInitParam

func WithSize(size int) Option {
	return func(param CacheInitParam) CacheInitParam {
		param.Size = size
		return param
	}
}
