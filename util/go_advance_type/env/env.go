package env

import (
	"os"
	"regexp"
	"strings"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

func Set(key string, value interface{}) error {
	return os.Setenv(key, convert.MustStr(value))
}
func Get(key string, defaultValue interface{}) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	} else {
		return convert.MustStr(defaultValue)
	}
}
func GetInt(key string, defaultValue int) int {
	return convert.SafeInt(Get(key, ""), defaultValue)
}
func GetFloat64(key string, defaultValue float64) float64 {
	return convert.SafeFloat64(Get(key, ""), defaultValue)
}

func GetEnvs(prefixes ...string) []string {
	var result []string
	for _, env := range os.Environ() {
		for _, prefix := range prefixes {
			if regexp.MustCompile(`^` + strings.ToLower(prefix)).MatchString(strings.ToLower(env)) {
				result = append(result, env)
				break
			}
		}
	}
	return result
}
func ToString(prefixes ...string) string {
	envs := GetEnvs(prefixes...)
	body := `
key = value
----------------------------
`
	for _, env := range envs {
		sp := strings.Split(env, "=")
		if len(sp) >= 2 {
			body += sp[0] + " = " + sp[1] + "\n"
		}
	}
	body += `
----------------------------` + "\n"
	return body
}
