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
	userEnv := os.Getenv("USER")

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

	groupCommand := exec.Command("/bin/sh", "-c", "groups  | awk '{print $1}'")

	groupEnv, err := groupCommand.Output()

	if err != nil {
		fmt.Println(err)
	}


	command := fmt.Sprintf("sudo mount.cifs //%s/%s %s -o user=%s,uid=%s,gid=%s", server, folderShare, path, user, userEnv, string(groupEnv))

	cmd := exec.Command("/bin/sh", "-c", command)
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}
}
