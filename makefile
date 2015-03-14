main: it
	./reqirement-matrix test.yaml > reqirements.html 
it:
	go build
view:
	chromium reqirements.html

