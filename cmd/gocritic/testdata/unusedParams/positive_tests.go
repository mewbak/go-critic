package checker_test

type foo struct{}

/// parameter `x` isn't used
func f1(x int, y float64) {
	_ = y
}

/// parameter `x` isn't used
/// parameter `y` isn't used
func (f *foo) f2(x int, y int) {
}

/// parameter `y` isn't used
func (f *foo) f3(_, y int, z float64) {
	_ = z
}

/// consider to name parameters as `_`
func (f *foo) f4(int, float64) {
}

/// parameter `x` isn't used
func (f *foo) f5(x int, _ float64) {
}
