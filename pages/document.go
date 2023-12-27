package pages

import "net/http"

type Document struct{}

func (p Document) Name() string {
	return "Document"
}

func (p Document) URL() string {
	return ""
}

func (p Document) Template() string {
	return `
	{{define "DocumentStart"}}
	<!DOCTYPE html>
	<html> 
		<head>
			<meta charset='utf-8'>
			<meta http-equiv='X-UA-Compatible' content='IE=edge'> 
			<meta name='viewport' content='width=device-width, initial-scale=1'>  
			<link rel="icon" href="/static/favicon.svg" type="image/svg+xml">

			{{template "Head"}} 
		 
			<script src="/static/htmx.min.js"></script> 
			<script src="/static/tailwind.min.js"></script> 
			<script>
				tailwind.config = {
					theme: {
						extend: {
							colors: {
								clifford: '#da373d',
							}
						}
					}
				}
			</script>
		</head> 
		<body>  
		 
	{{end}}
	{{define "DocumentEnd"}}   
			
		</body> 
	</html>
	{{end}}
	`
}

func (p Document) Props(r *http.Request) any {
	return nil
}
