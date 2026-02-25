package main

import (
"bufio"
"crypto/sha256"
"encoding/hex"
"fmt"
"io"
"os"
"path/filepath"
"strings"
)

const (
ExitOK   = 0
ExitFail = 10
)

func sha256File(path string) (string, error) {
f, err := os.Open(path)
if err != nil {
return "", err
}
defer f.Close()

h := sha256.New()
if _, err := io.Copy(h, f); err != nil {
return "", err
}
return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
repoRoot, err := os.Getwd()
if err != nil {
fmt.Println("verdict=FAIL reason=getwd_error")
os.Exit(ExitFail)
}

manifestRel := filepath.FromSlash("spec/sha256_manifest.v0.1.0.txt")
manifestPath := filepath.Join(repoRoot, manifestRel)

mf, err := os.Open(manifestPath)
if err != nil {
fmt.Printf("verdict=FAIL reason=manifest_missing file=%s\n", manifestRel)
os.Exit(ExitFail)
}
defer mf.Close()

sc := bufio.NewScanner(mf)
lineNo := 0
checked := 0

for sc.Scan() {
lineNo++
line := strings.TrimSpace(sc.Text())
if line == "" {
continue
}
// format: "<sha256>  <relative_path>"
parts := strings.SplitN(line, "  ", 2)
if len(parts) != 2 {
fmt.Printf("verdict=FAIL reason=bad_manifest_format line=%d\n", lineNo)
os.Exit(ExitFail)
}
want := strings.ToLower(strings.TrimSpace(parts[0]))
rel := strings.TrimSpace(parts[1])

// Normalize Windows backslashes that may exist in manifest
rel = strings.ReplaceAll(rel, "/", string(os.PathSeparator))
abs := filepath.Join(repoRoot, rel)

if _, err := os.Stat(abs); err != nil {
fmt.Printf("verdict=FAIL reason=file_missing file=%s\n", rel)
os.Exit(ExitFail)
}

got, err := sha256File(abs)
if err != nil {
fmt.Printf("verdict=FAIL reason=hash_read_error file=%s\n", rel)
os.Exit(ExitFail)
}
if got != want {
fmt.Printf("verdict=FAIL reason=sha_mismatch file=%s want=%s got=%s\n", rel, want, got)
os.Exit(ExitFail)
}
checked++
}
if err := sc.Err(); err != nil {
fmt.Println("verdict=FAIL reason=manifest_read_error")
os.Exit(ExitFail)
}

fmt.Printf("verdict=PASS reason=manifest_ok checked=%d\n", checked)
os.Exit(ExitOK)
}
