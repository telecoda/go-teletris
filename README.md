# go-teletris
Go port of my python Teletris game

![alt text][Screenshot]
[Screenshot]: https://raw.githubusercontent.com/telecoda/go-teletris/master/assets/screenshot.png "Screen shot"

## Dependencies
[go-mobile](https://github.com/golang/go/wiki/Mobile)

    $ go get golang.org/x/mobile/cmd/gomobile
    $ gomobile init # it might take a few minutes

## building the code (desktop)

    go get -u -v
    go build -tags=release && ./go-teletris

   
## building the code (mobile)
This command will create an .apk file

    gomobile build


## installing the code (mobile)
This command will build and install an .apk file

    gomobile install -tags=release

Make sure you connect you phone via USB and enable developer tools USB debugging


## Disclaimer
Please note this is very inefficient unoptimised code, so only runs at about 6-7 frames per second.  Currently code is rendering 300+ sprites every frame.  Next step is to optimise this by rendering to an offscreen buffer first.
