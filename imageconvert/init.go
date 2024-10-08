package imageconvert

import "github.com/davidbyttow/govips/v2/vips"

func init() {
	config := vips.Config{
		ConcurrencyLevel: 2,
		MaxCacheSize:     100,
		CollectStats:     false,
	}
	vips.Startup(&config)
	vips.LoggingSettings(LoggingHandlerFunction, vips.LogLevelCritical)
}
