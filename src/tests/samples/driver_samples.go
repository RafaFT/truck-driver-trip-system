package samples

import (
	"testing"
	"time"

	"github.com/rafaft/truck-driver-trip-system/entity"
)

func GetDrivers(t *testing.T) []*entity.Driver {
	now := time.Now()

	driversInput := []struct {
		CPF        string
		BirthDate  time.Time
		CNH        string
		Gender     string
		HasVehicle bool
		Name       string
	}{
		{
			CPF:        "33510345398",
			BirthDate:  now.AddDate(-71, 0, 0),
			CNH:        "a",
			Gender:     "m",
			HasVehicle: true,
			Name:       "Kaique João Teixeira",
		},
		{
			CPF:        "52742089403",
			BirthDate:  now.AddDate(-65, 0, 0),
			CNH:        "b",
			Gender:     "f",
			HasVehicle: false,
			Name:       "Lorenzo Ian Carlos Eduardo Drumond",
		},
		{
			CPF:        "70286951150",
			BirthDate:  now.AddDate(-60, 0, 0),
			CNH:        "c",
			Gender:     "o",
			HasVehicle: true,
			Name:       "Gustavo Francisco Cardoso",
		},
		{
			CPF:        "94982599769",
			BirthDate:  now.AddDate(-57, 0, 0),
			CNH:        "d",
			Gender:     "m",
			HasVehicle: false,
			Name:       "Fátima Vanessa Monteiro",
		},
		{
			CPF:        "08931283849",
			BirthDate:  now.AddDate(-52, 0, 0),
			CNH:        "e",
			Gender:     "f",
			HasVehicle: true,
			Name:       "Ana Giovanna Porto",
		},
		{
			CPF:        "31803413603",
			BirthDate:  now.AddDate(-46, 0, 0),
			CNH:        "a",
			Gender:     "o",
			HasVehicle: false,
			Name:       "Rayssa Emanuelly Andrea Viana",
		},
		{
			CPF:        "72595478710",
			BirthDate:  now.AddDate(-41, 0, 0),
			CNH:        "b",
			Gender:     "m",
			HasVehicle: true,
			Name:       "Raimundo Erick Nicolas Souza",
		},
		{
			CPF:        "94665431728",
			BirthDate:  now.AddDate(-37, 0, 0),
			CNH:        "c",
			Gender:     "f",
			HasVehicle: false,
			Name:       "Bernardo Rafael Julio Figueiredo",
		},
		{
			CPF:        "27188079463",
			BirthDate:  now.AddDate(-33, 0, 0),
			CNH:        "d",
			Gender:     "o",
			HasVehicle: true,
			Name:       "Thales Marcos Fogaça",
		},
		{
			CPF:        "45464490388",
			BirthDate:  now.AddDate(-30, 0, 0),
			CNH:        "e",
			Gender:     "m",
			HasVehicle: false,
			Name:       "Bárbara Valentina Ana Barros",
		},
		{
			CPF:        "56820381930",
			BirthDate:  now.AddDate(-28, 0, 0),
			CNH:        "a",
			Gender:     "f",
			HasVehicle: true,
			Name:       "Vitória Adriana Freitas",
		},
		{
			CPF:        "69048144620",
			BirthDate:  now.AddDate(-27, 0, 0),
			CNH:        "b",
			Gender:     "o",
			HasVehicle: false,
			Name:       "Fernanda Beatriz Brenda Costa",
		},
		{
			CPF:        "97863413923",
			BirthDate:  now.AddDate(-26, 0, 0),
			CNH:        "c",
			Gender:     "m",
			HasVehicle: true,
			Name:       "Pietro Victor Sebastião Barros",
		},
		{
			CPF:        "57765277677",
			BirthDate:  now.AddDate(-25, 0, 0),
			CNH:        "d",
			Gender:     "f",
			HasVehicle: false,
			Name:       "Carolina Ester da Paz",
		},
		{
			CPF:        "77163670303",
			BirthDate:  now.AddDate(-23, 0, 0),
			CNH:        "e",
			Gender:     "o",
			HasVehicle: true,
			Name:       "Allana Louise Bianca Nogueira",
		},
		{
			CPF:        "17844926805",
			BirthDate:  now.AddDate(-20, 0, 0),
			CNH:        "a",
			Gender:     "m",
			HasVehicle: false,
			Name:       "Evelyn Catarina Nascimento",
		},
		{
			CPF:        "63503201238",
			BirthDate:  now.AddDate(-18, 0, 0),
			CNH:        "b",
			Gender:     "f",
			HasVehicle: true,
			Name:       "Otávio Benício Ricardo Ramos",
		},
		{
			CPF:        "01764805100",
			BirthDate:  now.AddDate(-67, 0, 0),
			CNH:        "c",
			Gender:     "o",
			HasVehicle: false,
			Name:       "Mariana Luciana Alice Silveira",
		},
		{
			CPF:        "31300454652",
			BirthDate:  now.AddDate(-53, 0, 0),
			CNH:        "d",
			Gender:     "m",
			HasVehicle: true,
			Name:       "Danilo Bryan Mateus Melo",
		},
		{
			CPF:        "64053595061",
			BirthDate:  now.AddDate(-58, 0, 0),
			CNH:        "e",
			Gender:     "f",
			HasVehicle: false,
			Name:       "Laís Sônia Pereira",
		},
	}

	drivers := make([]*entity.Driver, len(driversInput))
	for i, input := range driversInput {
		driver, err := entity.NewDriver(
			input.CPF,
			input.Name,
			input.Gender,
			input.CNH,
			input.BirthDate,
			input.HasVehicle,
		)

		if err != nil {
			t.Fatalf("%d: could not generate driver. [input: %v]", i, input)
		}

		drivers[i] = driver
	}

	return drivers
}
