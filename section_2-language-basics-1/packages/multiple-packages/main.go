package main

import (
	"fmt"

	"github.com/steevehook/udemy-go101/section_2-language-basics-1/packages/multiple-packages/pkg1"
	"github.com/steevehook/udemy-go101/section_2-language-basics-1/packages/multiple-packages/pkg2"
)

func main() {
	fmt.Println(pkg1.A, pkg2.A)
}
