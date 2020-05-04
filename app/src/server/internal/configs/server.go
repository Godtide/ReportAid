package configs

import (
    "time"
)

const (
    // ServerPort - port
    ServerPort  = ":10000"
    // WriteTimeout - disallow writes
    WriteTimeout     = 0 * time.Second
    // ReadTimeout - reads from blockchain can be slow
    ReadTimeout     = 60 * time.Second
)
