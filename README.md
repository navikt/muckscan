# muckscan

Scan Git repositories for secrets. Muckscan uses
[Truffle Hog](https://github.com/dxa4481/truffleHog)
to discover commits containing

* high entropy strings,
* birth identification numbers (fødselsnummer),
* NAV idents,
* RSA private keys,
* and various other OAuth and API keys.

## Usage

Mount your git repository to the Docker container's `/data` directory, and run
the container:

```
docker run --rm -it -v /path/to/git/repository:/data navikt/muckscan
```

## Building

```
make
```
