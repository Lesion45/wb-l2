package downloader

import (
	"dev09/models"
	"dev09/utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

type Downloader struct {
	Website *models.Website
}

func NewDownloader(website *models.Website) *Downloader {
	return &Downloader{Website: website}
}

func (d *Downloader) Download() error {
	err := d.downloadHTML()
	if err != nil {
		return err
	}

	for _, resourceURL := range d.Website.Assets {
		err := d.downloadAndSaveResource(resourceURL)
		if err != nil {
			fmt.Printf("Failed to download resource %s: %v\n", resourceURL, err)
		}
	}

	err = d.updateHTMLWithLocalLinks()
	if err != nil {
		return err
	}

	err = utils.SaveHTML(d.Website.DirName, d.Website.HTML)
	if err != nil {
		return err
	}

	return nil
}

func (d *Downloader) downloadHTML() error {
	resp, err := http.Get(d.Website.URL)
	if err != nil {
		return fmt.Errorf("failed to download HTML: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTML body: %v", err)
	}

	d.Website.HTML = string(body)

	d.Website.Assets = extractResources(d.Website.HTML)

	err = os.MkdirAll(d.Website.DirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create directory: %v", err)
	}

	return nil
}

func (d *Downloader) downloadAndSaveResource(resourceURL string) error {
	parsedURL, err := url.Parse(resourceURL)
	if err != nil {
		return fmt.Errorf("invalid resource URL: %v", err)
	}

	resourcePath := filepath.Join(d.Website.DirName, utils.GetFileName(parsedURL))
	err = os.MkdirAll(filepath.Dir(resourcePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create directory for resource: %v", err)
	}

	resp, err := http.Get(resourceURL)
	if err != nil {
		return fmt.Errorf("failed to download resource %s: %v", resourceURL, err)
	}
	defer resp.Body.Close()

	outFile, err := os.Create(resourcePath)
	if err != nil {
		return fmt.Errorf("failed to create file for resource %s: %v", resourceURL, err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save resource %s: %v", resourceURL, err)
	}

	return nil
}

func (d *Downloader) updateHTMLWithLocalLinks() error {
	re := regexp.MustCompile(`(src|href)=[\'\"](https?://[^\s\'\"]+)`)
	d.Website.HTML = re.ReplaceAllStringFunc(d.Website.HTML, func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 3 {
			return match
		}

		resourceURL := matches[2]
		parsedURL, err := url.Parse(resourceURL)
		if err != nil {
			return match
		}
		fileName := utils.GetFileName(parsedURL)

		return fmt.Sprintf("%s=%q", matches[1], filepath.Join(d.Website.DirName, fileName))
	})

	return nil
}

func extractResources(html string) []string {
	var urls []string
	re := regexp.MustCompile(`(src|href)=[\'\"](https?://[^\s\'\"]+)`)
	matches := re.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		urls = append(urls, match[2])
	}
	return urls
}
