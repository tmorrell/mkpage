
# Installation

*mkpage* is a command line program run from a shell like Bash. You can find compiled
version in the [releases](https://github.com/caltechlibrary/mkpage/releases/latest). 
Download the zip file and unzip it. The filename is in the form of `mkpage-VERSION_NO-release.zip`.
Inside the zip file look for the directory that matches your computer and copy that someplace
defined in your path (e.g. $HOME/bin if you're running things on Unix/Linux/Mac OS X). 

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (ARM7)

## Mac OS X

1. Go to [github.com/caltechlibrary/mkpage/releases/latest](https://github.com/caltechlibrary/mkpage/releases/latest)
2. Click on the green "mkpage-VERSION_NO-release.zip" link and download
3. Open a terminal and type `cd ~/Downloads/` then unzip the file `unzip mkpage-VERSION_NO-release.zip` and `cd dist/macosx-amd64/`
4. Copy the *mkpage* to a "bin" directory in your path.  For example, type `sudo cp mkpage /usr/local/bin`
5. Test by typing `mkpage -h`

## Windows

1. Go to [github.com/caltechlibrary/mkpage/releases/latest](https://github.com/caltechlibrary/mkpage/releases/latest)
2. Click on the green "mkpage-release.zip" link and download
3. Open the file manager find the downloaded file and unzip it (e.g. mkpage-release.zip)
4. Look in the unziped folder and find dist/windows-amd64/mkpage.exe
5. Drag (or copy) the *mkpage.exe* to a "bin" directory in your path (a good option is C\Users\username\bin)
6. Open Bash and and test by typing `mkpage -h`
7. If it doesn't work type `echo $PATH` and copy *mkpage.exe* to one of the directories listed

## Linux

1. Go to [github.com/caltechlibrary/mkpage/releases/latest](https://github.com/caltechlibrary/mkpage/releases/latest)
2. Click on the green "mkpage-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/mkpage-release.zip)
4. In the unziped directory and find for dist/linux-amd64/mkpage
5. copy the *mkpage* to a "bin" directory (e.g. cp ~/Downloads/mkpage-release/dist/linux-amd64/mkpage ~/bin/)
6. From the shell prompt run `mkpage -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/caltechlibrary/mkpage/releases/latest](https://github.com/caltechlibrary/mkpage/releases/latest)
2. Click on the green "mkpage-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/mkpage-release.zip)
4. In the unziped directory and find for dist/raspberrypi-arm7/mkpage
5. copy the *mkpage* to a "bin" directory (e.g. cp ~/Downloads/mkpage-release/dist/raspberrypi-arm7/mkpage ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
6. From the shell prompt run `mkpage -h`

