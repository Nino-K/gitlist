[![Build Status](https://travis-ci.org/Nino-K/gitlist.svg?branch=master)](https://travis-ci.org/Nino-K/gitlist)

#GitList

Helper tool to get details about golang repositories on Github

## Install:

`make install`

## Usage:

#### - Get all repos that include a given keyword

`gitlist show tail`

```
+----+--------------+-------------------------------------+
| #  |     NAME     |                 URL                 |
+----+--------------+-------------------------------------+
|  1 | tail         | github.com/hpcloud/tail             |
|  2 | tailf        | github.com/aybabtme/tailf           |
|  3 | tail         | github.com/go-zoo/tail              |
|  4 | tail         | github.com/mangalaman93/tail        |
|  5 | tail         | github.com/cwiggers/tail            |
|  6 | tail         | github.com/errnoh/tail              |
|  7 | tail         | github.com/timperman/tail           |
|  8 | tail         | github.com/paulstuart/tail          |
+----+--------------+-------------------------------------+

```
#### - Get details about a specific repo

`gitlist detail 1`

```
+------+--------------------------------+-----------------------------------+-------------------------------------+---------------------------------+
| NAME |          DESCRIPTION           |              GITURL               |              CLONEURL               |             SSHURL              |
+------+--------------------------------+-----------------------------------+-------------------------------------+---------------------------------+
| tail | Go package for reading from    | git://github.com/hpcloud/tail.git | https://github.com/hpcloud/tail.git | git@github.com:hpcloud/tail.git |
|      | continously updated files      |                                   |                                     |                                 |
|      | (tail -f)                      |                                   |                                     |                                 |
+------+--------------------------------+-----------------------------------+-------------------------------------+---------------------------------+
```

