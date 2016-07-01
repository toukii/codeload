# 	codeload

github codeload url(sample): https://codeload.github.com/everfore/codeload/zip/master

#	get

	go get github.com/everfore/codeload

	go install

	cd cdwn && go install

	cd pull && go install


#	usage


_with default by @filepath.Base(dir)_

_if branch is nil, use master as default branch._

_if user is nil, use filepath.Base as default user._


## No .git repo

```
codeload #or cdln

[user/]repo[:branch]  > $
```

`https://codeload.github.com/everfore/codeload/zip/master`


>No .git repo && go install


```
cdln -i

[user/]repo[:branch]  > $

-i : go install
```

`https://codeload.github.com/everfore/codeload/zip/master` and `go install`


## With .git repo


 - git clone 

>git clone || go install


```
pull -i -r

[user/]repo[:branch]  > $

```

`git clone git:// or git@` and `go install`


 - git pull


```
pull


```

`git pull`


##	license

Apache License Version 2.0
