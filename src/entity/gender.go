package entity

type Gender string

var genderValues = []byte{
	'F', // Female
	'M', // Male
	'O', // Other
}

func NewGender(gender string) (Gender, error) {
	if len(gender) != 1 {
		return "", NewErrInvalidGender(gender)
	}

	genderByte := gender[0]
	if genderByte >= 'Z' {
		genderByte -= 32 // convert to uppercase
	}

	for _, b := range genderValues {
		if b == genderByte {
			return Gender(genderByte), nil
		}
	}

	return "", NewErrInvalidGender(gender)
}
