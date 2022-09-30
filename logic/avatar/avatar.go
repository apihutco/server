package avatar

import (
	"apihut-server/config"
	"apihut-server/models"
	"fmt"
	"github.com/nullrocks/identicon"
	"os"
	"path"
	"time"
)

func NewAvatar(req *models.AvatarReq) error {
	gen, err := identicon.New(req.GetNamespace(), req.GetSize(), req.GetDensity())
	if err != nil {
		return err
	}
	ident, err := gen.Draw(req.GetHash())
	if err != nil {
		return err
	}

	filePath := path.Join(config.Share.File.Avatar, fmt.Sprintf("%d.png", time.Now().UnixNano()))
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return err
	}

	err = ident.Png(32, f)
	if err != nil {
		return err
	}
	return nil
}
