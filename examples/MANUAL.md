hello-world(1) -- Example Service for GoLift Application Builder
===

SYNOPSIS
---
`hello-world -c /etc/hello-world/helloworld.conf`

This daemon prints hello world.

OPTIONS
---
`hello-world [-c <config-file>] [-h] [-v]`

    -c, --config <config-file>
        Provide a configuration file (instead of the default).

    -v, --version
        Display version and exit.

    -h, --help
        Display usage and exit.

CONFIGURATION
---

`Config File Parameters`

    hellos       default: 2
        How many hellos to print.

    worlds       default: 1
        How many worlds to print.


AUTHOR
---
*   David Newhall II (still going) 2019

LOCATION
---
*   Application Builder: [https://github.com/golift/application-builder](https://github.com/golift/application-builder)
