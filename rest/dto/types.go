package dto

type Request struct {
	Password string `json:"password"`
	Rules    []Rule `json:"rules"`
}

type Rule struct {
	RuleName string `json:"rule"`
	Value    int    `json:"value"`
}

type Response struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
