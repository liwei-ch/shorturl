# ShortUrl Service

## Get Start

```shell
mkdir -p $GOPATH/src/github.com/rammiah
cd $GOPATH/src/github.com/rammiah
git clone https://github.com/rammiah/shorturl
cd shorturl
go build
./shorturl
```
You have to config yours redis server and mysql server

see in `url/cache/redis.go` and `url/db/mysql.go`.

## Usage

```bash
GET <server:port>/generate?url=<url to shorten>
return {"surl": "short url", "url": "url you input"}

GET surl
will be redirected to url
```
