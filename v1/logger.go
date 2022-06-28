package proxmox

type Logger interface {
    Trace(args ...interface{})
    Tracef(format string, args ...interface{})

    Debug(args ...interface{})
    Debugf(format string, args ...interface{})

    Info(args ...interface{})
    Infof(format string, args ...interface{})
    
    Warn(args ...interface{})
    Warnf(format string, args ...interface{})
}
