package transport

type Option func(*Transport)

func WithFetchMode(m FetchMode) Option {
	return func(t *Transport) {
		t.Mode = m
	}
}

func WithFetchRedirect(r FetchRedirect) Option {
	return func(t *Transport) {
		t.Redirect = r
	}
}

func WithBeforeRequestFunc(f BeforeRequestFunc) Option {
	return func(t *Transport) {
		t.beforeRequest = f
	}
}

func WithModifyRequestFunc(f ModifyRequestFunc) Option {
	return func(t *Transport) {
		t.modifyRequest = f
	}
}

func WithModifyResponseFunc(f ModifyResponseFunc) Option {
	return func(t *Transport) {
		t.modifyResponse = f
	}
}
