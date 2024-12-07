package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

const version = "0.1"

type command struct {
    flagSet *flag.FlagSet
    run func(args []string) error
}

// Main examines the args and delegates to the specified subcommand.
//
// If no subcommand was specified, we default to the "table" subcommand.
func main() {
    subcommands := map[string]command {
        "table": tableCmd(),
        "tree": treeCmd(),
    }

    mainFlagSet := flag.NewFlagSet("gits", flag.ExitOnError)
    versionFlag := mainFlagSet.Bool("version", false, "Print version and exit")

    mainFlagSet.Parse(os.Args[1:])

    if *versionFlag {
        fmt.Printf("%s\n", version)
        return
    }

    args := mainFlagSet.Args()

    var cmd command
    if len(args) == 0 {
        cmd = subcommands["table"]
    } else {
        var ok bool
        cmd, ok = subcommands[args[0]]
        if !ok {
            cmd = subcommands["table"]
        }
    }

    if err := cmd.run(args[len(args):]); err != nil {
        log.Fatal(err)
    }
}

// -------------------- Subcommand Definitions  --------------------------------

func tableCmd() command {
    flagSet := flag.NewFlagSet("gits table", flag.ExitOnError)

    return command{
        flagSet: flagSet,
        run: func(args []string) error {
            fmt.Println("Run tableCmd()")
            return nil
        },
    }
}

func treeCmd() command {
    flagSet := flag.NewFlagSet("gits tree", flag.ExitOnError)

    return command{
        flagSet: flagSet,
        run: func(args []string) error {
            fmt.Println("Run treeCmd()")
            return nil
        },
    }
}
