package cli

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"
    "io/ioutil"
    "context"
)

var (
    server *http.Server
    mux    = http.NewServeMux() // ServeMux global
)

func Host(hostMode *bool, commands []string, cliPrefix *string, folderPath *string, port *string) {
    if len(commands) < 3 {
        fmt.Println("Invalid arguments")
        return
    }

    if commands[2] == "start" && commands[3] == "-d" && len(commands) >= 5 && !*hostMode {
        fmt.Println("Starting hosting... on " + commands[4])
        *cliPrefix = "host"
        *hostMode = true
        *folderPath = commands[4]

        if len(commands) > 6 && commands[5] == "-p" && commands[6] != "" {
            *port = commands[6]
        } else {
            *port = "7000"
        }

        startServer(commands[4], *port)

    } else if commands[2] == "stop" && *hostMode {
        fmt.Println("Stopping hosting...")
        *cliPrefix = "user"
        *hostMode = false
        *folderPath = ""

        stopServer()

    } else if commands[2] == "scan" && len(commands) >= 4 {
        ipPort := commands[3]
        scanFolder(ipPort)

    } else {
        fmt.Println("Args detection error")
        return
    }
}

func startServer(folderPath, port string) {
    // Resolve absolute path
    absPath, err := filepath.Abs(folderPath)
    if err != nil {
        log.Fatal("Error resolving folder path:", err)
    }

    // Create file server
    fileServer := http.FileServer(http.Dir(absPath))

    // Register handler with global ServeMux
    mux.Handle("/", fileServer)

    // Start HTTP server
    server = &http.Server{Addr: ":" + port, Handler: mux} // Use global mux
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal("Error starting HTTP server:", err)
        }
    }()
    fmt.Printf("Server started. Hosting %s on port %s.\n", absPath, port)
}


func stopServer() {
    if server != nil {
        // Arrêter le serveur HTTP
        if err := server.Shutdown(context.Background()); err != nil {
            log.Fatal("Error stopping HTTP server:", err)
        }
        fmt.Println("Server stopped.")
        
        // Réinitialiser le ServeMux
        resetServeMux()

        // Définir server à nil pour indiquer qu'il n'est plus en cours d'exécution
        server = nil

    } else {
        fmt.Println("Server is not running.")
    }
}

func resetServeMux() {
    // Créer un nouveau ServeMux pour remplacer le global
    mux = http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Gestionnaire par défaut pour le chemin racine
        http.Error(w, "Not found", http.StatusNotFound)
    }))
    fmt.Println("ServeMux reset.")
}

func scanFolder(ipPort string) {
    // Send HTTP request to server at ipPort to get directory listing
    resp, err := http.Get("http://" + ipPort + "/")
    if err != nil {
        log.Fatal("Error scanning folder:", err)
    }
    defer resp.Body.Close()

    // Read response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading response body:", err)
    }

    // Print directory listing
    fmt.Println("Directory listing from " + ipPort + ":\n" + string(body))
}
