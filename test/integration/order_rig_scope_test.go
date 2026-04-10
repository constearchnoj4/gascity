//go:build integration

package integration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOrderRunRigScopedFormulaUsesRigStore(t *testing.T) {
	cityDir, rigDir := setupOrderRunRigCity(t)
	writeRigFormulaOrder(t, rigDir, "rig-formula", "rig-formula", "polecat")

	out, err := gcDolt(cityDir, "order", "run", "--rig", "hello-world", "rig-formula")
	if err != nil {
		t.Fatalf("gc order run rig-formula failed: %v\n%s", err, out)
	}

	wispID := extractOrderWispID(t, out)
	if !strings.HasPrefix(wispID, "hw-") {
		t.Fatalf("formula order root id = %q, want rig prefix hw-\n%s", wispID, out)
	}
}

func TestOrderRunRigScopedExecUsesRigStore(t *testing.T) {
	cityDir, rigDir := setupOrderRunRigCity(t)
	writeRigExecOrder(t, rigDir, "rig-exec")

	out, err := gcDolt(cityDir, "order", "run", "--rig", "hello-world", "rig-exec")
	if err != nil {
		t.Fatalf("gc order run rig-exec failed: %v\n%s", err, out)
	}

	createdID := extractBeadID(t, out)
	if !strings.HasPrefix(createdID, "hw-") {
		t.Fatalf("exec order created bead id = %q, want rig prefix hw-\n%s", createdID, out)
	}
}

func setupOrderRunRigCity(t *testing.T) (cityDir, rigDir string) {
	t.Helper()

	t.Setenv("GC_SESSION", "subprocess")

	cityName := uniqueCityName()
	cityDir = filepath.Join(t.TempDir(), cityName)
	rigDir = filepath.Join(cityDir, "rigs", "hello-world")
	if err := os.MkdirAll(filepath.Join(rigDir, ".git"), 0o755); err != nil {
		t.Fatalf("creating rig git dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(rigDir, ".git", "config"), []byte("[core]\n\tbare = false\n"), 0o644); err != nil {
		t.Fatalf("writing rig git config: %v", err)
	}

	configPath := filepath.Join(t.TempDir(), "order-rig-city.toml")
	cityToml := "[workspace]\nname = \"order-rig-city\"\n\n[session]\nprovider = \"subprocess\"\n\n[[agent]]\nname = \"mayor\"\nstart_command = \"echo hello\"\n\n[[rigs]]\nname = \"hello-world\"\npath = \"rigs/hello-world\"\nformulas_dir = \"rigs/hello-world/formulas\"\n"
	if err := os.WriteFile(configPath, []byte(cityToml), 0o644); err != nil {
		t.Fatalf("writing city config: %v", err)
	}

	out, err := gcDolt("", "init", "--skip-provider-readiness", "--file", configPath, cityDir)
	if err != nil {
		t.Fatalf("gc init failed: %v\n%s", err, out)
	}
	t.Cleanup(func() {
		gcDolt("", "stop", cityDir) //nolint:errcheck // best-effort cleanup
	})

	return cityDir, rigDir
}

func writeRigFormulaOrder(t *testing.T, rigDir, orderName, formulaName, pool string) {
	t.Helper()

	formulasDir := filepath.Join(rigDir, "formulas")
	orderDir := filepath.Join(formulasDir, "orders", orderName)
	if err := os.MkdirAll(orderDir, 0o755); err != nil {
		t.Fatalf("creating order dir: %v", err)
	}

	formulaText := "formula = \"" + formulaName + "\"\nversion = 1\n\n[[steps]]\nid = \"work\"\ntitle = \"Work\"\n"
	if err := os.WriteFile(filepath.Join(formulasDir, formulaName+".formula.toml"), []byte(formulaText), 0o644); err != nil {
		t.Fatalf("writing formula: %v", err)
	}

	orderText := "[order]\nformula = \"" + formulaName + "\"\ngate = \"cooldown\"\ninterval = \"5m\"\npool = \"" + pool + "\"\n"
	if err := os.WriteFile(filepath.Join(orderDir, "order.toml"), []byte(orderText), 0o644); err != nil {
		t.Fatalf("writing order.toml: %v", err)
	}
}

func writeRigExecOrder(t *testing.T, rigDir, orderName string) {
	t.Helper()

	formulasDir := filepath.Join(rigDir, "formulas")
	orderDir := filepath.Join(formulasDir, "orders", orderName)
	scriptDir := filepath.Join(orderDir, "scripts")
	if err := os.MkdirAll(scriptDir, 0o755); err != nil {
		t.Fatalf("creating exec order script dir: %v", err)
	}

	script := "#!/bin/sh\nset -eu\nbd create 'Exec order created bead'\n"
	if err := os.WriteFile(filepath.Join(scriptDir, "create-bead.sh"), []byte(script), 0o755); err != nil {
		t.Fatalf("writing exec script: %v", err)
	}

	orderText := "[order]\nexec = \"$ORDER_DIR/scripts/create-bead.sh\"\ngate = \"cooldown\"\ninterval = \"5m\"\n"
	if err := os.WriteFile(filepath.Join(orderDir, "order.toml"), []byte(orderText), 0o644); err != nil {
		t.Fatalf("writing exec order.toml: %v", err)
	}
}

func extractOrderWispID(t *testing.T, output string) string {
	t.Helper()

	const prefix = "wisp "
	idx := strings.Index(output, prefix)
	if idx < 0 {
		t.Fatalf("could not find wisp id in output: %s", output)
	}
	rest := strings.TrimSpace(output[idx+len(prefix):])
	fields := strings.Fields(rest)
	if len(fields) == 0 {
		t.Fatalf("could not parse wisp id from output: %s", output)
	}
	return fields[0]
}
