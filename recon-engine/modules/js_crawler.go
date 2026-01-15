package modules

import (
	"io"
	"regexp"
	"strings"

	"ai-autonomous-redteam/recon-engine/model"
	scopeguard "ai-autonomous-redteam/scope-guard"
)

type JSCrawler struct{}

func (j *JSCrawler) Name() string {
	return "js-crawler"
}

func (j *JSCrawler) Run(domain string, result *model.ReconResult) error {
	client := scopeguard.NewSafeClient()

	// ---- Step 1: Fetch homepage ----
	resp, err := client.Get("https://" + domain)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	html := string(body)

	// ---- Step 2: Extract JS files ----
	jsFiles := extractJSFiles(html)
	result.JSFiles = jsFiles

	// ---- Step 3: Fetch & analyze JS ----
	for _, js := range jsFiles {
		jsURL := normalizeJSURL(domain, js)

		resp, err := client.Get(jsURL)
		if err != nil {
			continue
		}

		jsBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		endpoints := extractEndpoints(string(jsBody))
		result.Endpoints = append(result.Endpoints, endpoints...)
	}

	return nil
}
func extractJSFiles(html string) []string {
	re := regexp.MustCompile(`<script[^>]+src=["']([^"']+)["']`)
	matches := re.FindAllStringSubmatch(html, -1)

	var files []string
	for _, m := range matches {
		files = append(files, m[1])
	}
	return files
}

func normalizeJSURL(domain, src string) string {
	if strings.HasPrefix(src, "http") {
		return src
	}
	if strings.HasPrefix(src, "//") {
		return "https:" + src
	}
	if strings.HasPrefix(src, "/") {
		return "https://" + domain + src
	}
	return "https://" + domain + "/" + src
}

func extractEndpoints(js string) []string {
	re := regexp.MustCompile(`["'](/api/[a-zA-Z0-9/_\-]+)["']`)
	matches := re.FindAllStringSubmatch(js, -1)

	var eps []string
	for _, m := range matches {
		eps = append(eps, m[1])
	}
	return eps
}
