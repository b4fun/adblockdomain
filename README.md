# adblockdomain

Parse adblock rules domain from adblock rules.

## Usage

```
$ ./adblockdomain EASY_LIST.txt
example.com
```

show exception domains:

```
$ ./adblockdomain -e EASY_LIST.txt
exception.example.com
```

decode baes64 content first

```
$ ./adblockdomain -b64 B64_ENCODED_EASY_LIST.txt
example.com
```

## LICENSE

MIT
