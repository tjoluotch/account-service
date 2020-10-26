package config

import (
	"crypto/rand"
	"math/big"
)

func IdGenerator() (string, error) {
	generator := rand.Reader
	max := big.NewInt(1000000)
	min := big.NewInt(100000)
	diff := max.Sub(max, min)
	res, err := rand.Int(generator, diff)
	if err != nil {
		return "", err
	}
	res = res.Add(res, min)
	return res.String(), nil
}
