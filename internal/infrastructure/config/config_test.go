package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadUsesDefaultsWhenConfigFileIsMissing(t *testing.T) {
	tmpDir := t.TempDir()
	changeWorkingDirectory(t, tmpDir)
	clearConfigEnv(t)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.App.Environment != "development" {
		t.Fatalf("expected default environment development, got %q", cfg.App.Environment)
	}
	if cfg.Server.Host != "0.0.0.0" {
		t.Fatalf("expected default host 0.0.0.0, got %q", cfg.Server.Host)
	}
	if cfg.Server.Port != 8080 {
		t.Fatalf("expected default port 8080, got %d", cfg.Server.Port)
	}
}

func TestLoadUsesDotEnvFile(t *testing.T) {
	tmpDir := t.TempDir()
	changeWorkingDirectory(t, tmpDir)
	clearConfigEnv(t)

	env := []byte("APP_APP_ENVIRONMENT=test\nAPP_SERVER_HOST=127.0.0.1\nAPP_SERVER_PORT=3000\n")
	if err := os.WriteFile(filepath.Join(tmpDir, ".env"), env, 0o600); err != nil {
		t.Fatalf("write .env: %v", err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.App.Environment != "test" {
		t.Fatalf("expected .env environment test, got %q", cfg.App.Environment)
	}
	if cfg.Server.Host != "127.0.0.1" {
		t.Fatalf("expected .env host 127.0.0.1, got %q", cfg.Server.Host)
	}
	if cfg.Server.Port != 3000 {
		t.Fatalf("expected .env port 3000, got %d", cfg.Server.Port)
	}
}

func TestLoadKeepsEnvironmentVariablesOverDotEnvFile(t *testing.T) {
	tmpDir := t.TempDir()
	changeWorkingDirectory(t, tmpDir)
	clearConfigEnv(t)
	t.Setenv("APP_SERVER_PORT", "9000")

	if err := os.WriteFile(filepath.Join(tmpDir, ".env"), []byte("APP_SERVER_PORT=3000\n"), 0o600); err != nil {
		t.Fatalf("write .env: %v", err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.Server.Port != 9000 {
		t.Fatalf("expected environment port 9000, got %d", cfg.Server.Port)
	}
}

func TestLoadReturnsErrorForInvalidConfig(t *testing.T) {
	tmpDir := t.TempDir()
	changeWorkingDirectory(t, tmpDir)
	clearConfigEnv(t)

	if err := os.WriteFile(filepath.Join(tmpDir, "config.yaml"), []byte("app: [invalid"), 0o600); err != nil {
		t.Fatalf("write invalid config: %v", err)
	}

	if _, err := Load(); err == nil {
		t.Fatal("expected invalid config to return an error")
	}
}

func TestServerAddress(t *testing.T) {
	cfg := ServerConfig{Host: "127.0.0.1", Port: 3000}

	if got := cfg.Address(); got != "127.0.0.1:3000" {
		t.Fatalf("expected address 127.0.0.1:3000, got %q", got)
	}
}

func changeWorkingDirectory(t *testing.T, dir string) {
	t.Helper()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("change working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	})
}

func clearConfigEnv(t *testing.T) {
	t.Helper()

	for _, key := range []string{"APP_APP_ENVIRONMENT", "APP_SERVER_HOST", "APP_SERVER_PORT"} {
		value, ok := os.LookupEnv(key)
		if err := os.Unsetenv(key); err != nil {
			t.Fatalf("unset %s: %v", key, err)
		}
		t.Cleanup(func() {
			if ok {
				if err := os.Setenv(key, value); err != nil {
					t.Fatalf("restore %s: %v", key, err)
				}
				return
			}
			if err := os.Unsetenv(key); err != nil {
				t.Fatalf("restore unset %s: %v", key, err)
			}
		})
	}
}
