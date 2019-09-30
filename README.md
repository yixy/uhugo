# uhugo #

uhugo is a CLI tool to update markdown content file for hugo in yaml front matter.

## Install ##

```
go install -ldflags="-X main.ver=0.0.1 -X 'main.env=`uname -mv`' -X 'main.buildTime=`date`'" github.com/yixy/uhugo
```

## How to use ##

Generate the markdown file list.

```
# we have markdown files as follows
$ cat > a.md <<EOF                  
this is a content.
EOF

$ cat > b.md <<EOF                  
this is b content.
EOF

$ cat > c.md <<EOF                  
this is c content.
EOF

# generate the markdown file list
$ uhugo list
.list file generate success.

# filename is .list
$ cat .list     
a|0f29ad984b2ad8973adb8ff6a429a20f
b|e894ae6683e091a244777d0cc04f0998
c|4f339eac345d3cd783d3c40e7ef469d2
```

You can manage the list to change the files' name.

```
# arranged .list file.
$ vim .list
aaa|0f29ad984b2ad8973adb8ff6a429a20f
bbb|e894ae6683e091a244777d0cc04f0998
ccc|4f339eac345d3cd783d3c40e7ef469d2
```

make the markdown files update by .list. We can also specify the tags, categories of file.

```
$ uhugo update -c=test -t=demo,abc 
file is updated.

$ ls
aaa.md bbb.md ccc.md

$ cat aaa.md
---
categories:
- test
date: "2019-09-30T09:08:25+08:00"
lastmod: "2019-09-30T09:08:25+08:00"
tags:
- demo
- abc
title: aaa
---
this is a content.

```

more usage see `uhugo -h`