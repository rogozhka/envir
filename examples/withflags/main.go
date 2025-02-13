package main

import (
    "os"

    "github.com/jessevdk/go-flags"
    "github.com/rogozhka/envir"
    "github.com/rogozhka/envir/decoders"
)

const envPrefix = "APPNAME"

func main() {
    var cliArgs struct {
        // ConfigFile is optional yaml configuration file to lookup values.
        ConfigFile string `long:"config" description:"Configuration yaml path" default:"config.yaml"`
    }
    parser := flags.NewParser(&cliArgs, flags.Default)
    if _, err := parser.Parse(); err != nil {
        code := 1
        if fe, ok := err.(*flags.Error); ok {
            if fe.Type == flags.ErrHelp {
                code = 0
            }
        }
        // logging is not ready here 
        println(err.Error())
        os.Exit(code)
    }

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
                        cliArgs.ConfigFile,
                        envir.WithDecoder(decoders.NewYaml()),
                    ),
                ),
                envir.WithOptional( // third priority, shell env w/ prefix
                    envir.NewLookupOs(),
                ),
            ),
        ),
    )

    // TODO: setup logging here

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
