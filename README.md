## Usage

Build greetclient

```bash
$ make greetclient
```

Build greetserver

```bash
$ make greetserver
```

Test loop echoing

```bash
# Start the server first
$ ./output/greetserver
# Then start the client
$ ./output/greetclient loopecho
```

Test stream ticking

```bash
# Start the server first
$ ./output/greetserver
# Then start the client
$ ./output/greetclient tick
```