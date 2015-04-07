main: html markdown csv
html: it
	./reqirement-matrix new-req.yaml html > index.html 
markdown: it
	./reqirement-matrix new-req.yaml markdown > reqirements.md 
csv: it
	./reqirement-matrix new-req.yaml csv > reqirements.csv 
it:
	go build
view:
	chromium index.html
watch:
	watch make

images:
	pdftoppm -png index.pdf matrix
