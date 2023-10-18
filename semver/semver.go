package semver

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

type SemVer struct {
  Major, Minor, Patch int
}

func ParseSemVer(version string) (SemVer, error) {
  parts := strings.Split(version, ".")
  if len(parts) != 3 {
    return SemVer{}, fmt.Errorf("invalid semver format")
  }

  major, err := strconv.Atoi(parts[0])
  if err != nil {
    return SemVer{}, err
  }
  minor, err := strconv.Atoi(parts[1])
  if err != nil {
    return SemVer{}, err
  }
  patch, err := strconv.Atoi(parts[2])
  if err != nil {
    return SemVer{}, err
  }

  return SemVer{Major: major, Minor: minor, Patch: patch}, nil
}

func (s *SemVer) IncrementBy(step SemVer) {
  s.Major += step.Major
  s.Minor += step.Minor
  s.Patch += step.Patch
}

func (s SemVer) String() string {
  return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func ProcessFileContent(content, increaseBy, prefix string) (string, error) {
  step, err := ParseSemVer(increaseBy)
  if err != nil {
    return content, err
  }

  regex := regexp.MustCompile(prefix + `.*` + `(\d+\.\d+\.\d+)`)
  newContent := regex.ReplaceAllStringFunc(content, func(s string) string {
    match := regex.FindStringSubmatch(s)
    if len(match) < 2 {
      return s
    }

    oldVersion, err := ParseSemVer(match[1])
    if err != nil {
      return s
    }

    oldVersion.IncrementBy(step)
    return strings.Replace(s, match[1], oldVersion.String(), 1)
  })

  return newContent, nil
}
