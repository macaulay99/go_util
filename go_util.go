package go_util

import (
        "os"
        "log"
        "encoding/json"
)

var Config Configuration
var Logger *log.Logger

type Configuration struct {
    IpAddress    string
    TcpPort      string
    ReadTimeout  int64
    WriteTimeout int64
    Static       string
}


func OpenLog(logFileName string) int {
    file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file", err)
    }
    Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
    return(0)
}

func LoadConfig(configFileName string) int {
    file, err := os.Open(configFileName)
    if err != nil {
        log.Fatalln("Cannot open config file", err)
    }
    decoder := json.NewDecoder(file)
    Config = Configuration{}
    err = decoder.Decode(&Config)
    if err != nil {
        log.Fatalln("Cannot get configuration from file", err)
    }
    return(0)
}
