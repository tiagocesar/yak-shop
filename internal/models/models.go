package models

const (
	MinShavingAge = 101
	MaxYakAge     = 1000
)

type WrapperImport struct {
	Yaks []YakImport `xml:"labyak"`
}

type YakImport struct {
	Name string  `xml:"name,attr"`
	Age  float32 `xml:"age,attr"`
	Sex  string  `xml:"sex,attr"`
}

type Yak struct {
	Name      string
	AgeInDays int
	Sex       string
	Dead      bool
	NextShave int
}

func (yak *Yak) Milk() float32 {
	producedMilk := 50 - (float32(yak.AgeInDays) * 0.03)

	return producedMilk
}

func (yak *Yak) Shave(day int) int {
	if yak.AgeInDays >= MinShavingAge && yak.NextShave < day {
		yak.NextShave = 8 + int(float32(yak.AgeInDays)*0.01)

		return 1
	}

	return 0
}

func (yak *Yak) Age() {
	yak.AgeInDays += 1
	if yak.AgeInDays >= MaxYakAge {
		yak.Dead = true
	}
}
