package models

import (
	"fmt"
	"strconv"
)

type USD int64

func ToUSD(f float64) USD {
	return USD((f * 100) + 0.5)
}

func (u *USD) Float64() float64 {
	x := float64(*u)
	x = x / 100
	return x
}

func (u *USD) Multiply(f float64) USD {
	x := (float64(*u) * f) + 0.5
	return USD(x)
}

func (u *USD) String() string {
	x := float64(*u)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

func (u *USD) MarshalCSV() (string, error) {
	if u == nil {
		return "", nil
	}
	return strconv.FormatInt(int64(*u), 10), nil
}

func (u *USD) UnmarshalCSV(csv string) error {
	if csv == "" {
		*u = 0
		return nil
	}

	i, err := strconv.ParseInt(csv, 10, 64)
	if err != nil {
		return err
	}

	*u = USD(i)
	return nil
}
