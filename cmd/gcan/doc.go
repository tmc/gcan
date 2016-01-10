// NAME
//        gcan − gcan cli
//
// SYNOPSIS
//        gcan [OPTIONS]
//
// DESCRIPTION
// OPTIONS
//        −v, −−verbose
//               Be verbose.
//
//        −s, −−bootstrap‐server <default: ":4226">
//               Server addr to bootstrap with.
//
//        −h, −−help
//               Show this help message
//
// COMMANDS
//    cat
//        cat consumes messages from stdin
//
//        Usage: gcan [OPTIONS] cat [cat‐OPTIONS]
//
//
//        −t, −−topic (required)
//               Topic to send to.
//
//        −k, −−keySeparator <default: "\t">
//               Value to separate key from values in lines
//
//        −h, −−help
//               Show this help message
//
//
//
package main
