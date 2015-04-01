main: html markdown csv
html: it
	./reqirement-matrix document.yaml html > reqirements.html 
markdown: it
	./reqirement-matrix document.yaml markdown > reqirements.md 
csv: it
	./reqirement-matrix document.yaml csv > reqirements.csv 
it:
	go build
view:
	chromium reqirements.html
watch:
	watch make

