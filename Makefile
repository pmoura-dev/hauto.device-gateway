deploy:
	git tag -d $(tag)
	git push -d origin $(tag)
	git tag $(tag)
	git push origin $(tag)
