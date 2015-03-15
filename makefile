main: html markdown
html: it
	./reqirement-matrix document.yaml html > reqirements.html 
markdown: it
	./reqirement-matrix document.yaml markdown > reqirements.md 
it:
	go build
view:
	chromium reqirements.html
watch:
	watch make

