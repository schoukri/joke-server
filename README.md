# joke-server
Get a random joke about a random person (except Chuck Norris).

## Author

Sam Choukri
`(sam at choukri dot net)`

## Running joke-server with Go

`joke-server` is written in Go. You can build and run the server directly from the source code if you have Go installed.

Change directories into the `joke-server` directory:

```
cd joke-server
```

Run the tests to make sure everything passes:

```
go test -v ./...
```

Build the `joke-server` server:

```
go build -o joke-server
```

That will produce a `joke-server` executable file inside of the same directory:


Run the `joke-server` server:

```
./joke-server
```

If the server starts up successfully, you will see a message like this:

```
2020/04/29 22:17:49 Listening on addr :5000
```

By default, the server will run on port 5000.

You can request a random joke in your web browser with the following url [http://localhost:5000/](http://localhost:5000/).

Or you can use curl:

```
curl http://localhost:5000/
```

To stop the server, press Ctrl+C (^C).

## Running joke-server with Docker

You can also run `joke-server` in Docker if you prefer.

Change directories into the `joke-server` directory:

```
cd joke-server
```

Build the `joke-server` docker image:

```
docker-compose build
```

The first time you run this command, docker may take several minutes to download all the required files.

After the build is finished, run `joke-server` in docker:

```
docker-compose up
```

That should start `joke-server` in docker and you should be able to access it at [http://localhost:5000/](http://localhost:5000/).

To stop the server, press Ctrl+C (^C).
