**envir** is a tiny library inspired by Viper. Focused on declarative OOP composition. You can setup desired priority and extend processing logic by adding your own variable expander or even different sources.

[![Go Report Card](https://goreportcard.com/badge/github.com/rogozhka/envir)](https://goreportcard.com/report/github.com/rogozhka/envir)
[![Releases](https://img.shields.io/github/release/rogozhka/envir/all.svg?style=flat-square)](https://github.com/rogozhka/envir/releases)
[![LICENSE](https://img.shields.io/github/license/rogozhka/envir.svg?style=flat-square)](https://github.com/rogozhka/envir/LICENSE)

Example: 

```go
import (
    "github.com/rogozhka/envir"
    "github.com/rogozhka/envir/decoders"
)

const envPrefix = "APPNAME"

func main() {
	env := envir.New(
		envir.WithPrefix(envPrefix),
		envir.WithEnvSubst(), // use "github.com/a8m/envsubst" as a custom vars processing 
		envir.WithLookup(
			envir.NewLookupComposition(
				envir.NewLookupFile(
					"app.yaml", // first priority, unprefixed vars inside env: section
					envir.WithCutPrefix(envPrefix),
					envir.WithDecoder(
						decoders.NewYaml(
							decoders.WithEntriesPath("env"),
						),
					),
				),
				envir.WithOptional( // second priority, prefixed vars in the document root
					envir.NewLookupFile(
						"dev.yaml",
						envir.WithDecoder(decoders.NewYaml()),
					),
				),
				envir.WithOptional( // third priority, shell env w/ prefix
					envir.NewLookupOs(),
				),
			),
		),
	)
	const (
		envPgHost = "DB_SERVER"
		envPgPort = "DB_PORT"
		envHome   = "HOME"
	)

	pgHost := env.String(envPgHost, "localhost")
	pgPort := env.Int(envPgPort, 5432)
	// ...

	home := env.MustString(envHome)
 
  // ...
}
```
