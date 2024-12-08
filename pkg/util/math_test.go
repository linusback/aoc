package util

import "testing"

func Test_Combinate(t *testing.T) {
	t.Log("running test")
	for c := range Combinate[byte](13, '0', '1') {
		t.Logf("%c\n", c)
	}
}

func Test_Combinate2(t *testing.T) {
	t.Log("running test")
	for _ = range Combinate[byte](13, '0', '1', '2') {
		//t.Logf("%c\n", c)
	}
}
