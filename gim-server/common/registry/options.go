package registry

func WithAddr(addr string) OptionFunc {
	return func(options *Options) {
		options.Addr = addr
	}
}

func WithName(name string) OptionFunc {
	return func(options *Options) {
		options.Name = name
	}
}

func WithVersion(version string) OptionFunc {
	return func(options *Options) {
		options.Name = version
	}
}

func WithEndpoints(endpoints []string) OptionFunc {
	return func(options *Options) {
		options.Endpoints = endpoints
	}
}

func WithMode(mode string) OptionFunc {
	return func(options *Options) {
		options.Mode = mode
	}
}
