package GOpom

import (
	"fmt"
	"time"
)

func (b *BrowserSteps) iNavigateTo(site string) error {
	u, err := b.GetURL(site)
	if err != nil {
		return err
	}
	return b.GetWebDriver().Get(u.String())
}

func (b *BrowserSteps) iWriteTo(text, selector, by string) error {
	// Click the element
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}

	err = element.Clear()
	if err != nil {
		return err
	}
	return element.SendKeys(text)
}

func (b *BrowserSteps) iClick() error {

	// Submit the element
	element, err := b.GetWebDriver().FindElement("name", "btnK")
	if err != nil {
		return err
	}
	return element.Click()
}

func (b *BrowserSteps) iSubmit(selector, by string) error {
	// Submit the element
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}
	return element.Submit()
}

func (b *BrowserSteps) iShouldSeePageTitleAs(expectedTitle string) error {
	title, err := b.GetWebDriver().Title()
	if err != nil {
		return err
	}

	if expectedTitle != title {
		return fmt.Errorf("Title does not match. Expected '%s', Found '%s'",
			expectedTitle, title)
	}
	return nil
}

func iWaitFor(amount int, unit string) error {
	u := time.Second
	fmt.Printf("Waiting for %d %s", amount, unit)
	time.Sleep(u * time.Duration(amount))
	return nil
}
