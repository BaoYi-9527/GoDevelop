package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	// cook xiHongShi
	xiHongShi := &XiHongShi{}
	doCook(xiHongShi)

	fmt.Println("\n another cooking")

	// cook ChaoJiDan
	chaoJiDan := &ChaoJiDan{}
	doCook(chaoJiDan)
}
