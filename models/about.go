package models

type About struct {
	Id    uint   `son:"id"`
	Title string `json:title`

	Desc1 string `json:desc1`
	Desc2 string `json:desc2`
	Desc3 string `json:desc3`
}
