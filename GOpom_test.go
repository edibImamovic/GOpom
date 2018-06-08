package GOpom

import (
	"os"
	"testing"
	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"

)



func FeatureContext(s *godog.Suite) {

	NewBrowserSteps(s, selenium.Capabilities{"browserName": "chrome"}, "")
}

func TestMain(m *testing.M) {
	status := godog.Run("GOpom", FeatureContext)
	os.Exit(status)
}
