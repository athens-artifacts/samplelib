package samplelib

import (
	"fmt"

	"github.com/athens-artifacts/samplelib/types"
)

func Hello() string {
	return fmt.Sprintf("%s", types.Lion{})
}
