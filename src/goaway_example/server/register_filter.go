package server

import (
	"github.com/labstack/gommon/log"
	"strings"
	"gateway/src/goaway/core"
	"gateway/src/goaway/util"
)

type BaseUriServiceFilter struct {
	core.BaseFilter
	Uri string
}

func (b *BaseUriServiceFilter) Matches(uri string) bool {
	return strings.HasPrefix(uri, b.Uri)
}

func NewBaseServiceFilter(uri string, filterName string) core.Filter {
	uri, err := util.NormalizeUri(uri)
	if err != nil {
		log.Printf("Invalid uri: %s", uri)
	}
	switch filterName {
	case "MSDOWNLOAD":
		return &MsDownloadFilter{BaseUriServiceFilter{Uri: uri},}
	case "FILTERTEXT":
		return &TextFilter{BaseUriServiceFilter{Uri: uri},}
	case "NOTJSON":
		return &NoneJsonFilter{BaseUriServiceFilter{Uri: uri},}
	case "UPDATE_FLIGHT":
		return &UpdateFlightFilter{BaseUriServiceFilter{Uri: uri},}
	case "RIGHTS":
		return &AirportRightsFilter{BaseUriServiceFilter{Uri: uri},}
	}
	log.Errorf("Unknown filter name: %s", filterName)
	return nil
}
