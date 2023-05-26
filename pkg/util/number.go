package util

import "github.com/shopspring/decimal"

// 两个浮点数字相加
func AddFloat64(value1 float64, value2 float64) float64 {
	v1 := decimal.NewFromFloat(value1)
	v2 := decimal.NewFromFloat(value2)
	result, _ := v1.Add(v2).Float64()
	return result
}

// 两个浮点数字相减
func SubFloat64(value1 float64, value2 float64) float64 {
	v1 := decimal.NewFromFloat(value1)
	v2 := decimal.NewFromFloat(value2)
	result, _ := v1.Sub(v2).Float64()
	return result
}

// 两个浮点数字相乘
func MulFloat64(value1 float64, value2 float64) float64 {
	v1 := decimal.NewFromFloat(value1)
	v2 := decimal.NewFromFloat(value2)
	result, _ := v1.Mul(v2).Float64()
	return result
}

// 两个浮点数字相减
func DivFloat64(value1 float64, value2 float64) float64 {
	v1 := decimal.NewFromFloat(value1)
	v2 := decimal.NewFromFloat(value2)
	result, _ := v1.Div(v2).Float64()
	return result
}
