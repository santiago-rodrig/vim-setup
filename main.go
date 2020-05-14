package main

import (
    "os"
    "os/exec"
    "log"
    "io/ioutil"
    "strings"
    "time"
    "path/filepath"
)

func main() {
    // sets the user home directory
    homeDir := os.Getenv("HOME")
    executable, err := os.Executable()
    if err != nil {
        log.Fatal(err)
    }
    executableDir := filepath.Dir(executable)
    // gets the contents of the vim configuration file
    vimConfigFile, err := ioutil.ReadFile(
        strings.Join([]string{executableDir,"assets/vimrc"},"/"),
    )
    if err != nil {
        log.Fatal(err)
        return
    }
    // writes the contents to the vim configuration file
    err = ioutil.WriteFile(
        strings.Join([]string{homeDir,".vimrc"},"/"),
        vimConfigFile,
        os.FileMode(0777),
    )
    if err != nil {
        log.Fatal(err)
        return
    }
    time.Sleep(time.Second)
    log.Println("$HOME/.vimrc has been set up")
    // look for the existence of vim on the system
    _, err = exec.LookPath("vim")
    if err != nil {
        log.Println("Please install vim: https://www.vim.org/download.php")
        log.Fatal(err)
        return
    }
    // look for the existence of curl on the system
    _, err = exec.LookPath("curl")
    if err != nil {
        log.Println("Please install curl: https://curl.haxx.se/download.html")
        log.Fatal(err)
        return
    }
    // look for the existence of git on the system
    _, err = exec.LookPath("git")
    if err != nil {
        log.Println("Please install git: https://git-scm.com/downloads")
        log.Fatal(err)
        return
    }
    // set the command for installing vim-plug
    vimPlugURL := strings.Join([]string{
        "https://raw.githubusercontent.com/junegunn",
        "vim-plug/master/plug.vim",
    },"/")
    cmd := exec.Command(
        "curl",
        "-fLo",
        homeDir + "/.vim/autoload/plug.vim",
        "--create-dirs",
        vimPlugURL,
    )
    err = cmd.Run()
    if err != nil {
        log.Fatal("Fetching vim plug did not work")
        return
    }
    // install vim plugins
    cmd = exec.Command(
        "vim",
        "+PlugInstall",
        "+qa",
    )
    err = cmd.Run()
    if err != nil {
        log.Fatal("Installing plugins with vim plug did not work")
    }
    time.Sleep(time.Second)
    log.Println("Vim plugins successfully installed")
}
