package option

import "time"

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

// NewConnect
// @Description: 创建一个连接
// @param addr
// @return *Connection
// @return error
func NewConnect(addr string) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   defaultCaching,
		timeout: defaultTimeout,
	}, nil
}

// NewConnectWithOptions
// @Description: 创建一个带选项的示例
// @param addr
// @param cache
// @param timeout
// @return *Connection
// @return error
func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   cache,
		timeout: timeout,
	}, nil
}

// 相对优雅实现

// ConnectionOptions
// @Description:
type ConnectionOptions struct {
	Caching bool
	Timeout time.Duration
}

func NewDefaultOptions() *ConnectionOptions {
	return &ConnectionOptions{
		Caching: defaultCaching,
		Timeout: defaultTimeout,
	}
}

func NewConnectPerf(addr string, opts *ConnectionOptions) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   opts.Caching,
		timeout: opts.Timeout,
	}, nil
}

// 选项模式实现

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func NewConnection(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}
