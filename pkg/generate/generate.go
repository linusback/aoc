package generate

import (
	"embed"
	"errors"
	"fmt"
	"github.com/linusback/aoc2024/pkg/errorsx"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"slices"
	"text/template"
)

const (
	ErrFailedToGetModuleName errorsx.SimpleError = "failed tog et name of go module"
)

//go:embed templates
var files embed.FS

type yearData struct {
	Imports []string
	Cases   []string
	Year    string
	ModName string
}

type dayData struct {
	Day  string
	Year string
}

func Generate(year string, days []string) error {
	if len(days) == 0 {
		return nil
	}
	moduleName, err := getModuleName()
	if err != nil {
		return err
	}

	err = generateSolver(moduleName, year)
	if err != nil {
		return err
	}

	err = generateYearSolve(moduleName, year, days)
	if err != nil {
		return err
	}
	for _, day := range days {
		err = generatDaySolve(moduleName, year, day)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func generateSolver(moduleName, year string) error {
	years := getYears(year)
	f, err := os.Create("./internal/solve.go")
	if err != nil {
		return err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			err = errors.Join(err, err2)
		}
	}()

	data := yearData{
		Imports: make([]string, 0, len(years)),
		Cases:   make([]string, 0, len(years)),
		Year:    year,
		ModName: moduleName,
	}
	intendent := ""
	for i, y := range years {
		data.Imports = append(data.Imports, fmt.Sprintf(`%s"%s/internal/year%s"`, intendent, moduleName, y))
		data.Cases = append(data.Cases, fmt.Sprintf(`%scase "%s":
		return year%s.Solve(day)`, intendent, y, y))
		if i == 0 {
			intendent = "\n\t"
		}
	}
	t, err := template.ParseFS(files, "templates/solve")
	if err != nil {
		return err
	}
	err = t.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func getYears(year string) []string {
	yearRegex, err := regexp.Compile(`\d{4}`)
	if err != nil {
		return nil
	}
	oldFile, err := os.ReadFile("./internal/solve.go")
	if err != nil {
		return nil
	}
	byteYears := yearRegex.FindAll(oldFile, -1)
	result := make([]string, 0, len(byteYears)+1)
	result = append(result, year)
	for _, yb := range byteYears {
		y := string(yb)
		if !slices.Contains(result, y) {
			result = append(result, y)
		}
	}
	slices.Sort(result)
	return result
}

func generateYearSolve(moduleName, year string, days []string) error {
	err := os.Mkdir(fmt.Sprintf("./internal/year%s", year), 0775)
	f, err := os.Create(fmt.Sprintf("./internal/year%s/solve.go", year))
	if err != nil {
		return err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			err = errors.Join(err, err2)
		}
	}()
	data := yearData{
		Imports: make([]string, 0, len(days)),
		Cases:   make([]string, 0, len(days)),
		Year:    year,
		ModName: moduleName,
	}
	intendent := ""
	for i, day := range days {
		data.Imports = append(data.Imports, fmt.Sprintf(`%s"%s/internal/year%s/day%s"`, intendent, moduleName, year, day))
		data.Cases = append(data.Cases, fmt.Sprintf(`%scase "%s":
		return day%s.Solve()`, intendent, day, day))
		if i == 0 {
			intendent = "\n\t"
		}
	}
	t, err := template.ParseFS(files, "templates/solve_year")
	if err != nil {
		return err
	}
	err = t.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func generatDaySolve(moduleName, year string, day string) error {
	data := dayData{
		Day:  day,
		Year: year,
	}

	t, err := template.ParseFS(files, "templates/solve_day")
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	return nil
}

func getModuleName() (string, error) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", ErrFailedToGetModuleName
	}

	// this might be hacky.
	if len(bi.Deps) == 0 {
		return "", ErrFailedToGetModuleName
	}
	return bi.Deps[0].Path, nil
}
