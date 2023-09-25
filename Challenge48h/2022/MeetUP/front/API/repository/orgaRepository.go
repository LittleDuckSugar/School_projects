package repository

import "meetupAPI/models"

func GetOrgaById(id string) models.Orga {
	return models.Orga{Id: id, Username: "to1", Tel: "00000000", Note: 4, Password: "1234", Email: "exemple@exemple.com"}
}

func GetAllOrga() []models.Orga {
	return []models.Orga{{Id: "1", Username: "to1", Tel: "00000000", Note: 4, Password: "1234", Email: "exemple@exemple.com"}, {Id: "2", Username: "to2", Tel: "00003000", Note: 5, Password: "12345678", Email: "exemple@e3xemple.com"}}
}
