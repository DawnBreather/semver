package main

import (
  "flag"
  "os"
  "path/filepath"
  "regexp"
  "semver/semver"
)

func main() {
  var increaseBy, prefix, filenamePattern, fsPath string
  flag.StringVar(&increaseBy, "increase-version-by-step", "", "Version to increase by")
  flag.StringVar(&prefix, "string-regexp-prefix", "", "Regexp prefix for version string")
  flag.StringVar(&filenamePattern, "filename-regexp-pattern", "", "Filename regexp pattern")
  flag.StringVar(&fsPath, "recursive-filesystem-path", "./", "Filesystem path to start the search")
  flag.Parse()

  filepath.Walk(fsPath, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    return processPath(path, filenamePattern, increaseBy, prefix)
  })
}

func processPath(path string, filenamePattern, increaseBy, prefix string) error {
  info, err := os.Stat(path)
  if err != nil {
    return err
  }

  matched, _ := regexp.MatchString(filenamePattern, info.Name())
  if !info.IsDir() && matched {
    content, err := os.ReadFile(path)
    if err != nil {
      return err
    }

    newContent, err := semver.ProcessFileContent(string(content), increaseBy, prefix)
    if err != nil {
      return err
    }

    if string(content) != newContent {
      return os.WriteFile(path, []byte(newContent), 0644)
    }
  }
  return nil
}
