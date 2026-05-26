package sprig

import (
	"github.com/shopspring/decimal"
)

// toFloat64 converts 64-bit floats
func toFloat64(v interface{}) float64 { _ = "STUB: not implemented"; return 0 }

func toInt(v interface{}) int { _ = "STUB: not implemented"; return 0 }

// toInt64 converts integer types to 64-bit integers
func toInt64(v interface{}) int64 { _ = "STUB: not implemented"; return 0 }

func max(a interface{}, i ...interface{}) int64 { _ = "STUB: not implemented"; return 0 }

func maxf(a interface{}, i ...interface{}) float64 { _ = "STUB: not implemented"; return 0 }

func min(a interface{}, i ...interface{}) int64 { _ = "STUB: not implemented"; return 0 }

func minf(a interface{}, i ...interface{}) float64 { _ = "STUB: not implemented"; return 0 }

func until(count int) []int { _ = "STUB: not implemented"; return nil }

func untilStep(start, stop, step int) []int { _ = "STUB: not implemented"; return nil }

func floor(a interface{}) float64 { _ = "STUB: not implemented"; return 0 }

func ceil(a interface{}) float64 { _ = "STUB: not implemented"; return 0 }

func round(a interface{}, p int, rOpt ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// converts unix octal to decimal
func toDecimal(v interface{}) int64 { _ = "STUB: not implemented"; return 0 }

func seq(params ...int) string { _ = "STUB: not implemented"; return "" }

func intArrayToString(slice []int, delimeter string) string { _ = "STUB: not implemented"; return "" }

// performs a float and subsequent decimal.Decimal conversion on inputs,
// and iterates through a and b executing the mathmetical operation f
func execDecimalOp(a interface{}, b []interface{}, f func(d1, d2 decimal.Decimal) decimal.Decimal) float64 {
	_ = "STUB: not implemented"
	return 0
}
