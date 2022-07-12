package main

import (
	"fmt"
	"os"
	"path/filepath"
	"os/exec"
)

func main() {
	var (
		path         string
		user         string
		folderShare  string
		server       string
	)

	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Ingrese ip del servidor: ")
	fmt.Scanln(&server)

	fmt.Printf("Ingrese la ruta a montar, Actualmente te encuentras en: %s", pwd)
	fmt.Println()
	fmt.Scanln(&path)

	fmt.Println("Ingrese el usuario: ")
	fmt.Scanln(&user)

	fmt.Println("Ingrese carpeta a mapear: ")
	fmt.Scanln(&folderShare)

	if path == "" {
		path = pwd
	} else {
		path = filepath.Join(pwd, path)
	}

	command := fmt.Sprintf("sudo mount.cifs //%s/%s %s -o user=%s", server, folderShare, path, user)

	cmd := exec.Command("/bin/sh", "-c", command)
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}
}
