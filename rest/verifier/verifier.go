package verifier

import (
	"log"
	"unicode"

	"github.com/guilhermewolke/studio-sol-test/rest/dto"
)

var (
	MinSizeRule         = "minSize"
	MinUppercaseRule    = "minUppercase"
	MinLowercaseRule    = "minLowercase"
	MinSpecialCharsRule = "minSpecialChars"
	MinDigitRule        = "minDigit"
	NoRepeatedRule      = "noRepeted"
)

func ReleaseTheKraken(request dto.Request) dto.Response {
	var response dto.Response
	response.Verify = true

	for _, v := range request.Rules {
		switch v.RuleName {
		case MinSizeRule:
			MinSize(request.Password, v, &response)
		case MinUppercaseRule:
			MinUppercase(request.Password, v, &response)
		case MinLowercaseRule:
			MinLowercase(request.Password, v, &response)
		case MinSpecialCharsRule:
			MinSpecialChars(request.Password, v, &response)
		case MinDigitRule:
			MinDigit(request.Password, v, &response)
		case NoRepeatedRule:
			NoRepeated(request.Password, v, &response)
		}

	}

	return response
}

func MinSize(password string, obj dto.Rule, response *dto.Response) {
	log.Printf("MinSize - String: %s | size: %d | mínimo solicitado: %d", password, len(password), obj.Value)
	if len(password) < obj.Value {
		response.Verify = false
		response.NoMatch = append(response.NoMatch, MinSizeRule)
	}
}

func MinUppercase(password string, obj dto.Rule, response *dto.Response) {
	var count int = 0
	for _, l := range password {
		if unicode.IsUpper(l) && unicode.IsLetter(l) {
			count++
		}
	}
	log.Printf("MinUppercase - String: %s | Quantidade de caracteres em maiúsculo: %d | mínimo solicitado: %d", password, count, obj.Value)
	if count < obj.Value {
		response.Verify = false
		response.NoMatch = append(response.NoMatch, MinUppercaseRule)
	}
}

func MinLowercase(password string, obj dto.Rule, response *dto.Response) {
	var count int = 0
	for _, l := range password {
		if unicode.IsLower(l) && unicode.IsLetter(l) {
			count++
		}
	}

	log.Printf("MinLowercase - String: %s | Quantidade de caracteres em minúsculo: %d | mínimo solicitado: %d", password, count, obj.Value)

	if count < obj.Value {
		response.Verify = false
		response.NoMatch = append(response.NoMatch, MinLowercaseRule)
	}
}

func MinSpecialChars(password string, obj dto.Rule, response *dto.Response) {
	var count int = 0
	for _, l := range password {
		if !unicode.IsNumber(l) && !unicode.IsLetter(l) {
			count++
		}
	}

	log.Printf("MinSpecialChars - String: %s | Quantidade de caracteres especiais: %d | mínimo solicitado: %d", password, count, obj.Value)

	if count < obj.Value {
		response.Verify = false
		response.NoMatch = append(response.NoMatch, MinSpecialCharsRule)
	}
}

func MinDigit(password string, obj dto.Rule, response *dto.Response) {
	var count int = 0
	for _, l := range password {
		if unicode.IsNumber(l) {
			count++
		}
	}

	log.Printf("MinDigit - String: %s | Quantidade de números: %d | mínimo solicitado: %d", password, count, obj.Value)

	if count < obj.Value {
		response.Verify = false
		response.NoMatch = append(response.NoMatch, MinDigitRule)
	}
}

func NoRepeated(password string, obj dto.Rule, response *dto.Response) {
	var lastLetter string
	for i := 0; i < len(password); i++ {
		log.Printf("NoRepeated - String: %s | índice da letra: %d | letra do índice: %s | última letra: %s", password, i, string(password[i]), lastLetter)
		if i == 0 {
			lastLetter = string(password[i])
			continue
		}

		if lastLetter == string(password[i]) {
			response.Verify = false
			response.NoMatch = append(response.NoMatch, NoRepeatedRule)
			break
		}

		lastLetter = string(password[i])
	}
}
