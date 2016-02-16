package core

const (
	GPIO0 = iota
	GPIO1
	GPIO2
	GPIO3
	GPIO4
	GPIO5
	GPIO6
	GPIO7
	GPIO8
	GPIO9
	GPIO10
	GPIO11
	GPIO12
	GPIO13
	GPIO14
	GPIO15
	GPIO16
	GPIO17
	GPIO18
	GPIO19
	GPIO20
	GPIO21
	GPIO22
	GPIO23

	A0 = iota
	A1
	A2
	A3
	A4
	A5
	A6
	A7
	A8
	A9
	A10
	A11
)

const (
	LOW = iota
	HIGH

	INPUT = iota
	OUTPUT
	INPUT_PULLUP

	LSBFIRST = iota
	MSBFIRST

	CHANGE = iota + 1
	FALLING
	RISING
)

const (
	DEV_GPIO_MEM = "/dev/gpiomem"
)
