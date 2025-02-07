// cmd/api/main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: Initialiser la configuration
	// TODO: Initialiser la base de données
	// TODO: Initialiser le router

	port := "8080" // À déplacer dans la configuration

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// Handler sera initialisé plus tard
	}

	// Canal pour gérer l'arrêt gracieux
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
		}
	}()

	log.Printf("Serveur démarré sur le port %s", port)

	<-done
	log.Print("Arrêt du serveur...")

	// TODO: Implémentation de l'arrêt gracieux
}
