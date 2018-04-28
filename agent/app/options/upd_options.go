package options

type Option func(c *Options)

type Options struct {
	Host           string
	Port           int
	ReceiveBuffLen int
}

func Host(h string) Option {
	return func(c *Options) {
		c.Host = h
	}
}

func Port(p int) Option {
	return func(c *Options) {
		c.Port = p
	}
}

func ReceiveBuffLen(l int) Option {
	return func(c *Options) {
		c.ReceiveBuffLen = l
	}
}

//申请结构体值
func ApplyOptions(options ...Option) Options {
	opts := Options{}
	for _, option := range options {
		option(&opts)
	}
	return opts
}
