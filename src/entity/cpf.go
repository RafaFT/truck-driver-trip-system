package entity

type CPF string

func NewCPF(cpf string) (CPF, error) {
	if !isCPFValid(cpf) {
		return "", NewErrInvalidCPF(cpf)
	}

	return CPF(cpf), nil
}

// https://dicasdeprogramacao.com.br/algoritmo-para-validar-cpf/
// https://pt.wikipedia.org/wiki/Cadastro_de_pessoas_f%C3%ADsicas
// http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/funcoes.js
func isCPFValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	// all CPF's with equal digits are invalid
	sameDigits := true
	for _, digit := range cpf {
		if digit != rune(cpf[0]) {
			sameDigits = false
			break
		}
	}
	if sameDigits {
		return false
	}

	var digits [11]int
	for i, digit := range cpf {
		// avoid heavy Atoi allocation by using code points
		if digit < 48 || digit > 57 {
			return false
		}
		digits[i] = int(digit) - 48
	}

	sum := 0
	for i, digit := range digits[:9] {
		sum += digit * (10 - i)
	}
	if firstDigitCheck := (sum * 10) % 11 % 10; firstDigitCheck != digits[9] {
		return false
	}

	sum = 0
	for i, digit := range digits[:10] {
		sum += digit * (11 - i)
	}
	if secondDigitCheck := (sum * 10) % 11 % 10; secondDigitCheck != digits[10] {
		return false
	}

	return true
}
