cldir
=====

I created this tool to make it easier for me to find the last file I downloaded
in my messy downloads folder until I clean it up, but it can work with any folder.


It will keep the last 5 modified files in the downloads
folder by default, and move the rest to `backed` folder inside the original
directory, the directory that will be cleaned and the number of files to keep
can be customized using flags `dir` and `remain`

#### Installation
to install it run
`go get github.com/shadi/cldir`

#### Usage
run `cldir`, it requires the environment variable $GOPATH/bin to be set and part of $PATH environment variable

to use it with another folder and keep the last 3 files only
`cldir -dir $SOME_DIR -remain 3`
