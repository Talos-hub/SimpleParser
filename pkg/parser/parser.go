package parser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Talos-hub/SimpleParser/pkg/adapters"
	"github.com/Talos-hub/SimpleParser/pkg/helper"
)

type parser struct {
	// logger
	logger adapters.Logging
	// path where is parsed data
	pathStorage string
	// pattern for search
	pattern string
	mu      *sync.Mutex
	client  *http.Client
}

func NewParser(logger adapters.Logging, pathStorage, pattern string) *parser {
	return &parser{
		logger:      logger,
		pathStorage: pathStorage,
		pattern:     pattern,
		mu:          &sync.Mutex{},
		client: &http.Client{
			Timeout: 30 * time.Second, // ← Add timeout
		},
	}
}

func (p *parser) StartData(url string) error {
	err := helper.CheckFolder(p.pathStorage, p.logger)
	if err != nil {
		p.logger.Error("Error starting", "error", err)
		return err
	}

	r, err := p.client.Get(url)
	if err != nil {
		p.logger.Error("Error getting data", "error", err)
		return err
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if !errors.Is(err, io.EOF) {
		p.logger.Error("Error reading from response", "error", err)
	}

	re, err := regexp.Compile(p.pattern)
	if err != nil {
		p.logger.Error("Error create regexp", "error", err)
		return err
	}

	match := re.FindStringSubmatch(string(data))
	if len(match) < 2 {
		p.logger.Warn("exchange rate not found on the page", "page", url)
		return nil
	}

	p.mu.Lock()
	file, err := os.OpenFile(p.pathStorage, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		p.logger.Error("Error open or creating file: ", "error", err)
		return fmt.Errorf("error open or creating file: %w", err)
	}
	defer file.Close()
	text := strings.Join(match, " ")
	_, err = file.WriteString("ru: " + text)
	if err != nil {
		p.logger.Error("Error writing to file", "error", err)
	}
	p.mu.Unlock()

	return nil
}

func (p *parser) ParseMultipleUrl(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			p.StartData(u)
		}(url)
	}
	wg.Wait()
}
