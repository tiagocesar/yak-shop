package models

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
