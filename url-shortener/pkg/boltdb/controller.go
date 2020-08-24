package boltdb

import (
	"fmt"

	urlshortener "github.com/bcatubig/go-projects/url-shortener"
	"github.com/boltdb/bolt"
	"github.com/hashicorp/go-hclog"
)

const (
	bucket = "urls"
)

type Controller struct {
	db     *bolt.DB
	logger hclog.Logger
}

type NewControllerOptions struct {
	DBLocation string
	Logger     hclog.Logger
}

// NewController creates a new BoltDB Controller
func NewController(options *NewControllerOptions) (*Controller, error) {
	if options == nil {
		options = &NewControllerOptions{
			DBLocation: "url-redirects.db",
			Logger:     hclog.NewNullLogger(),
		}
	}

	if options.DBLocation == "" {
		options.DBLocation = "url-redirects.db"
	}

	if options.Logger == nil {
		options.Logger = hclog.NewNullLogger()
	}

	db, err := bolt.Open(options.DBLocation, 0600, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// initialize urls bucket in boltdb
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("boltdb init failed: %w", err)
	}

	return &Controller{
		db:     db,
		logger: options.Logger,
	}, nil

}

// Shutdown closes the boltdb database
func (c *Controller) Shutdown() error {
	c.logger.Info("closing database")
	return c.db.Close()
}

// Create creates a new short url that targets the targetURL and stores it in the DB
func (c *Controller) Create(targetURL string) (*urlshortener.Redirect, error) {
	c.logger.Info("creating shortened url", "target_url", targetURL)
	r := urlshortener.NewRedirect(targetURL)

	err := c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(r.ShortUrl), urlshortener.EncodeRedirect(r))
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to store url: %w", err)
	}

	return r, nil
}

// Lookup looks up a short url to resolve the target url
func (c *Controller) Lookup(shortURL string) (*urlshortener.Redirect, error) {
	var r *urlshortener.Redirect
	var err error

	c.logger.Info("looking up redirect", "short_url", shortURL)

	err = c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(shortURL))
		r, err = urlshortener.DecodeRedirect(v)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}
