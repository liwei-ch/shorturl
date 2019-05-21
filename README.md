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

## Usage

```bash
POST <server:port>/generate
data body is url to shorten
return {"surl": "short url", "url": "url you input"}

GET surl
will be redirected to url
```

## Example


<pre><font color="#55FF55"><b>➜  </b></font><font color="#55FFFF"><b>shorturl</b></font> <font color="#5555FF"><b>git:(</b></font><font color="#FF5555"><b>master</b></font><font color="#5555FF"><b>) </b></font><font color="#FFFF55"><b>✗</b></font> curl -X POST rammiah.org:8080/generate -d &quot;https://www.baidu.com&quot;
{&quot;surl&quot;:&quot;http://rammiah.org:8080/to/fiCEerXQ&quot;,&quot;url&quot;:&quot;https://www.baidu.com&quot;}<span style="background-color:#D3DAE3"><font color="#404552"><b>%</b></font></span>                                               
<font color="#55FF55"><b>➜  </b></font><font color="#55FFFF"><b>shorturl</b></font> <font color="#5555FF"><b>git:(</b></font><font color="#FF5555"><b>master</b></font><font color="#5555FF"><b>) </b></font><font color="#FFFF55"><b>✗</b></font> curl http://rammiah.org:8080/to/fiCEerXQ
&lt;a href=&quot;https://www.baidu.com&quot;&gt;Moved Permanently&lt;/a&gt;.

<font color="#55FF55"><b>➜  </b></font><font color="#55FFFF"><b>shorturl</b></font> <font color="#5555FF"><b>git:(</b></font><font color="#FF5555"><b>master</b></font><font color="#5555FF"><b>) </b></font><font color="#FFFF55"><b>✗</b></font> 
</pre>
## Attention
url must be starts with `http://` or `https://`. Or it will be parsed as a link belong to your domain

Example:
```bash
url: baidu.com
server: rammiah.org
redirect to: rammiah.org/to/baidu.com
```
