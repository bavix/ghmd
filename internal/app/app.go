package app

import "regexp"

type ReplacerFunc func([]byte) []byte
type Option func(replacer) replacer

type ReplacerInterface interface {
	Replace(input []byte) []byte
}

type replacer struct {
	replacers []ReplacerFunc
}

func New(ops ...Option) ReplacerInterface {
	rpl := replacer{}

	for _, o := range ops {
		rpl = o(rpl)
	}

	return &rpl
}

func (r *replacer) Replace(input []byte) []byte {
	output := append(make([]byte, 0, len(input)), input...)

	for _, rpl := range r.replacers {
		output = rpl(output)
	}

	return output
}

func WithUserReplacer() Option {
	return func(r replacer) replacer {
		r.replacers = append(r.replacers, userReplacer)

		return r
	}
}

var userRegexp = regexp.MustCompile("([^[])@([A-Za-z0-9_-]+)([^]])")

func userReplacer(input []byte) []byte {
	return userRegexp.ReplaceAll(input, []byte("$1[@$2](https://github.com/$2)$3"))
}
