// +build !windows

package fastzip

import (
	"archive/zip"
	"io"
	"math/big"
	"syscall"

	"github.com/saracen/zipextra"
)

func (a *Archiver) createHeader(hdr *zip.FileHeader) (io.Writer, error) {
	stat, ok := hdr.FileInfo().Sys().(*syscall.Stat_t)
	if ok {
		hdr.Extra = append(hdr.Extra, zipextra.NewInfoZIPNewUnix(big.NewInt(int64(stat.Uid)), big.NewInt(int64(stat.Gid))).Encode()...)
	}

	return a.zw.CreateHeader(hdr)
}
