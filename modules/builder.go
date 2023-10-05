package modules

import (
	dianomiPrebid_seat_module "github.com/prebid/prebid-server/modules/dianomi/prebid_seat_module"
	prebidOrtb2blocking "github.com/prebid/prebid-server/modules/prebid/ortb2blocking"
)

// builders returns mapping between module name and its builder
// vendor and module names are chosen based on the module directory name
func builders() ModuleBuilders {
	return ModuleBuilders{
		"dianomi": {
			"prebid_seat_module": dianomiPrebid_seat_module.Builder,
		},
		"prebid": {
			"ortb2blocking": prebidOrtb2blocking.Builder,
		},
	}
}
