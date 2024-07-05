package main

import (
	"fmt"
	"net"
	"log"
)

var (
	// Variables globales pour l'application
	command string
	hostMode = false
	
	// Variables pour l'interface CLI
	cliPrefix = "user"
	cliType = "@"
	cliSuffix = ""
	cliEnd = "~>"
	cliSTR string
)

func main() {
	cliSuffix = GetOutboundIP().String()
	
	for {
		cliSTR = cliPrefix + cliType + cliSuffix + cliEnd
		fmt.Print(cliSTR)
		fmt.Scanln(&command)
		command := os.Args
		Command()
	}
}

func Command() {
	if command == "oft" {
		// Si la commande est "oft", on vérifie la sous-commande
		switch command {
		case "help":
			fmt.Println("Affichage de l'aide...")
		case "kill":
			fmt.Println("Arrêt du programme...")
			// Ajoutez ici le code pour arrêter proprement votre programme si nécessaire
		default:
			fmt.Println("Commande 'oft : ", command, "' -> non reconnue.")
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
