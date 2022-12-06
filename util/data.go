package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func cfgDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(homeDir, ".config", "advent-of-code")
}

func cacheDir() string {
	return path.Join(cfgDir(), "cache")
}

func mustGetAOCSession() string {
	session, err := ioutil.ReadFile(path.Join(cfgDir(), "session"))
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(session))
}

func cacheFile(year, day int) string {
	return path.Join(cacheDir(), fmt.Sprintf("%d-%d.txt", year, day))
}

func mustCacheData(year, day int, data string) {
	if _, err := os.Stat(cacheDir()); os.IsNotExist(err) {
		if err := os.Mkdir(cacheDir(), os.ModePerm); err != nil {
			panic(err)
		}
	}
	if err := os.WriteFile(cacheFile(year, day), []byte(data), 0600); err != nil {
		panic(err)
	}
}

func cacheData(year, day int) (string, bool) {
	if _, err := os.Stat(cacheFile(year, day)); err == nil {
		data, err := os.ReadFile(cacheFile(year, day))
		if err != nil {
			panic(err)
		}
		return string(data), true
	}
	return "", false
}

func MustDailyInput(year, day int) string {
	if d, ok := cacheData(year, day); ok {
		return d
	}

	fmt.Printf("Getting input for %d/%d from AoC site... ", year, day)
	u := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: mustGetAOCSession(),
		// Value: "53616c7465645f5f39a5bf58fc16f0adc15fdf098c016aa3652afaaab310cc1e26f39b2cbc87f5dafa993f49d6cc60bc5e63b9097f1c218258a75afa5cfe276b",
	})
	req.Header.Add("Accept", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("unexpected response: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := string(body)

	mustCacheData(year, day, data)
	fmt.Printf("✅\n")

	return string(data)
}
