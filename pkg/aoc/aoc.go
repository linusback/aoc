package aoc

import "os/exec"

type Part string
type SimpleError string

func (se SimpleError) Error() string {
	return string(se)
}

const (
	aocCli      = "aoc"
	Part1  Part = "1"
	Part2  Part = "2"
)

const (
	ErrCmdIsNil SimpleError = "cmd was nil for some reason"
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
