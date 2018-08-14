package api

import "encoding/json"

// Email payload's pattern to mount the Json of each mail
type Email struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Img                  string `json:"img"`
	DateDelivery         string `json:"date_delivery"`
	DateDeliveryFormated string `json:"date_delivery_formated"`
}

// GetEmailList gets list of emails
func GetEmailList() ([]byte, error) {
	data := []Email{
		{1, "Cael", "Descrição do cael", "/assets/images/hyoga.gif", "2018-06-25", "25/06/2018"},
		{2, "ueda", "Descrição do Ueda", "/assets/images/seiya.jpg", "2018-06-20", "20/06/2018"},
		{3, "celestino", "Descrição do Celestino", "/assets/images/shun.png", "2018-05-24", "24/05/2018"},
		{4, "vasco", "Descrição do Vasco", "/assets/images/vegeta.jpg", "2018-03-15", "15/03/2018"},
	}

	return json.Marshal(data)
}

//func que pega o objeto, iterando para filtrar o buscado
