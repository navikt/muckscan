# muckscan

Scan Git repositories for secrets. Muckscan uses
[Truffle Hog](https://github.com/dxa4481/truffleHog)
to discover commits containing

* high entropy strings,
* birth identification numbers (fødselsnummer),
* NAV idents,
* RSA private keys,
* and various other OAuth and API keys.

The configuration specific to NAV can be found in the
[sensitive regex](sensitive-regex.json) file.

This repository also contains a [wrapper for Truffle Hog](truffletool.go) that
distills its output to make it more readable and easier to work with in case of
huge repositories with a lot of sensitive history.

## Usage

### Public repository
Give the url to your git repository as first parameter to muckscan:
```
docker run --rm -it navikt/muckscan https://github.com/navikt/muckscan.git
```

### Private repository
If you need to access a private repository, create a file `~/.git-credentials`
in the following format:
```
https://<username>:<personal_access_token>@github.com/
```

Then, run the container as you would with a public repository, and mount your
credentials file into the container:
```
docker run -v ~/.git-credentials:/root/.git-credentials (...)
```

### Local repository
A local repository can be mounted in this manner: 

```
docker run --rm -it -v /path/to/git/repository:/data navikt/muckscan /data
```

## Output

If Muckscan thinks your repository is clean, the command will give no output.

When commits containing sensitive data are found, Muckscan will give an output
similar to this:

```
[*] file.....: foo
[*] commits..: 2d0ee374015ce7c12adfd1c872bd60f4536e7e36
[*] branches.: master (1)
[*] reasons..: Fødselsnummer (1)
12125678910
---
[*] file.....: key
[*] commits..: 5db78374fed1483550456fa9d96de8a79e64ba0c
[*] branches.: master (2)
[*] reasons..: RSA private key (1), High Entropy (1)
(...snip...)
-----BEGIN RSA PRIVATE KEY-----
(...snip...)
---
[*] file.....: key.pub
[*] commits..: 5db78374fed1483550456fa9d96de8a79e64ba0c
[*] branches.: master (1)
[*] reasons..: High Entropy (1)
AAAAB3NzaC1yc2EAAAADAQABAAABAQDkZMeH1wsHd6M5Q7VaA1KPs1Oia78embqES7Lat0U+VF60q2p2...(292 characters truncated)
```

### Exit status

Muckscan exits with one of the following values:

* `0` when no sensitive commits are found.
* `1` otherwise.

## Building

```
make
```
