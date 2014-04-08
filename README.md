VBAControl
=======

This is a Go package that lets you remote control a GameBoy Advance Emulator (VisualBoyAdvance).

## Examples
This repository includes two examples for using this package to control an instance of VBA through a web interface or from STDIN.
To try them out, clone this repository into your $GOPATH/src and then execute `go run web_control.go /path/to/rom_file.gba`. Make sure you have all the dependencies installed.

## Dependencies
- Go
- xdotool
- VBA-M
- X11