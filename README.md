# go-teletris
Go port of my python [Teletris](https://github.com/telecoda/teletris) game

![alt text][Screenshot]
[Screenshot]: https://raw.githubusercontent.com/telecoda/go-teletris/master/original_artwork/screenshot.png "Screen shot"

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

## Building for the App Store
These are more detailed instructions for building a .apk file that can be deployed to the app store and has a pretty icon to go with it. (iOS also supported but I've not tested it).

### Build dependencies

    go get github.com/nobonobo/gomobileapp
    
    pip install icons
    
### Building the .apk

    gomobileapp build -icon icon.png -target android github.com/telecoda/go-teletris

	adb install go-teletris.apk

or

    gomobileapp install -icon icon.png -target android github.com/telecoda/go-teletris


The AndroidManifest.xml file is required to make the app full screen.  Gomobileapp then applies its own extra settings to the file during the build process.

### Troubleshooting
If you get any weird adb installation error, try installing the app from your phone first.  ADB has a habit of returning non-sensical error messages...

Also when build the app with gomobileapp, make sure you have an internet connection.  This is because the code makes a call out to the internet to a time server for signing the jar.?? Apparently.

## Acknowledgements

Font: [Karmatic Arcade](http://www.1001freefonts.com/karmatic_arcade.font) by Vic Fieger

Music: [Tetris dubstep - Mr Straightface](https://soundcloud.com/kaseythompson/tetris-dubstep-remix-free)

Gophers: me ;)

