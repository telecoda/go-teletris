# go-teletris
Go port of my python Teletris game

![alt text][Screenshot]
[Screenshot]: https://raw.githubusercontent.com/telecoda/go-teletris/master/orginal_arkwork/screenshot.png "Screen shot"

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


## References

Font: [Karmatic Arcade](http://www.1001freefonts.com/karmatic_arcade.font) by Vic Fieger

