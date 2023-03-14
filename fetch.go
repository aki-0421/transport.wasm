package transport

// FetchMode present an option to the Fetch API mode setting.
// Valid values are: "cors", "no-cors", "same-origin", "navigate". The default is "same-origin".
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters
type FetchMode string

const (
	FetchModeCORS       FetchMode = "cors"
	FetchModeNoCORS     FetchMode = "no-cors"
	FetchModeSameOrigin FetchMode = "same-origin"
	FetchModeNavigate   FetchMode = "navigate"
)

func (f FetchMode) String() string {
	return string(f)
}

func (f FetchMode) IsValid() bool {
	switch f {
	case FetchModeCORS:
		return true
	case FetchModeNoCORS:
		return true
	case FetchModeSameOrigin:
		return true
	case FetchModeNavigate:
		return true
	}
	return false
}

// FetchRedirect present an option to the Fetch API redirect setting.
// Valid values are: "follow", "error", "manual". The default is "follow".
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters
type FetchRedirect string

const (
	FetchRedirectFollow FetchRedirect = "follow"
	FetchRedirectError  FetchRedirect = "error"
	FetchRedirectManual FetchRedirect = "manual"
)

func (f FetchRedirect) String() string {
	return string(f)
}

func (f FetchRedirect) IsValid() bool {
	switch f {
	case FetchRedirectFollow:
		return true
	case FetchRedirectError:
		return true
	case FetchRedirectManual:
		return true
	}
	return false
}
