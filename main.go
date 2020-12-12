package main

import (
    "fmt"
    fast "git.9885.net/daima.mobi/lib/fasthttp-client"
    "github.com/atotto/clipboard"
    "github.com/go-vgo/robotgo"
    hook "github.com/robotn/gohook"
    "log"
)

var (
    rainbowPee = "https://chp.shadiao.app/api.php"
    dirtyWords = "https://zuanbot.com/api.php"

    client = fast.NewClient().AddHeader("referer", "https://zuanbot.com/")
    err error
    resp *fast.Response

    all = true
    startRainbowPee = true
)

func main() {
    add()
}

func add() {

    robotgo.EventHook(hook.KeyDown, []string{"shift", "b"}, func(e hook.Event) {
        fmt.Println("start success")
        all = true
    })

    fmt.Println("--- Please press shift + e to stop hook ---")
    robotgo.EventHook(hook.KeyDown, []string{"shift", "e"}, func(e hook.Event) {
        fmt.Println("stop success")
        all = false
    })

    robotgo.EventHook(hook.KeyDown, []string{"shift", "d"}, func(e hook.Event) {
        fmt.Println("start dirty word success")
        startRainbowPee = false
    })

    robotgo.EventHook(hook.KeyDown, []string{"shift", "r"}, func(e hook.Event) {
        fmt.Println("start rainbowPee success")
        startRainbowPee = true
    })


    robotgo.EventHook(hook.KeyDown, []string{"ctrl", "v"}, func(e hook.Event) {
        rainbowWords()
    })

    robotgo.EventHook(hook.KeyDown, []string{"command", "v"}, func(e hook.Event) {
        rainbowWords()
    })

    s := robotgo.EventStart()
    <-robotgo.EventProcess(s)
}

func rainbowWords() {
    if !all {
        return
    }
    url := rainbowPee
    if !startRainbowPee {
        url = dirtyWords
    }
    resp, err = client.Get(url)
    if err != nil {
        log.Println("request err", err)
    }
    err = clipboard.WriteAll(string(resp.Body))
    if err != nil {
        log.Println("write err", err)
    }
}