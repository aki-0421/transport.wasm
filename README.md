<!---
Thanks for reading 💖!

## Commit Message Naming Convention

- F for Feature (Additions, Fixes, Adjustments of functionalities, etc.)
- T for Testing (New tests / specs, Test refactoring, etc.)
- R for Refactor (Adjustments of code structure, naming, typing, comments, etc.)
- D for Documentation (Documentation, README, etc.)
- S for Style (Styling, Storybook, Theme, Visual Design Adjustments, etc.)
- V for Version (Versioning, Dependencies, Licensing, etc.)
- C for Configuration (Building, Linting, CLI Tooling, etc.)

-->

# transport.wasm

Go written Implement HTTP Transport. (Include useful hooks for Edge computing)

# simple example

```go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aki-0421/transport.wasm"
)

func main() {
	t, err := transport.New()
	if err != nil {
		panic(err)
	}

	cli := &http.Client{
		Transport: t,
	}
	resp, err := cli.Get("https://api.github.com/repos/aki-0421/transport.wasm/stargazers")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

```

# environment

- TinyGo v0.28.0-dev-0e94553b 