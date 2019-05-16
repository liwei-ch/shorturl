package url

import "shorturl/url/cache"

func ParseSurl(surl string) (string, bool) {
	// 内存中不存放数据
	// 参数为短链接的最后一段
	return cache.GetRecord(surl)
}
