package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Farewell(name string) string {
	return fmt.Sprintf("Goodbye %s, It's been fun!", name)
}

func (a *App) GetCat(tag string) string {
	client := http.DefaultClient
	uri := "https://cataas.com/cat"
	if tag != "" {
		uri += "/" + tag
	}
	resp, err := client.Get(uri + "?json=true")
	if err != nil {
		panic(err)
	}
	url := struct {
		URL string
	}{}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &url)
	return "https://cataas.com" + url.URL
}

func (a *App) GetTags() []string {
	client := http.DefaultClient
	resp, err := client.Get("https://cataas.com/api/tags")
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	b = b[1 : len(b)-1] // remove '[' and ']'
	fields := bytes.Split(b, []byte{','})
	ans := []string{}
	for _, field := range fields {
		ans = append(ans, string(field[1:len(field)-1])) // remove `"` and `"`
	}

	return ans
}
