package entity

type CNH string

const (
	cnhA byte = 65 + iota // "A"
	cnhB                  // "B"
	cnhC                  // "C"
	cnhD                  // "D"
	cnhE                  // "E"
	cnhMax
)

func NewCNH(cnh string) (CNH, error) {
	if len(cnh) != 1 {
		return "", NewErrInvalidCNH(cnh)
	}

	cnhByte := cnh[0]
	if cnhByte >= 'Z' {
		cnhByte -= 32 // convert to uppercase
	}
	if cnhByte < cnhA || cnhByte >= cnhMax {
		return "", NewErrInvalidCNH(cnh)
	}

	return CNH(cnhByte), nil
}
