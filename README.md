# Portable HTTP Server

## Features

- [x] Directory Listing
- [x] Directory Explore
- [x] Upload file `/upload`

## Installation 

If you already have Go installed on your computer, use :

```sh
go get github.com/codenoid/portable-http-server
```

## Usage

```sh
$ portable-http-server -help
  -port string
        -port <wanted-port> (default "3000")
# upload a file
$ curl -F "file=@path/to/file.jpg" 127.0.0.1:3000/upload
```
