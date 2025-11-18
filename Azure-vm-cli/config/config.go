package config


import (
"os"
"path/filepath"
)


func ConfigDir() string {
home, err := os.UserHomeDir()
if err != nil {
return "."
}
return filepath.Join(home, ".config", "azvmctl")
}


func NewConfig() *struct{} {
return &struct{}{}
}
