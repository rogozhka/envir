package main

import (
    "github.com/rogozhka/envir"
    "github.com/rogozhka/envir/decoders"
)

const envPrefix = "APPNAME"

func main() {
    env := envir.New(
        envir.WithPrefix(envPrefix),
        envir.WithLookup(
            envir.NewLookupComposition(
                envir.NewLookupFile(
                    ".env", // first priority
                    envir.WithDecoder(decoders.NewEnv()),
                ),
                envir.WithOptional( // second priority
                    envir.NewLookupFile(
                        "config.yaml",
                        envir.WithDecoder(decoders.NewYaml()),
                        envir.WithCutPrefix(envPrefix), // yaml is unprefixed
                    ),
                ),
                envir.WithOptional( // third priority
                    envir.NewLookupOs(),
                ),
            ),
        ),
    )
    const (
        envPgHost      = "PG_HOST"
        envKeystoneURL = "AUTH_KEYSTONE_URL"
    )

    pgHost := env.String(envPgHost, "localhost")
    keystoneURL := env.String(envKeystoneURL, "localhost:35000")
    // ...

    println(envPgHost, ":", pgHost)
    println(envKeystoneURL, ":", keystoneURL)
    // ...
}
