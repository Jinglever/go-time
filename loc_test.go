package jgtime

import "testing"

func TestLoadLocation(t *testing.T) {
	tzCache := NewLocationCache()

	// 测试已知的时区
	knownTimeZones := []string{"Europe/Paris", "America/New_York", "Asia/Shanghai"}
	for _, tz := range knownTimeZones {
		_, err := tzCache.LoadLocation(tz)
		if err != nil {
			t.Errorf("Failed to load location '%s': %v", tz, err)
		}
	}

	// 测试一个未知的时区
	_, err := tzCache.LoadLocation("Unknown/Location")
	if err == nil {
		t.Error("Expected an error for 'Unknown/Location', but did not get one")
	}
}
