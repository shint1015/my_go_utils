package utils

import (
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type RoundOptions struct {
	RoundType string
	Precision int
}

// ToDecimal は、引数をDecimal型に変換
func ToDecimal(num interface{}, precision string) decimal.Decimal {
	var dec decimal.Decimal
	switch v := num.(type) {
	case string:
		dec, _ = decimal.NewFromString(v)
	case int:
		dec = decimal.NewFromInt(int64(v))
	case int64:
		dec = decimal.NewFromInt(v)
	case float32:
		dec = decimal.NewFromFloat32(v)
	case float64:
		dec = decimal.NewFromFloat(v)
	default:
		dec = decimal.NewFromInt(0)
	}
	if precision != "" {
		precNum, _ := strconv.ParseInt(precision, 10, 32)
		dec = dec.RoundFloor(int32(precNum))
	}
	return dec
}

// DecimalAdd は、引数の和を返す。(num1+num2)
func DecimalAdd(num1, num2 decimal.Decimal) decimal.Decimal {
	return num1.Add(num2)
}

// DecimalMul は、引数の積を返す。(num1*num2)
func DecimalMul(num1, num2 decimal.Decimal) decimal.Decimal {
	return num1.Mul(num2)
}

// DecimalQuo は、引数の商を返す。(num1/num2)
func DecimalQuo(num1, num2 decimal.Decimal, precision int) (decimal.Decimal, decimal.Decimal) {
	prec := int32(precision)
	return num1.QuoRem(num2, prec)
}

func DecimalDiv(num1, num2 decimal.Decimal) decimal.Decimal {
	return num1.Div(num2)
}

// DecimalSub は、引数の差を返す。(num1-num2)
func DecimalSub(num1, num2 decimal.Decimal) decimal.Decimal {
	return num1.Sub(num2)
}

/*
DecimalRound は、引数を指定の少数以下桁数で丸め込む。

roundType: "round", "ceil", "floor"
*/
func DecimalRound(num decimal.Decimal, roundType string, roundDigit int) decimal.Decimal {
	var calc decimal.Decimal
	if roundType == "round" {
		calc = num.Round(int32(roundDigit))
	} else if roundType == "ceil" {
		calc = num.RoundCeil(int32(roundDigit))
	} else if roundType == "floor" {
		calc = num.RoundFloor(int32(roundDigit))
	} else {
		calc = num
	}
	return calc
}

// IsPositive は、引数が正の数かどうかを判定する。
func IsPositive(num decimal.Decimal) bool {
	return num.Cmp(decimal.NewFromInt(int64(0))) > 0
}

// CalcChangeRate は、引数の変化率を計算する。
func CalcChangeRate(num1, num2 decimal.Decimal, precision int) decimal.Decimal {
	numDiff := DecimalSub(num1, num2)
	calc, _ := DecimalQuo(numDiff, num2, precision)
	quoNum := ToDecimal(100, strconv.Itoa(precision))
	calc2 := DecimalMul(calc, quoNum)
	return DecimalRound(
		calc2,
		"round",
		precision,
	)
}

// AdjustPrecision adjusts the precision of the decimal number.
// TargetDecimal: trueの場合、小数点以下の桁数を指定する。
func AdjustPrecision(num decimal.Decimal, digit int, targetDecimal bool) decimal.Decimal {
	precision := digit
	if targetDecimal == false {
		numCount := strings.Split(num.String(), ".")
		if len(numCount) == 1 {
			return num
		}
		intLen := len(numCount[0]) + 1
		precision = digit - intLen
	}
	return num.Truncate(int32(precision))
}

// ToString は、引数を文字列に変換する。
func ToString(num decimal.Decimal, digit int, targetDecimal bool) string {
	// TODO: 整数の頭には対応していない
	return AdjustPrecision(num, digit, targetDecimal).String()
}

func CountDecimalPrecision(num decimal.Decimal) (int, int) {
	numCount := strings.Split(num.String(), ".")
	if len(numCount) == 1 {
		return len(numCount[0]), 0
	}
	return len(numCount[0]), len(numCount[1])
}

func CmpTargetPrecision(num, num2 decimal.Decimal, targetPrecision int) int {
	numFormat := AdjustPrecision(num, targetPrecision, true)
	num2Format := AdjustPrecision(num2, targetPrecision, true)
	return numFormat.Cmp(num2Format)
}
