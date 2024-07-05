package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"oft/config/colors"
	"oft/config/cli"
	"os"
	// "os/exec"
	"strings"
	"unicode"
)

var (
	// Variables globales pour l'application
	hostMode = false
	folderPath = ""
	port string

	// Variables pour l'interface CLI
	cliPrefix = "user"
	cliType   = "@"
	cliSuffix = ""
	cliEnd    = "~>"
	cliSTR    string
)

func main() {
	cliSuffix = GetOutboundIP().String()
	reader := bufio.NewReader(os.Stdin)
	cli.AsciiStart()

	for {
		if cliPrefix == "user" {
			cliSTR = color.Green + cliPrefix + cliType + color.Orange + cliSuffix + color.Green + cliEnd + " " + color.Reset
		} else if(cliPrefix == "host") {
			cliSTR = color.Red + cliPrefix + color.Green + cliType + color.Orange + cliSuffix + color.Green + cliEnd + " " + color.Reset
		} else{
			cliSTR = color.Red+"[ERROR cliSTR]"+color.Reset
		}

		fmt.Print(cliSTR)

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée:", err)
			continue
		}
		
		command = strings.TrimSpace(command)
		commands := splitCommand(command)
		Command(commands)
		fmt.Println("")
	}
}

func Command(commands []string) {
	if len(commands) == 0 {
		fmt.Println("Aucune commande fournie.")
		return
	}

	if commands[0] == "oft" {
		// Si la commande est "oft", on vérifie la sous-commande
		if len(commands) < 2 {
			fmt.Println("Veuillez spécifier une sous-commande pour 'oft'.")
			return
		}

		subCommand := commands[1]
		switch subCommand {
		case "help":
			cli.Help()
		case "kill":
			fmt.Println("Arrêt du programme...")
			// Ajoutez ici le code pour arrêter proprement votre programme si nécessaire
			os.Exit(0)
		case "clear":
			cli.ClearScreen()
		case "status":
			cli.Status(cliPrefix, cliType, cliSuffix, cliEnd, folderPath, port)
		case "host":
			cli.Host(&hostMode, commands, &cliPrefix, &folderPath, &port)
		case "login":
			// cli.Login(commands[2])
		default:
			fmt.Printf("Commande 'oft %s' non reconnue.\n", subCommand)
		}
	} else {
		fmt.Println("Veuillez écrire d'abord \"oft\"")
	}
}

// Fonction pour obtenir l'adresse IP sortante préférée de cette machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func splitCommand(command string) []string {
	var parts []string
	var current strings.Builder
	inQuotes := false

	for _, char := range command {
		switch {
		case char == '"' && !inQuotes:
			inQuotes = true
		case char == '"' && inQuotes:
			inQuotes = false
		case unicode.IsSpace(char) && !inQuotes:
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}
