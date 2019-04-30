
# Deployment
$ bee pack -be GOOS=linux -be GOARCH=amd64
$ bee pack -be GOOS=windows
$ bee pack -be GOOS=windows -be GOARCH=386

# Push to github.com

```
  $ git init
  $ git add -A
  $ git  commit -m "first commit"
  $ git remote add origin https://github.com/ujnzxw/go-wallpaper.git
  $ git push origin master
```
