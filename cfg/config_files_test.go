package cfg

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	testFileName := "test_config"
	configDir, _ := WtfConfigDir()
	expectedFile := fmt.Sprintf("%s/%s", configDir, testFileName)

	t.Run("File does not exist yet", func(t *testing.T) {
		actual, err := CreateFile(testFileName)

		if err != nil {
			t.Errorf("Got error: %s", err)
		}

		if actual != expectedFile {
			t.Errorf("Expected %s, got %s", expectedFile, actual)
		}

		_, statErr := os.Stat(expectedFile)
		if os.IsNotExist(statErr) {
			t.Errorf("File %s should have been created, but does not exist.", expectedFile)
		}
	})

	t.Run("File already exists", func(t *testing.T) {
		os.Create(expectedFile)
		statsBefore, _ := os.Stat(expectedFile)

		actual, err := CreateFile(testFileName)

		if err != nil {
			t.Errorf("Got error: %s", err)
		}

		if actual != expectedFile {
			t.Errorf("Expected %s, got %s", expectedFile, actual)
		}

		statsAfter, _ := os.Stat(expectedFile)
		if statsAfter.ModTime() != statsBefore.ModTime() {
			t.Errorf("Existing file %s should not have been modified.", expectedFile)
		}
	})

	os.Remove(expectedFile)
}
