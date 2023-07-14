package ranobes

import (
	"github/cam-per/nst"
	"io"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func (src *source) ID() string {
	return host.Hostname()
}

func (src *source) Search(title string) ([]nst.BookHeader, error) {
	var result []nst.BookHeader

	wd, err := nst.ChromeDriver()
	if err != nil {
		return result, err
	}
	defer wd.Quit()

	if err := wd.Get("https://ranobes.com/index.php?do=search"); err != nil {
		return result, err
	}

	wd.ExecuteScript(`
		function post(path, params, method='post') {
			const form = document.createElement('form');
			form.method = method;
			form.action = path;

			for (const key in params) {
			if (params.hasOwnProperty(key)) {
				const hiddenField = document.createElement('input');
				hiddenField.type = 'hidden';
				hiddenField.name = key;
				hiddenField.value = params[key];

				form.appendChild(hiddenField);
			}
			}

			document.body.appendChild(form);
			form.submit();
		}

		post('/index.php?do=search', {
			do: 'search',
			subaction: 'search',
			search_start: 0,
			full_search: 1,
			result_from: 1,
			story: arguments[0],
			titleonly: 3
		});
	`, []interface{}{
		title,
	})

	var (
		articles []selenium.WebElement
	)

	wd.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		searchResult, err := wd.FindElement(selenium.ByCSSSelector, "div.search_result_num.grey")
		if err != nil {
			return false, nil
		}

		text, _ := searchResult.Text()
		if strings.TrimSpace(text) == "" {
			return false, nil
		}

		articles, _ = wd.FindElements(selenium.ByCSSSelector, "article.block.story.shortstory.mod-poster")
		return true, nil
	}, 10*time.Second)

	for _, article := range articles {
		a, err := article.FindElement(selenium.ByCSSSelector, "h2.title > a")
		if err != nil {
			continue
		}

		href, err := a.GetAttribute("href")
		if err != nil {
			continue
		}

		text, err := a.Text()
		if err != nil {
			continue
		}

		result = append(result, nst.BookHeader{
			Name: text,
			URL:  href,
		})
	}

	return result, nil
}

func (src *source) Info(path string) (nst.BookInfo, error) {
	var result nst.BookInfo

	return result, nil
}

func (src *source) Save(w io.Writer, path string) error {

	return nil
}

func (src *source) OnProgress(handler nst.HandlerProgress) {
	src.handlerProgress = handler
}
