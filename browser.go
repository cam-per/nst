package nst

import (
	"fmt"
	"github/cam-per/nst/config"
	"os"

	"github.com/tebeka/selenium"
)

var (
	SeleniumService *selenium.Service
)

func InitSelenium() error {
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr),
	}

	if config.Selenium.XVFB.Enabled {
		opts = append(opts, selenium.StartFrameBuffer())
	}

	if config.Selenium.Drivers.ChromePath != "" {
		opts = append(opts, selenium.ChromeDriver(config.Selenium.Drivers.ChromePath))
	}

	if config.Selenium.Drivers.GeckoPath != "" {
		opts = append(opts, selenium.GeckoDriver(config.Selenium.Drivers.GeckoPath))
	}

	selenium.SetDebug(true)

	var err error
	SeleniumService, err = selenium.NewSeleniumService(config.Selenium.ServerPath, config.Selenium.Port, opts...)

	return err
}

func FirefoxDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{
		"browserName": "firefox",
	}

	url := fmt.Sprintf("http://localhost:%d/wd/hub", config.Selenium.Port)
	return selenium.NewRemote(caps, url)
}

func ChromeDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	url := fmt.Sprintf("http://localhost:%d/wd/hub", config.Selenium.Port)
	return selenium.NewRemote(caps, url)
}

func SeleniumStop() {
	SeleniumService.Stop()
}
