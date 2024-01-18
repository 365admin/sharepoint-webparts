package matrix

import (
	"fmt"
	"testing"
)

func TestAnalysePages(t *testing.T) {

	p, err := AnalysePages()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for k, v := range p {
		fmt.Println(k, len(v))

	}

}

func TestWriteMatrix(t *testing.T) {
	WriteMatrix()
}
