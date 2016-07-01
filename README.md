# 	codeload

github codeload url(sample): https://codeload.github.com/everfore/codeload/zip/master

##	get

go get github.com/everfore/codeload

##	usage

### No .git repo

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


>cd No .git repo || go install

>cdwn -i

[user/]repo[:branch]  > $ 同上
```
-i : go install
```

### .git repo

#### git clone 

>git clone || go install


```
pull -i -r

[user/]repo[:branch]  > $

```

```
-i : go install
-r: git://github.com/user/repo.git , default: git@github.com:user/repo.git
```

#### git pull

>git pull


```
pull



```

##	license

Apache License Version 2.0
