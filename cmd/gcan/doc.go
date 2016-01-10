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
//        valid compression codecs: NONE, SNAPPY, ZLIB
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
//        −c, −−compression <default: "NONE">
//               Compression codec to use.
//
//        −h, −−help
//               Show this help message
//
//    tail
//        tail consumes messages from gcan and produces them on stdout
//
//        Usage: gcan [OPTIONS] tail [tail‐OPTIONS]
//
//
//        −t, −−topic (required)
//               Topic to read from.
//
//        −p, −−partiition <default: "0">
//               Topic partition to consume.
//
//        −o, −−offset <default: "0">
//               Topic offset to resume from.
//
//        −h, −−help
//               Show this help message
//
//
//
package main
