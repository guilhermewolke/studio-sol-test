package internal

import (
	"log"
	"unicode"

	"github.com/guilhermewolke/studio-sol-test/graphql/graph/model"
)

type Verifier struct {
	Request  model.Request
	Response model.Response
}

var (
	MinSizeRule         = "minSize"
	MinUppercaseRule    = "minUppercase"
	MinLowercaseRule    = "minLowercase"
	MinSpecialCharsRule = "minSpecialChars"
	MinDigitRule        = "minDigit"
	NoRepeatedRule      = "noRepeted"
)

func NewVerifier(request model.Request) *Verifier {
	return &Verifier{Request: request}
}

func (v *Verifier) ReleaseTheKraken() {
	var response model.Response
	response.Verify = true
	response.NoMatch = make([]*string, 0, 6)

	for _, rule := range v.Request.Rules {
		switch rule.Rule {
		case MinSizeRule:
			v.MinSize(rule, &response)
		case MinUppercaseRule:
			v.MinUppercase(rule, &response)
		case MinLowercaseRule:
			v.MinLowercase(rule, &response)
		case MinSpecialCharsRule:
			v.MinSpecialChars(rule, &response)
		case MinDigitRule:
			v.MinDigit(rule, &response)
		case NoRepeatedRule:
			v.NoRepeated(rule, &response)
		}
	}

	if len(response.NoMatch) > 0 {
		response.Verify = false
	}

	v.Response = response
}

func (v *Verifier) MinSize(obj *model.Rule, response *model.Response) {
	log.Printf("MinSize - String: %s | size: %d | mínimo solicitado: %d", v.Request.Password, len(v.Request.Password), obj.Value)
	if len(v.Request.Password) < obj.Value {
		response.NoMatch = append(response.NoMatch, &MinSizeRule)
	}
}

func (v *Verifier) MinUppercase(obj *model.Rule, response *model.Response) {
	var count int = 0
	for _, l := range v.Request.Password {
		if unicode.IsUpper(l) && unicode.IsLetter(l) {
			count++
		}
	}
	log.Printf("MinUppercase - String: %s | Quantidade de caracteres em maiúsculo: %d | mínimo solicitado: %d", v.Request.Password, count, obj.Value)
	if count < obj.Value {
		response.NoMatch = append(response.NoMatch, &MinUppercaseRule)
	}
}

func (v *Verifier) MinLowercase(obj *model.Rule, response *model.Response) {
	var count int = 0
	for _, l := range v.Request.Password {
		if unicode.IsLower(l) && unicode.IsLetter(l) {
			count++
		}
	}

	log.Printf("MinLowercase - String: %s | Quantidade de caracteres em minúsculo: %d | mínimo solicitado: %d", v.Request.Password, count, obj.Value)

	if count < obj.Value {
		response.NoMatch = append(response.NoMatch, &MinLowercaseRule)
	}
}

func (v *Verifier) MinSpecialChars(obj *model.Rule, response *model.Response) {
	var count int = 0
	for _, l := range v.Request.Password {
		if !unicode.IsNumber(l) && !unicode.IsLetter(l) {
			count++
		}
	}

	log.Printf("MinSpecialChars - String: %s | Quantidade de caracteres especiais: %d | mínimo solicitado: %d", v.Request.Password, count, obj.Value)

	if count < obj.Value {
		response.NoMatch = append(response.NoMatch, &MinSpecialCharsRule)
	}
}

func (v *Verifier) MinDigit(obj *model.Rule, response *model.Response) {
	var count int = 0
	for _, l := range v.Request.Password {
		if unicode.IsNumber(l) {
			count++
		}
	}

	log.Printf("MinDigit - String: %s | Quantidade de números: %d | mínimo solicitado: %d", v.Request.Password, count, obj.Value)

	if count < obj.Value {
		response.NoMatch = append(response.NoMatch, &MinDigitRule)
	}
}

func (v *Verifier) NoRepeated(obj *model.Rule, response *model.Response) {
	var lastLetter string
	for i := 0; i < len(v.Request.Password); i++ {
		log.Printf("NoRepeated - String: %s | índice da letra: %d | letra do índice: %s | última letra: %s", v.Request.Password, i, string(v.Request.Password[i]), lastLetter)
		if i == 0 {
			lastLetter = string(v.Request.Password[i])
			continue
		}

		if lastLetter == string(v.Request.Password[i]) {
			response.NoMatch = append(response.NoMatch, &NoRepeatedRule)
			break
		}

		lastLetter = string(v.Request.Password[i])
	}
}
