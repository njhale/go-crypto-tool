package commands

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Hash(data, algo string) (string, error) {
	if data == "" {
		return "", fmt.Errorf("A non-empty data argument must be provided")
	}

	if algo == "" {
		algo = "sha256"
	}

	sum, ok := hashFunctions[algo]
	if !ok {
		return "", fmt.Errorf("Unsupported hash algorithm: %s not in [%s]", algo, hashFunctions)
	}

	hash, err := json.Marshal(hashResult{
		Algo: algo,
		Hash: hex.EncodeToString(sum([]byte(data))),
	})
	if err != nil {
		return "", fmt.Errorf("Failed to marshal hash result: %w", err)
	}

	return string(hash), nil
}

type hashResult struct {
	Algo string `json:"algo"`
	Hash string `json:"hash"`
}

var hashFunctions = hashFuncSet{
	"sha256": func(d []byte) []byte { h := sha256.Sum256(d); return h[:] },
	"md5":    func(d []byte) []byte { h := md5.Sum(d); return h[:] },
}

type hashFuncSet map[string]func([]byte) []byte

func (s hashFuncSet) String() string {
	return strings.Join(keys(s), ", ")
}

func keys[V any](m map[string]V) []string {
	set := make([]string, 0, len(m))
	for k := range m {
		set = append(set, k)
	}

	sort.Strings(set)
	return set
}
