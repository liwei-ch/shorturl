package url

func (u *UrlServer) ParseSurl(surl string) (string, bool) {
	// 内存中不存放数据
	// 参数为短链接的最后一段
	return u.c.GetUrl(surl)
}
