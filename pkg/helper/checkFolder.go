package helper

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Talos-hub/SimpleParser/pkg/adapters"
)

func CheckFolder(pathTofile string, logger adapters.Logging) error {
	folder := path.Dir(pathTofile)

	// get info
	info, err := os.Stat(folder)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			logger.Warn("warning, the fodler %s, is not exist, err: %w", folder, err)
			logger.Info("creating new folder: %s", folder)
			// creat new folder
			err := os.Mkdir(folder, 0755)
			if err != nil {
				logger.Error("error creating a folder", "error", err)
				return fmt.Errorf("error creating a folder: %s, err: %w", folder, err)
			}
		case errors.Is(err, os.ErrPermission):
			logger.Warn("warning, permission denied", "error", err)
			return fmt.Errorf("warning, permission denied, folder: %s, err: %w", folder, err)
		}
		logger.Error("error: %w", err)
		return err
	}

	// it check that's actually a folder
	if !info.IsDir() {
		logger.Error("error path is exist but isn't folder", "error", err)
		return fmt.Errorf("error, path exists: %s, but it is not a folder, err: %w", folder, err)
	}

	return nil
}
