package vba_control

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Client struct {
	rom, windowID string
	controls      chan string
}

func New(rom string) *Client {
	return &Client{
		rom:      rom,
		controls: make(chan string),
	}
}

func (v *Client) SendInput(key string) {
	v.controls <- strings.TrimSpace(key)
}

func (v *Client) waitForInput() {
	for {
		v.pressButton(<-v.controls)
	}
}

func (v *Client) Start() {
	go cmdWithDisplay("vbam " + v.rom)
	v.findWindowID()

	go v.waitForInput()
}

// The window ID is needed for xdotool to send keyboard input to the GBA emulator's window
func (v *Client) findWindowID() {
	time.Sleep(2 * time.Second) // TODO: do this in a loop to make sure the window is there already
	v.windowID = strings.TrimSpace(cmdWithDisplay("xdotool search vba-m | head -1"))
}

func (v *Client) pressButton(key string) {
	keys := map[string]string{
		"a":     "z",
		"b":     "x",
		"up":    "Up",
		"right": "Right",
		"down":  "Down",
		"left":  "Left",
		"start": "Return",
	}
	val, exists := keys[key]

	if exists {
		v.sendKeyToWindow(val)
	}
}

func (v *Client) sendKeyToWindow(key string) {
	cmdWithDisplay("xdotool windowactivate " + v.windowID)
	time.Sleep(20 * time.Millisecond)
	cmdWithDisplay("xdotool keydown " + key)
	time.Sleep(20 * time.Millisecond)
	cmdWithDisplay("xdotool keyup " + key)
}

func cmdWithDisplay(cmd string) string {
	output, err := exec.Command("sh", "-c", "DISPLAY=:0 "+cmd).Output()

	if err != nil {
		fmt.Println(err)
	}

	return string(output)
}
