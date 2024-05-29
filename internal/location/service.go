package location

type LocationService struct {
	client *LocationClient
}

func NewLocationService() *LocationService {
	return &LocationService{
		client: NewLocationClient(),
	}
}

func (ls *LocationService) GetLocation(cep string) (string, error) {
	viaCEP, err := ls.client.GetLocation(cep)
	if err != nil {
		return "", err
	}
	return viaCEP.Localidade, nil
}
