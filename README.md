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

# display by alphabet order default
$ ls
a.md    b.md    c.md 

# generate the markdown file list
$ uhugo list
.list file generate success.

# filename is .list
$ cat .list     
a
b
c
```

You can manage the list to arrange the files.

```
# arranged .list file.
$ vim .list
b
a
c
```

make the markdown files arranged by rename file in list order, and update front matter of files. We can mod the title, tags, categories of file.

```
$ uhugo update -c=test -t=demo,abc 
file is updated.

$ ls
1-b.md 2-a.md 3-c.md

$ cat 1-b.md 
---
categories:
- test
date: "2019-09-25T14:14:52+08:00"
lastmod: "2019-09-25T14:22:26+08:00"
tags:
- demo
- abc
title: 1-b
---
this is b content.

```

more usage see `uhugo -h`