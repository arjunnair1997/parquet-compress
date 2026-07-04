package reportpdf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var (
	imagePattern = regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	linkPattern  = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
)

// WriteMarkdownAndPDF writes markdownPath and a sibling PDF with the same base
// name. PDF generation uses a local Chrome/Chromium installation.
func WriteMarkdownAndPDF(markdownPath string, data []byte) error {
	if err := os.WriteFile(markdownPath, data, 0o644); err != nil {
		return err
	}
	_, err := Generate(markdownPath)
	return err
}

func PathForMarkdown(markdownPath string) string {
	return strings.TrimSuffix(markdownPath, filepath.Ext(markdownPath)) + ".pdf"
}

func Generate(markdownPath string) (string, error) {
	absMarkdown, err := filepath.Abs(markdownPath)
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(absMarkdown)
	if err != nil {
		return "", err
	}
	chromePath, err := findChrome()
	if err != nil {
		return "", err
	}

	pdfPath := PathForMarkdown(absMarkdown)
	if err := os.Remove(pdfPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	tmpDir, err := os.MkdirTemp("", "parquet-reportpdf-*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmpDir)

	htmlPath := filepath.Join(tmpDir, "report.html")
	if err := os.WriteFile(htmlPath, []byte(markdownHTML(data, filepath.Dir(absMarkdown))), 0o644); err != nil {
		return "", err
	}

	profileDir := filepath.Join(tmpDir, "chrome-profile")
	args := []string{
		"--headless",
		"--disable-gpu",
		"--disable-dev-shm-usage",
		"--no-sandbox",
		"--allow-file-access-from-files",
		"--no-pdf-header-footer",
		"--run-all-compositor-stages-before-draw",
		"--user-data-dir=" + profileDir,
		"--print-to-pdf=" + pdfPath,
		fileURI(htmlPath),
	}
	if err := runChromePrint(chromePath, args, pdfPath, markdownPath); err != nil {
		return "", err
	}
	return pdfPath, nil
}

func runChromePrint(chromePath string, args []string, pdfPath, markdownPath string) error {
	cmd := exec.Command(chromePath, args...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Start(); err != nil {
		return err
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	deadline := time.After(2 * time.Minute)
	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	var readyAt time.Time
	var lastSize int64
	for {
		select {
		case err := <-done:
			if err != nil && !pdfReady(pdfPath) {
				return fmt.Errorf("generate PDF from %s: %w\n%s", markdownPath, err, strings.TrimSpace(output.String()))
			}
			return nil
		case <-ticker.C:
			size, ready := pdfSize(pdfPath)
			if ready && size == lastSize {
				if readyAt.IsZero() {
					readyAt = time.Now()
				}
				if time.Since(readyAt) >= time.Second {
					_ = cmd.Process.Kill()
					<-done
					return nil
				}
			} else {
				readyAt = time.Time{}
				lastSize = size
			}
		case <-deadline:
			_ = cmd.Process.Kill()
			<-done
			if pdfReady(pdfPath) {
				return nil
			}
			return fmt.Errorf("generate PDF from %s timed out\n%s", markdownPath, strings.TrimSpace(output.String()))
		}
	}
}

func pdfReady(path string) bool {
	_, ready := pdfSize(path)
	return ready
}

func pdfSize(path string) (int64, bool) {
	info, err := os.Stat(path)
	if err != nil || info.IsDir() || info.Size() == 0 {
		return 0, false
	}
	return info.Size(), true
}

func findChrome() (string, error) {
	for _, env := range []string{"CHROME_PATH", "CHROMIUM_PATH"} {
		if path := strings.TrimSpace(os.Getenv(env)); path != "" {
			if executable(path) {
				return path, nil
			}
		}
	}

	for _, name := range []string{
		"google-chrome",
		"google-chrome-stable",
		"chromium",
		"chromium-browser",
		"chrome",
		"msedge",
		"brave-browser",
	} {
		if path, err := exec.LookPath(name); err == nil {
			return path, nil
		}
	}

	for _, path := range []string{
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/Applications/Chromium.app/Contents/MacOS/Chromium",
		"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
		"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
	} {
		if executable(path) {
			return path, nil
		}
	}

	return "", errors.New("Chrome/Chromium executable not found; set CHROME_PATH or run with --generate-pdf=false")
}

func executable(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir() && info.Mode()&0o111 != 0
}

func markdownHTML(markdown []byte, baseDir string) string {
	body := renderMarkdown(markdown, baseDir)
	return `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<style>
@page { size: Letter; margin: 0.45in; }
html, body { margin: 0; padding: 0; }
body {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Arial, sans-serif;
  color: #111827;
  font-size: 10.5px;
  line-height: 1.42;
}
h1 { font-size: 24px; margin: 0 0 16px; border-bottom: 2px solid #d1d5db; padding-bottom: 8px; }
h2 { font-size: 17px; margin: 22px 0 8px; padding-top: 6px; border-top: 1px solid #e5e7eb; break-after: avoid; page-break-after: avoid; }
h3 { font-size: 13px; margin: 14px 0 6px; break-after: avoid; page-break-after: avoid; }
p { margin: 5px 0 8px; }
ol, ul { margin: 5px 0 10px 22px; padding: 0; }
li { margin: 3px 0; }
code {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 0.94em;
  background: #f3f4f6;
  padding: 1px 3px;
  border-radius: 3px;
}
pre { background: #f3f4f6; padding: 8px; overflow-wrap: anywhere; white-space: pre-wrap; }
pre code { background: transparent; padding: 0; }
table { border-collapse: collapse; width: 100%; margin: 8px 0 12px; font-size: 9px; table-layout: fixed; }
th, td { border: 1px solid #d1d5db; padding: 4px 5px; text-align: left; vertical-align: top; overflow-wrap: anywhere; }
th { background: #f3f4f6; }
img { display: block; max-width: 100%; height: auto; margin: 8px 0 12px; border: 1px solid #e5e7eb; break-inside: avoid; page-break-inside: avoid; }
hr { border: 0; border-top: 1px solid #e5e7eb; margin: 18px 0; }
</style>
</head>
<body>
` + body + `
</body>
</html>
`
}

func renderMarkdown(markdown []byte, baseDir string) string {
	var b strings.Builder
	scanner := bufio.NewScanner(bytes.NewReader(markdown))
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var list string
	inCode := false
	var code strings.Builder
	lines := make([]string, 0, 1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	closeList := func() {
		if list != "" {
			fmt.Fprintf(&b, "</%s>\n", list)
			list = ""
		}
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			if inCode {
				fmt.Fprintf(&b, "<pre><code>%s</code></pre>\n", html.EscapeString(strings.TrimRight(code.String(), "\n")))
				code.Reset()
				inCode = false
			} else {
				closeList()
				inCode = true
			}
			continue
		}
		if inCode {
			code.WriteString(line)
			code.WriteByte('\n')
			continue
		}
		if trimmed == "" {
			closeList()
			continue
		}
		if isTableStart(lines, i) {
			closeList()
			writeTable(&b, lines, &i, baseDir)
			continue
		}
		if strings.HasPrefix(trimmed, "### ") {
			closeList()
			fmt.Fprintf(&b, "<h3>%s</h3>\n", renderInline(strings.TrimSpace(trimmed[4:]), baseDir))
			continue
		}
		if strings.HasPrefix(trimmed, "## ") {
			closeList()
			fmt.Fprintf(&b, "<h2>%s</h2>\n", renderInline(strings.TrimSpace(trimmed[3:]), baseDir))
			continue
		}
		if strings.HasPrefix(trimmed, "# ") {
			closeList()
			fmt.Fprintf(&b, "<h1>%s</h1>\n", renderInline(strings.TrimSpace(trimmed[2:]), baseDir))
			continue
		}
		if strings.HasPrefix(trimmed, "---") && strings.Trim(trimmed, "-") == "" {
			closeList()
			b.WriteString("<hr>\n")
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			if list != "ul" {
				closeList()
				list = "ul"
				b.WriteString("<ul>\n")
			}
			fmt.Fprintf(&b, "<li>%s</li>\n", renderInline(strings.TrimSpace(trimmed[2:]), baseDir))
			continue
		}
		if n := orderedListPrefix(trimmed); n > 0 {
			if list != "ol" {
				closeList()
				list = "ol"
				b.WriteString("<ol>\n")
			}
			fmt.Fprintf(&b, "<li>%s</li>\n", renderInline(strings.TrimSpace(trimmed[n:]), baseDir))
			continue
		}
		closeList()
		fmt.Fprintf(&b, "<p>%s</p>\n", renderInline(trimmed, baseDir))
	}
	closeList()
	if inCode {
		fmt.Fprintf(&b, "<pre><code>%s</code></pre>\n", html.EscapeString(strings.TrimRight(code.String(), "\n")))
	}
	return b.String()
}

func isTableStart(lines []string, i int) bool {
	if i+1 >= len(lines) {
		return false
	}
	return strings.Contains(lines[i], "|") && isTableSeparator(lines[i+1])
}

func isTableSeparator(line string) bool {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "|") {
		line = strings.TrimPrefix(line, "|")
	}
	if strings.HasSuffix(line, "|") {
		line = strings.TrimSuffix(line, "|")
	}
	cells := strings.Split(line, "|")
	if len(cells) == 0 {
		return false
	}
	for _, cell := range cells {
		cell = strings.TrimSpace(cell)
		cell = strings.Trim(cell, ":")
		if len(cell) < 3 || strings.Trim(cell, "-") != "" {
			return false
		}
	}
	return true
}

func writeTable(b *strings.Builder, lines []string, i *int, baseDir string) {
	headers := splitTableRow(lines[*i])
	*i += 1
	fmt.Fprintln(b, "<table>")
	fmt.Fprintln(b, "<thead><tr>")
	for _, h := range headers {
		fmt.Fprintf(b, "<th>%s</th>", renderInline(strings.TrimSpace(h), baseDir))
	}
	fmt.Fprintln(b, "</tr></thead>")
	fmt.Fprintln(b, "<tbody>")
	for *i+1 < len(lines) {
		next := strings.TrimSpace(lines[*i+1])
		if next == "" || !strings.Contains(next, "|") || isTableSeparator(next) {
			break
		}
		*i += 1
		cells := splitTableRow(lines[*i])
		fmt.Fprintln(b, "<tr>")
		for _, cell := range cells {
			fmt.Fprintf(b, "<td>%s</td>", renderInline(strings.TrimSpace(cell), baseDir))
		}
		fmt.Fprintln(b, "</tr>")
	}
	fmt.Fprintln(b, "</tbody>")
	fmt.Fprintln(b, "</table>")
}

func splitTableRow(line string) []string {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "|") {
		line = strings.TrimPrefix(line, "|")
	}
	if strings.HasSuffix(line, "|") {
		line = strings.TrimSuffix(line, "|")
	}
	return strings.Split(line, "|")
}

func orderedListPrefix(s string) int {
	i := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		i++
	}
	if i == 0 || i+1 >= len(s) || s[i] != '.' || s[i+1] != ' ' {
		return 0
	}
	return i + 2
}

func renderInline(s, baseDir string) string {
	parts := splitCodeSpans(s)
	for i := range parts {
		if parts[i].code {
			parts[i].text = "<code>" + html.EscapeString(parts[i].text) + "</code>"
		} else {
			parts[i].text = renderInlineNoCode(parts[i].text, baseDir)
		}
	}
	var b strings.Builder
	for _, part := range parts {
		b.WriteString(part.text)
	}
	return b.String()
}

type inlinePart struct {
	text string
	code bool
}

func splitCodeSpans(s string) []inlinePart {
	var parts []inlinePart
	for {
		start := strings.IndexByte(s, '`')
		if start < 0 {
			if s != "" {
				parts = append(parts, inlinePart{text: s})
			}
			return parts
		}
		if start > 0 {
			parts = append(parts, inlinePart{text: s[:start]})
		}
		rest := s[start+1:]
		end := strings.IndexByte(rest, '`')
		if end < 0 {
			parts = append(parts, inlinePart{text: s[start:]})
			return parts
		}
		parts = append(parts, inlinePart{text: rest[:end], code: true})
		s = rest[end+1:]
	}
}

func renderInlineNoCode(s, baseDir string) string {
	var replacements []replacement
	for _, m := range imagePattern.FindAllStringSubmatchIndex(s, -1) {
		alt := s[m[2]:m[3]]
		src := strings.TrimSpace(s[m[4]:m[5]])
		replacements = append(replacements, replacement{
			start: m[0],
			end:   m[1],
			html:  fmt.Sprintf(`<img src="%s" alt="%s">`, html.EscapeString(resolveAssetURI(baseDir, src)), html.EscapeString(alt)),
		})
	}
	for _, m := range linkPattern.FindAllStringSubmatchIndex(s, -1) {
		overlapsImage := false
		for _, repl := range replacements {
			if m[0] >= repl.start && m[0] < repl.end {
				overlapsImage = true
				break
			}
		}
		if overlapsImage {
			continue
		}
		text := s[m[2]:m[3]]
		href := strings.TrimSpace(s[m[4]:m[5]])
		replacements = append(replacements, replacement{
			start: m[0],
			end:   m[1],
			html:  fmt.Sprintf(`<a href="%s">%s</a>`, html.EscapeString(resolveLinkURI(baseDir, href)), renderInline(text, baseDir)),
		})
	}
	if len(replacements) == 0 {
		return html.EscapeString(s)
	}
	sortReplacements(replacements)
	var b strings.Builder
	pos := 0
	for _, repl := range replacements {
		if repl.start < pos {
			continue
		}
		b.WriteString(html.EscapeString(s[pos:repl.start]))
		b.WriteString(repl.html)
		pos = repl.end
	}
	b.WriteString(html.EscapeString(s[pos:]))
	return b.String()
}

func sortReplacements(replacements []replacement) {
	for i := 1; i < len(replacements); i++ {
		for j := i; j > 0 && replacements[j].start < replacements[j-1].start; j-- {
			replacements[j], replacements[j-1] = replacements[j-1], replacements[j]
		}
	}
}

type replacement struct {
	start int
	end   int
	html  string
}

func resolveAssetURI(baseDir, raw string) string {
	if isExternalURI(raw) {
		return raw
	}
	if filepath.IsAbs(raw) {
		return fileURI(raw)
	}
	return fileURI(filepath.Join(baseDir, filepath.FromSlash(raw)))
}

func resolveLinkURI(baseDir, raw string) string {
	if isExternalURI(raw) || strings.HasPrefix(raw, "#") {
		return raw
	}
	if filepath.IsAbs(raw) {
		return fileURI(raw)
	}
	return fileURI(filepath.Join(baseDir, filepath.FromSlash(raw)))
}

func isExternalURI(raw string) bool {
	u, err := url.Parse(raw)
	return err == nil && u.Scheme != ""
}

func fileURI(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		abs = path
	}
	u := url.URL{Scheme: "file", Path: filepath.ToSlash(abs)}
	return u.String()
}
