package boot

import (
	"log"
	"net/http"

	"product/internal/config"
	"product/pkg/httpclient"

	"github.com/jmoiron/sqlx"

	doData "product/internal/data/do"
	mpData "product/internal/data/mp"
	outletData "product/internal/data/outlet"
	productData "product/internal/data/product"
	productServer "product/internal/delivery/http"
	productHandler "product/internal/delivery/http/product"
	productService "product/internal/service/product"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		srv      productServer.Server // HTTP Server Object
		mpD      mpData.Data
		outletD  outletData.Data
		doD      doData.Data
		productD productData.Data        // Domain data layer
		productS productService.Service  // Domain service layer
		productH *productHandler.Handler // Domain handler
		cfg      *config.Config          // Configuration object
	)

	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()
	// HTTP Client
	httpc := httpclient.NewClient()
	//API
	mpD = mpData.New(httpc, cfg.API.MP)
	outletD = outletData.New(httpc, cfg.API.Outlet)
	doD = doData.New(httpc, cfg.API.DO)
	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Diganti dengan domain yang anda buat
	productD = productData.New(db)
	productS = productService.New(productD, mpD, outletD, doD)
	productH = productHandler.New(productS)

	srv = productServer.Server{
		Product: productH,
	}

	if err := srv.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
