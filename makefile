main: html markdown csv
html: it
	./reqirement-matrix req.yaml html > index.html 
markdown: it
	./reqirement-matrix req.yaml markdown > reqirements.md 
csv: it
	./reqirement-matrix req.yaml csv > reqirements.csv 
it:
	go build
view:
	chromium index.html
watch:
	watch make
image:
	rm images/*.png
	pdftoppm -png index.pdf images/matrix
	convert -trim images/*.png 
