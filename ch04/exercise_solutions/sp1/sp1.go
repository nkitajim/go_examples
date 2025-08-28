package main

import (
	"fmt"

	// "math"
	"math/rand"
	"time"
)

func switch_return(ix int) int {
	switch {
	case 0 == ix%6:
		return 6
	case 0 == ix%2:
		return 2
	case 0 == ix%3:
		return 3
	default:
		return 1
	}
}

func switch_return_10(ix int) int {
	switch ix {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return 4
	case 5:
		return 5
	case 6:
		return 6
	case 7:
		return 7
	case 8:
		return 8
	case 9:
		return 9
	default:
		return 10
	}
}

func if_return(ix int) int {
	if 0 == ix%6 {
		return 6
	} else if 0 == ix%2 {
		return 2
	} else if 0 == ix%3 {
		return 3
	} else {
		return 1
	}
}

func if_return_10(ix int) int {
	if 0 == ix {
		return 0
	} else if 1 == ix {
		return 1
	} else if 2 == ix {
		return 2
	} else if 3 == ix {
		return 3
	} else if 4 == ix {
		return 4
	} else if 5 == ix {
		return 5
	} else if 6 == ix {
		return 6
	} else if 7 == ix {
		return 7
	} else if 8 == ix {
		return 8
	} else if 9 == ix {
		return 9
	} else {
		return 10
	}
}

func main() {
	var start, end time.Time

	is := make([]int, 0, 1000000)
	rand.Seed(time.Now().UnixNano())
	for ix := 0; ix < 1000000; ix++ {
		// is = append(is, rand.Intn(math.MaxInt32))
		is = append(is, rand.Intn(10))
	}

	sum := 0
	start = time.Now()
	for _, ix := range is {
		switch {
		case 0 == ix%6:
			sum += 6
		case 0 == ix%2:
			sum += 2
		case 0 == ix%3:
			sum += 3
		default:
			sum += 1
		}
	}
	end = time.Now()
	fmt.Println("switch:", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		switch ix {
		case 6:
			sum += 6
		case 0, 2, 4, 8:
			sum += 2
		case 3, 9:
			sum += 3
		default:
			sum += 1
		}
	}
	end = time.Now()
	fmt.Println("switch multi key:", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		if 0 == ix%6 {
			sum += 6
		} else if 0 == ix%2 {
			sum += 2
		} else if 0 == ix%3 {
			sum += 3
		} else {
			sum += 1
		}
	}
	end = time.Now()
	fmt.Println("if", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		switch ix {
		case 6:
			sum += 6
		case 2:
			sum += 2
		case 3:
			sum += 3
		default:
			sum += 1
		}
	}
	end = time.Now()
	fmt.Println("simple switch:", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		if 6 == ix {
			sum += 6
		} else if 2 == ix {
			sum += 2
		} else if 3 == ix {
			sum += 3
		} else {
			sum += 1
		}
	}
	end = time.Now()
	fmt.Println("simple if", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		switch ix {
		case 0:
			sum += 0
		case 1:
			sum += 1
		case 2:
			sum += 2
		case 3:
			sum += 3
		case 4:
			sum += 4
		case 5:
			sum += 5
		case 6:
			sum += 6
		case 7:
			sum += 7
		case 8:
			sum += 8
		case 9:
			sum += 9
		default:
			sum += 10
		}
	}
	end = time.Now()
	fmt.Println("simple 10 switch:", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		if 0 == ix {
			sum += 0
		} else if 1 == ix {
			sum += 1
		} else if 2 == ix {
			sum += 2
		} else if 3 == ix {
			sum += 3
		} else if 4 == ix {
			sum += 4
		} else if 5 == ix {
			sum += 5
		} else if 6 == ix {
			sum += 6
		} else if 7 == ix {
			sum += 7
		} else if 8 == ix {
			sum += 8
		} else if 9 == ix {
			sum += 9
		} else {
			sum += 10
		}
	}
	end = time.Now()
	fmt.Println("simple 10 if", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		sum += switch_return(ix)
	}
	end = time.Now()
	fmt.Println("func switch_return", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		sum += if_return(ix)
	}
	end = time.Now()
	fmt.Println("func if_return", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		sum += switch_return_10(ix)
	}
	end = time.Now()
	fmt.Println("func switch_return_10", sum, end.Sub(start))

	sum = 0
	start = time.Now()
	for _, ix := range is {
		sum += if_return_10(ix)
	}
	end = time.Now()
	fmt.Println("func if_return_10", sum, end.Sub(start))
}
