# go-converter

This tool is written in golang and it simply converts video files using ffmpeg in the same
directory to '.ogg' files encoded with the opus codec.

In addition this program uses all available CPUs on the system.

## Installation

Simply call go build

    go build
    mv converter ~/bin

Add your ~/bin directory to your $PATH variable.

## Using it

Then go into your directory where your video files are listed:

    cd ~/myVideoDir

Then call the converter

    converter           # converts all *.mp4 and *.mkv, *.webm to *.ogg

It uses all your available CPUs on the system and calls for every logical CPU ffmpeg.

## Dependencies

This program calls *ffmpeg*. So you have to install it.
