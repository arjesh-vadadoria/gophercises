package main

import (
	"fmt"
	"scrum"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	openChrome()
}

func readFileSendMail() {
	data := demo.YamlHandler("scrumnote.yaml")
	projects := data[1].Projects

	body := "1. What was I accomplished today?\n"

	for _, k := range projects {
		body += fmt.Sprintf("%s: %s\n", k.Name, k.Time)
		for _, j := range k.Tasks {
			body += fmt.Sprintf("* %s\n", j)
		}
		body += fmt.Sprint("\n")
	}
	body += fmt.Sprintf("2. What will I do Tomorrow?\n* Continue with work\n\n")
	body += fmt.Sprintf("3. What challenges did I face today?\n\n4. What did I learn today?\n\n5. Have you pushed the code in Gitlab? (Yes/No), if No then Reason?\n* Yes")

	date := data[0].Date
	fmt.Println("body: \n", body)

	mail := demo.Mail{
		Sender:  "arjesh.vadadoriya@grewon.com",
		To:      []string{"arjeshvadadoriya123@gmail.com"},
		Bcc:     []string{"2274movies@gmail.com"},
		Subject: date,
		Body:    body,
	}
	_ = mail
	// demo.EmailController(mail)
}

func openChrome() {
	// Run Chrome browser
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		panic(err.Error())
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		// "--headless",  // comment out this line to see the browser
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	driver.Get("https://www.google.com")
}
