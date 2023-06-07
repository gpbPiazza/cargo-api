package repositories

type LocationRepository struct{}

func (lr LocationRepository) FindByPortCode(code string) {}
func (lr LocationRepository) FindByCityName(city string) {}
