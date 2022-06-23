package upload

import (
	"gostars/global"
	"io"
)

type FtpOptions struct {
}

func FtpCreateDir(dir string) {
	global.GFtp.Cwd("/var/ftp")
	global.GFtp.Mkd(dir)
}

func FtpCreateFile(path string, file io.Reader) {
	global.GFtp.Cwd("/var/ftp")
	global.GFtp.Stor(path, file)
}
