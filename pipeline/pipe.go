package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	data, _ := ioutil.ReadFile("acetest.webm")

	wav := []string{
		"ffmpeg", "-y",
		"-loglevel", "panic",
		"-i", "pipe:0",
		"-f", "wav",
		"pipe:1",
	}
	flac := []string{
		"ffmpeg", "-y",
		"-loglevel", "panic",
		"-i", "pipe:0",
		"-f", "flac",
		"pipe:1",
	}
	mp3 := []string{
		"ffmpeg", "-y",
		"-loglevel", "panic",
		"-i", "pipe:0",
		"-f", "mp3",
		"pipe:1",
	}

	out, err := pipe(
		data,
		wav,
		flac,
		mp3,
	)

	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("acetest.mp3", out, 0644)
}

// https://github.com/mattn/go-pipeline
func pipe(buf []byte, commands ...[]string) ([]byte, error) {
	cmds := make([]*exec.Cmd, len(commands))
	var err error

	for i, c := range commands {
		cmds[i] = exec.Command(c[0], c[1:]...)

		if i == 0 && len(buf) > 0 {
			cmds[i].Stdin = bytes.NewBuffer(buf)
		}
		if i > 0 {
			if cmds[i].Stdin, err = cmds[i-1].StdoutPipe(); err != nil {
				return nil, err
			}
		}
		cmds[i].Stderr = os.Stderr
	}

	var out bytes.Buffer

	cmds[len(cmds)-1].Stdout = &out
	for _, c := range cmds {
		if err = c.Start(); err != nil {
			return nil, err
		}
	}
	for _, c := range cmds {
		if err = c.Wait(); err != nil {
			return nil, err
		}
	}

	return out.Bytes(), nil
}
