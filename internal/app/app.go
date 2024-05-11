package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/inview-team/veles.assistant/internal/config"
	httpapi "github.com/inview-team/veles.assistant/internal/controller/http"
	"github.com/inview-team/veles.assistant/internal/controller/ws"
	"github.com/inview-team/veles.assistant/internal/hub"
	"github.com/inview-team/veles.assistant/internal/service"
	log "github.com/sirupsen/logrus"
)

type App struct {
	config         *config.Config
	wsSrv          *http.Server
	httpSrv        *http.Server
	sessionService service.SessionService
	matchService   service.MatchService
	executeService service.ExecuteService
	hub            hub.Hub
}

func NewApp(cfg *config.Config, ss service.SessionService, ms service.MatchService, es service.ExecuteService, hub hub.Hub) *App {
	app := &App{
		config:         cfg,
		sessionService: ss,
		matchService:   ms,
		executeService: es,
	}
	return app
}

func (a *App) Start() {
	go a.startWS()
	go a.startHTTP()
	a.awaitSignals()
}

func (a *App) Stop() error {
	var err error
	if a.httpSrv != nil {
		log.Info("Shutting down http server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if shutdownErr := a.httpSrv.Shutdown(ctx); shutdownErr != nil {
			log.Errorf("Failed to shutdown http server: %v", shutdownErr)
			err = shutdownErr
		}

		log.Info("Http server stopped")
	}

	if a.wsSrv != nil {
		log.Info("Shutting down WebSocket server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if shutdownErr := a.wsSrv.Shutdown(ctx); shutdownErr != nil {
			log.Errorf("Failed to shutdown WebSocket server: %v", shutdownErr)
			err = shutdownErr
		}

		log.Info("WebSocket server stopped")
	}

	return err
}

func (a *App) awaitSignals() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	sig := <-interrupt
	log.Infof("Received signal: %v", sig)
	err := a.Stop()
	if err != nil {
		log.Errorf("Error during server shutdown: %v", err)
	}
}

func (a *App) startHTTP() {
	log.Info("Starting HTTP server")

	router := mux.NewRouter()
	httpHandler := httpapi.NewHttpHandler(a.sessionService, a.matchService, a.executeService)
	router.HandleFunc("/api/v1/sessions", httpHandler.StartSession).Methods("POST")
	router.HandleFunc("/api/v1/actions", httpHandler.HandleAction).Methods("POST")

	a.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.config.HTTPHost, a.config.HTTPPort),
		Handler: router,
	}

	if err := a.httpSrv.ListenAndServe(); err != nil {
		log.Errorf("HTTP Server failed: %v", err)
	}
}

func (a *App) startWS() {
	log.Infof("Starting WebSocket server")

	address := fmt.Sprintf("%s:%d", a.config.WebSocketHost, a.config.WebSocketPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("Failed to create TCP listener: %v", err)
		return
	}

	router := mux.NewRouter()
	wsHandler := ws.NewWsHandler(a.sessionService, a.matchService, a.executeService, a.hub)
	router.HandleFunc("/ws", wsHandler.HandleWs)

	a.wsSrv = &http.Server{
		Handler: router,
	}

	log.Infof("WebSocket server started on %s", address)
	if err := a.wsSrv.Serve(lis); err != nil {
		if err == http.ErrServerClosed {
			log.Infof("WebSocket server stopped")
			return
		}
		log.Errorf("Closed WebSocket connection: %v", err)
		return
	}
}
