package secretsafe

import (
    "fmt"
    "strconv"
    "strings"
)

type Version struct {
    Major int
    Minor int
    Patch int
}

func ParseVersion(v string) (*Version, error) {
    parts := strings.Split(v, ".")
    if len(parts) != 3 {
        return nil, fmt.Errorf("invalid version format: %s", v)
    }

    major, err := strconv.Atoi(parts[0])
    if err != nil {
        return nil, fmt.Errorf("invalid major version: %s", parts[0])
    }

    minor, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil, fmt.Errorf("invalid minor version: %s", parts[1])
    }

    patch, err := strconv.Atoi(parts[2])
    if err != nil {
        return nil, fmt.Errorf("invalid patch version: %s", parts[2])
    }

    return &Version{Major: major, Minor: minor, Patch: patch}, nil
}

func (v *Version) String() string {
    return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) Increment(incrementType string) {
    switch incrementType {
    case "major":
        v.Major++
        v.Minor = 0
        v.Patch = 0
    case "minor":
        v.Minor++
        v.Patch = 0
    case "patch":
        v.Patch++
    }
}

func (s *SecretStore) IncrementVersion(incrementType string) error {
    version, err := ParseVersion(s.Version)
    if err != nil {
        // If parsing fails, initialize with a default version
        version = &Version{Major: 1, Minor: 0, Patch: 0}
    }

    version.Increment(incrementType)
    s.Version = version.String()
    return nil
}