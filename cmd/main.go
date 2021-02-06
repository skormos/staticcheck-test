package main

import (
	"fmt"

	utils2 "github.com/skormos/staticcheck-test/internal/utils"
	"github.com/skormos/staticcheck-test-dependency/pkg/utils"
)

func main() {
	fmt.Println(utils.SoSad("missing chocolate"))
	fmt.Println(utils2.AlsoSoSad("missing more chocolate"))
}
