/*
Требуется вычислить период колебаний (t) математического маятника https://stepik.org/lesson/229321/step/7?unit=201907
*/
package main

import "math"

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}
