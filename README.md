# 	codeload

github codeload url(sample): https://codeload.github.com/everfore/codeload/zip/master

##	get

go get github.com/everfore/codeload

##	usage

>codeload

>user/repo:master

	or

>codeload

>user/repo

	or

>codeload

>repo    with default by @filepath.Base(dir)

if branch is nil, use master as default branch.
if user is nil, use filepath.Base as default user.

##	license

Apache License Version 2.0
