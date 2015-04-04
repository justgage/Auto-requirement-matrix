main: html markdown csv
html: it
	./reqirement-matrix new-req.yaml html > reqirements.html 
markdown: it
	./reqirement-matrix new-req.yaml markdown > reqirements.md 
csv: it
	./reqirement-matrix new-req.yaml csv > reqirements.csv 
it:
	go build
view:
	chromium reqirements.html
watch:
	watch make

