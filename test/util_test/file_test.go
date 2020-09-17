package uitl_test

import (
	myUtils "go-core/utils"
	"testing"
)

func TestFile(t *testing.T) {
	path := myUtils.GetPath()
	t.Log(path)
}
