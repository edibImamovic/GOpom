package GOpom

import "github.com/DATA-DOG/godog"

func (b *BrowserSteps) googleSearch(s *godog.Suite) {
	s.Step(`^I am a anonymous user$`, func() error { return b.GetWebDriver().DeleteAllCookies() })
	s.Step(`^I navigate to "([^"]*)"$`, b.iNavigateTo)
	s.Step(`^I write "([^"]*)" to "([^"]*)" `+ByOption+`$`, b.iWriteTo)
	s.Step(`^I click  button$`, b.iClick)
	s.Step(`^I submit "([^"]*)" `+ByOption+`$`, b.iSubmit)
	s.Step(`^I wait for (\d+) (milliseconds|millisecond|seconds|second)$`, iWaitFor)
	s.Step(`^I should see "([^"]*)" in link text$`, b.iShouldSeePageTitleAs)

}
