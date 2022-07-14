// This file contains common regex filters
package commonregex

const (
	Ip     = `\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`
	Domain = `\b((xn--)?[a-z0-9][a-z0-9-_]{0,61}[a-z0-9]{0,1}\.){1,9}(xn--)?([a-z0-9\-]{1,61}|[a-z0-9-]{1,30}\.[a-z]{2,})\b`
	Sha1   = `\b[a-f0-9]{40}\b`
	Md5    = `\b[a-f0-9]{32}\b`
	Sha256 = `\b[a-f0-9]{64}\b`
)
