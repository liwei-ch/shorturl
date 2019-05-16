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

## Example


<pre><font color="#8AE234"><b>➜  </b></font><font color="#34E2E2"><b>shorturl</b></font> <font color="#729FCF"><b>git:(</b></font><font color="#EF2929"><b>master</b></font><font color="#729FCF"><b>) </b></font><font color="#FCE94F"><b>✗</b></font> curl &quot;http://rammiah.org:8080/generate?url=https://baidu.com&quot;
{&quot;surl&quot;:&quot;rammiah.org:8080/to/MnYAAMcf&quot;,&quot;url&quot;:&quot;https://baidu.com&quot;}<span style="background-color:#3B3E45"><font color="#FFFFFF"></font></span>                                                                                                                                            <font color="#8AE234"><b>➜  </b></font>
<font color="#34E2E2"><b>shorturl</b></font> <font color="#729FCF"><b>git:(</b></font><font color="#EF2929"><b>master</b></font><font color="#729FCF"><b>) </b></font><font color="#FCE94F"><b>✗</b></font> curl rammiah.org:8080/to/MnYAAMcf
&lt;a href=&quot;https://baidu.com&quot;&gt;Moved Permanently&lt;/a&gt;.
<font color="#8AE234"><b>➜  </b></font><font color="#34E2E2"><b>shorturl</b></font> <font color="#729FCF"><b>git:(</b></font><font color="#EF2929"><b>master</b></font><font color="#729FCF"><b>) </b></font><font color="#FCE94F"><b>✗</b></font> </pre>

## Attention
url must be starts with `http://` or `https://`. Or it will be parsed as a link belong to your domain

Example:
```bash
url: baidu.com
server: rammiah.org
redirect to: rammiah.org/to/baidu.com
```