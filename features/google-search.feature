Feature: Google Search

@remote
    Scenario: Submit a form
        Given I navigate to "https://google.com/"
        When I write "Google" to "q" name
        And I click  button
        And I wait for 2 seconds
        Then I should see "Google - Google Pretraživanje" in link text
