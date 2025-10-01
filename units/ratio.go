package units

type Ratio float64

func (f Ratio) ToPercentage() float64 {
	return float64(100 - (100 / f))
}

type Percentage float64

func (f Percentage) ToRatio() float64 {
	return float64(100 / (100 - f))
}
