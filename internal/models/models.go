package models

type Herd struct {
	Yaks []Yak `xml:"labyak"`
}

type Yak struct {
	Name string  `xml:"name,attr"`
	Age  float32 `xml:"age,attr"`
	Sex  string  `xml:"sex,attr"`
}
