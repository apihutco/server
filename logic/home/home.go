package home

import (
	"fmt"

	"apihut-server/config"
)

func Page() string {
	return fmt.Sprintf("APIHut.\nDocs: %s", config.Conf.Site.DocsUrl)
}
