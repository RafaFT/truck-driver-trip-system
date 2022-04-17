package entity

import (
	"regexp"
	"strconv"
)

var (
	cpfPattern = regexp.MustCompile(`^\d{11}$`)

	invalidCPFs = map[string]bool{
		"00000000000": true,
		"11111111111": true,
		"22222222222": true,
		"33333333333": true,
		"44444444444": true,
		"55555555555": true,
		"66666666666": true,
		"77777777777": true,
		"88888888888": true,
		"99999999999": true,
	}
)

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
	if _, isKnownInvalid := invalidCPFs[cpf]; isKnownInvalid || !cpfPattern.MatchString(cpf) {
		return false
	}

	sum := 0
	for i, strDigit := range cpf[:9] {
		digit, _ := strconv.Atoi(string(strDigit))
		sum += digit * (10 - i)
	}
	if firstDigitCheck := (sum * 10) % 11 % 10; strconv.Itoa(firstDigitCheck) != string(cpf[9]) {
		return false
	}

	sum = 0
	for i, strDigit := range cpf[:10] {
		digit, _ := strconv.Atoi(string(strDigit))
		sum += digit * (11 - i)
	}
	if secondDigitCheck := (sum * 10) % 11 % 10; strconv.Itoa(secondDigitCheck) != string(cpf[10]) {
		return false
	}

	return true
}
