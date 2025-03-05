package goconf

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestParseConstants(t *testing.T) {
	content := `
ConstA = "value1"
ConstB = "value2"
`
	tmpfile, err := ioutil.TempFile("", "testconstants*.go")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	constants, err := parseConstants(tmpfile.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if constants["ConstA"] != "value1" {
		t.Errorf("expected ConstA to be value1, got %s", constants["ConstA"])
	}
	if constants["ConstB"] != "value2" {
		t.Errorf("expected ConstB to be value2, got %s", constants["ConstB"])
	}
}

func TestRunUpdateConstantsDryRun(t *testing.T) {
	newContent := `NewConst = "common_value"`
	oldContent := `OldConst = "common_value"`

	newFile, err := ioutil.TempFile("", "new*.go")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(newFile.Name())
	if _, err := newFile.Write([]byte(newContent)); err != nil {
		t.Fatal(err)
	}
	newFile.Close()

	oldFile, err := ioutil.TempFile("", "old*.go")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(oldFile.Name())
	if _, err := oldFile.Write([]byte(oldContent)); err != nil {
		t.Fatal(err)
	}
	oldFile.Close()

	tempDir, err := ioutil.TempDir("", "updateconsttest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a sample .go file that uses the old constant.
	filePath := filepath.Join(tempDir, "sample.go")
	originalFileContent := "package main\nvar x = OldConst\n"
	err = os.WriteFile(filePath, []byte(originalFileContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Run in dry-run mode so the file should remain unchanged.
	if err := UpdateConstants(oldFile.Name(), newFile.Name(), true, false, []string{tempDir}); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != originalFileContent {
		t.Errorf("file content changed in dry-run mode")
	}
}

func TestBuildTreeAndFlatten(t *testing.T) {
	// Sample YAML to simulate config.
	yamlContent := `
key1:
  subkey1: value1
  subkey2: value2
# prefix: Custom
key2: value3
`
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlContent), &node)
	if err != nil {
		t.Fatal(err)
	}
	if len(node.Content) == 0 {
		t.Fatal("expected YAML document content")
	}
	tree := buildTree(node.Content[0])
	if tree.Children["key1"] == nil || tree.Children["key2"] == nil {
		t.Fatal("expected both key1 and key2 in tree")
	}
	leaves := flattenTree(tree, []string{}, "", "", false)
	if len(leaves) == 0 {
		t.Error("expected some leaves from flattenTree")
	}
}

func TestGenerateConstName(t *testing.T) {
	fr := flattenResult{
		Path:              []string{"Parent", "Child"},
		EffectiveOverride: "Override",
	}
	constName := generateConstName(fr)
	expected := "OverrideChild"
	if constName != expected {
		t.Errorf("expected %s, got %s", expected, constName)
	}
}
