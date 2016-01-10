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

**Usage**: gcan \[OPTIONS\] cat \[cat-OPTIONS\]

**−t**, **−−topic** (*required*)

Topic to send to.

**−k**, **−−keySeparator** &lt;default: *"\\t"*&gt;

Value to separate key from values in lines

**−h**, **−−help**

Show this help message

------------------------------------------------------------------------
