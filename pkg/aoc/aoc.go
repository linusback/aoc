package aoc

import (
	"errors"
	"fmt"
	"github.com/linusback/aoc2024/pkg/errorsx"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Part string

const (
	aocCli          = "aoc"
	Part1      Part = "1"
	Part2      Part = "2"
	InputFile       = "input"
	PuzzleFile      = "puzzle.md"
)

const (
	ErrCmdIsNil                  errorsx.SimpleError = "cmd was nil for some reason"
	ErrCouldNotFindSessionCookie errorsx.SimpleError = "could not find session cookie"
)

var (
	sessionCookie    string
	sessionCookieErr error
)

func Send(part Part, day, answer string) error {
	cmd := exec.Command(aocCli, "submit", string(part), answer, "-d", day)
	if cmd == nil {
		return ErrCmdIsNil
	}
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Download(year string, day string) error {
	location := fmt.Sprintf("./internal/year%s/day%s/", year, day)
	inputFound, _, err := filesAlreadyExists(location)
	if err != nil {
		return err
	}

	client, err := getSessionClient()
	if err != nil {
		return err
	}

	errCh := make(chan error, 2)
	wg := new(sync.WaitGroup)
	downloadPuzzleAsync(wg, errCh, client, year, day, location)
	if !inputFound {
		downloadInputAsync(wg, errCh, client, year, day, location)
	}

	wg.Wait()
	close(errCh)
	for e := range errCh {
		err = errors.Join(err, e)
	}
	return err
}

func getSessionClient() (client *http.Client, err error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("https://adventofcode.com/")
	if err != nil {
		return nil, err
	}

	val, err := getSessionCookie()
	if err != nil {
		return nil, err
	}

	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "session",
			Value: val,
		},
	})

	client = &http.Client{
		Jar: jar,
	}
	return client, nil
}

func downloadInputAsync(wg *sync.WaitGroup, errCh chan error, client *http.Client, year, day, location string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := downloadInput(client, year, day, location)
		if err != nil {
			errCh <- err
			return
		}
	}()
}

func downloadInput(client *http.Client, year, day, location string) error {
	return nil
}

func downloadPuzzleAsync(wg *sync.WaitGroup, errCh chan<- error, client *http.Client, year, day, location string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := downloadPuzzle(client, year, day, location)
		if err != nil {
			errCh <- err
			return
		}
	}()
}

func downloadPuzzle(client *http.Client, year, day, location string) error {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day), nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed request status: %d", resp.StatusCode)
	}
	f, err := os.OpenFile(location+PuzzleFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func getSessionCookie() (string, error) {
	sync.OnceFunc(func() {
		var (
			ok   bool
			s    string
			err  error
			home string
		)
		s, ok = os.LookupEnv("ADVENT_OF_CODE_SESSION")
		if ok {
			sessionCookie = s
			return
		}
		home, err = os.UserHomeDir()
		if err != nil {
			sessionCookieErr = err
			return
		}

		s, err = getSessionFromFile(home + "/.config/adventofcode.session")
		if err == nil {
			sessionCookie = s
			return
		}
		sessionCookieErr = err
		s, err = getSessionFromFile(home + "/.adventofcode.session")
		if err == nil {
			sessionCookieErr = nil
			sessionCookie = s
			return
		}
		sessionCookieErr = errors.Join(ErrCouldNotFindSessionCookie, sessionCookieErr, err)
	})()
	return sessionCookie, sessionCookieErr
}

func getSessionFromFile(s string) (string, error) {
	f, err := os.Open(s)
	if err != nil {
		return "", err
	}
	sb := new(strings.Builder)
	_, err = io.Copy(sb, f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(sb.String()), nil
}

func filesAlreadyExists(dir string) (inputFound, puzzleFound bool, err error) {
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, e := range dirEntry {
		if e.Name() == InputFile && e.Type().IsRegular() {
			inputFound = true
		}
		if e.Name() == PuzzleFile && e.Type().IsRegular() {
			puzzleFound = true
		}
		if inputFound && puzzleFound {
			return
		}
	}

	return
}
