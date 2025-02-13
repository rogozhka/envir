package main

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

    println(envPgHost, ":", pgHost)
    println(envPgPort, ":", pgPort)
    println(envHome, ":", home)
    // ...
}
