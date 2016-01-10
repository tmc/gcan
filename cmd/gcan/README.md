gcan
====

[NAME](#NAME)
 [SYNOPSIS](#SYNOPSIS)
 [DESCRIPTION](#DESCRIPTION)
 [OPTIONS](#OPTIONS)
 [COMMANDS](#COMMANDS)

------------------------------------------------------------------------

NAME
----

gcan − gcan cli

SYNOPSIS
--------

**gcan** \[OPTIONS\]

DESCRIPTION
-----------

OPTIONS
-------

**−v**, **−−verbose**

Be verbose.

**−s**, **−−bootstrap-server** &lt;default: *":4226"*&gt;

Server addr to bootstrap with.

**−h**, **−−help**

Show this help message

COMMANDS
--------

**cat**
 cat consumes messages from stdin

valid compression codecs: NONE, SNAPPY, ZLIB

**Usage**: gcan \[OPTIONS\] cat \[cat-OPTIONS\]

**−t**, **−−topic** (*required*)

Topic to send to.

**−k**, **−−keySeparator** &lt;default: *"\\t"*&gt;

Value to separate key from values in lines

**−c**, **−−compression** &lt;default: *"NONE"*&gt;

Compression codec to use.

**−h**, **−−help**

Show this help message

**tail**
 tail consumes messages from gcan and produces them on stdout

**Usage**: gcan \[OPTIONS\] tail \[tail-OPTIONS\]

**−t**, **−−topic** (*required*)

Topic to read from.

**−p**, **−−partiition** &lt;default: *"0"*&gt;

Topic partition to consume.

**−o**, **−−offset** &lt;default: *"0"*&gt;

Topic offset to resume from.

**−h**, **−−help**

Show this help message

------------------------------------------------------------------------
