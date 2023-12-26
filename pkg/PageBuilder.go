package pkg

import "fmt"

type PageBuilder struct {
	Page       string
	Components string
	Title      string
}

func (b *PageBuilder) AddComponents(components []string) {
	for i := 0; i < len(components); i++ {
		b.Components = b.Components + components[i]
	}
}

func (b *PageBuilder) HtmlBytes() []byte {
	return []byte(fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset='UTF-8'>
			<meta name='viewport' content='width=device-width, initial-scale=1.0'>
			<script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script>
			<script src='https://unpkg.com/htmx.org@1.9.9'></script>
			<script src='https://unpkg.com/hyperscript.org@0.9.12'></script>
			<script src='/public/js/index.js'></script>
			<link rel='stylesheet' href='/public/css/styles.css'>
			<link rel='stylesheet' href='/public/css/animate.css'>
			<link rel=\"preconnect\" href=\"https://fonts.googleapis.com\">
			<link rel=\"preconnect' href=\"https://fonts.gstatic.com\" crossorigin>
			<link href=\"https://fonts.googleapis.com/css2?family=Chelsea+Market&display=swap\" rel=\"stylesheet\">
			<title>%s</title>
		</head>
		<body>
			%s
		</body>
		</html>
	`, b.Title, b.Components))
}