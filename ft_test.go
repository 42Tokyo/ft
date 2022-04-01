package ft_test

import "github.com/42tokyo/ft"

func ExamplePrintRune() {
	ft.PrintRune('4')
	ft.PrintRune('2')
	ft.PrintRune('\n')
	ft.PrintRune('♥')
	invalidRune := rune(-1)
	err := ft.PrintRune(invalidRune)
	if err == nil {
		panic("ft.PrintRune should fail with an invalid rune")
	}
	// Output:
	// 42
	// ♥
}