main: html markdown csv
html: it
	./reqirement-matrix req.yaml html > index.html 
markdown: it
	./reqirement-matrix req.yaml markdown > reqirements.md 
csv: it
	./reqirement-matrix req.yaml csv > reqirements.csv 
csv2: it
	./reqirement-matrix req-together.yaml csv > reqirements.csv 
it:
	go build
view:
	chromium index.html
watch:
	watch make
image:
	rm images/*.png
	pdftoppm -png index.pdf images/matrix
trim:
	convert -trim images/*.png 
	rm images/matrix-0*
	rm images/matrix-10.png
	zip images.zip images/*.png
