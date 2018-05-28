// WebServerSettings_test.go
package main

import (
	"testing"
)

func TestWebServerSettings_init_set_get(t *testing.T) {
	sett := initWebServerSettings("~/go/src/www/")
	sett.setCssPath("css/")
	sett.setTmplPath("tmpl/")
	sett.setDataPath("data/")

	if sett.getWWWPath() != "~/go/src/www/" {
		t.Error("Expected ~/go/src/www/, got", sett.getWWWPath())
	}

	if sett.getCssPath() != "~/go/src/www/css/" {
		t.Error("Expected ~/go/src/www/css/, got", sett.getCssPath())

	}

	if sett.getTmplPath() != "~/go/src/www/tmpl/" {
		t.Error("Expected ~/go/src/www/tmpl/, got", sett.getTmplPath())
	}

	if sett.getDataPath() != "~/go/src/www/data/" {
		t.Error("Expected ~/go/src/www/data/, got", sett.getDataPath())
	}
}
