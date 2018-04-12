
```elixir
$ cat acetest.webm |
    ffmpeg -y -loglevel panic -i pipe:0 -f wav  pipe:1 |
    ffmpeg -y -loglevel panic -i pipe:0 -f flac pipe:1 |
    ffmpeg -y -loglevel panic -i pipe:0 -f mp3  pipe:1 |
    cat - > acetest.mp3
```


```go
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
```
