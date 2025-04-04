# Go HTTP Server

A HTTP server that's capable of handling simple GET/POST requests, serving files, compression and handling multiple concurrent connections.

## 🚀 Quick Start

Ensure you have either [Docker](https://www.docker.com/get-started) or [Go](https://golang.org/doc/install) environments set up. Additionally, make sure you have you have an API testing tool installed. I used [cURL](https://curl.se/), but you can use anything you like(eg. [Postman](https://www.postman.com/), [Thunder Client](https://www.thunderclient.com/), ...).

### Clone the project:

```bash
git clone https://github.com/danilovict2/go-http-server.git
cd go-http-server
```

### Run with Docker:

```bash
docker build .
docker container run -p "4221:4221" <IMAGE ID>
```

### Run locally:

```bash
./your_program.sh
```

## 📖 Usage

### Available endpoints

* `GET /`
* `GET /echo/{str}`
* `GET /user-agent`
* `GET /files/{filename}`
* `POST /files/{filename}`

## Examples

### 404 - Not Found

If you send a request to an endpoint that is not listed in the available endpoints, the server will respond with a `404 Not Found` status code.

```bash
$ curl -i http://localhost:4221/abcdefg
HTTP/1.1 404 Not Found
```

### Echo

The `/echo/{str}` endpoint takes a path parameter and responds with a body containing the same value.

```bash
$ curl http://localhost:4221/echo/abc
abc
```

### User-Agent

The `/user-agent` endpoint reads the `User-Agent` request header and returns it in the response body.

```bash
$ curl --header "User-Agent: foobar/1.2.3" http://localhost:4221/user-agent
foobar/1.2.3
```

### Files

To read/create files, provide a `--directory` that specifies the directory where the files are stored, as an absolute path.

```bash
# With Docker
docker container run -p "4221:4221" <IMAGE ID> --directory /tmp/

# Locally
./your_program.sh --directory /tmp/
```

The `GET /files/{filename}` endpoint takes a filename parameter and will respond with one of the following status codes:
- `200 OK` along with the file's contents if the file is successfully retrieved.
- `404 Not Found` if the specified file does not exist.
- `500 Internal Server Error` if an unexpected error occurs during the operation.

```bash
$ echo -n 'Hello, World!' > /tmp/foo
$ curl -i http://localhost:4221/files/foo
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: application/octet-stream

Hello, World!
```

The `POST /files/{filename}` endpoint allows you to create a file named `{filename}` with the content provided in the request body. The server will respond with one of the following status codes:

- `201 Created`: The file was successfully created.
- `500 Internal Server Error`: An unexpected error occurred during the operation.

```bash
$ curl -i --data "12345" -H "Content-Type: application/octet-stream" http://localhost:4221/files/file_123
HTTP/1.1 201 Created

$ cat /tmp/file_123
12345
```

### Compression

### Compression

To enable response body compression, include the `Accept-Encoding` header in your request with the desired compression schemes. The server will respond with a `Content-Encoding` header indicating the compression scheme it applied, or none if no supported scheme was specified.

**Note:** Currently, only the `gzip` compression scheme is supported.

```bash
$ curl -H "Accept-Encoding: gzip" http://localhost:4221/echo/abc -o - | gunzip
abc
```

## 🤝 Contributing

### Build the project

```bash
go build -o server app/*.go
```

### Run the project

```bash
./server
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `master` branch.