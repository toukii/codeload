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


### cdwn

>cd cdwn && go install

>cdwn -i

[user/]repo[:branch] 同上
```
-i : go install
```

### cdln

>cd cdln && go install

>cdln -i -w

[user/]repo[:branch] 同上
```
-i : go install
-w: git@github.com:user/repo.git
```

##	license

Apache License Version 2.0
