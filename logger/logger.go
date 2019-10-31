package logger

import (
	"io"
	"log"
	"os"
)

func SetupLogger() {
	logFile, err := os.OpenFile("rootly.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

}

// func getLogger() *log.Logger {
// 	logFile, err := os.OpenFile("rootly.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return log.New(logFile, "", log.LstdFlags)
// }

// // Println log
// func Println(v ...interface{}) {
// 	getLogger().Println(v...)
// }

// // Printf log
// func Printf(format string, v ...interface{}) {
// 	getLogger().Printf(format, v...)
// }

// // Fatalln log
// func Fatalln(v ...interface{}) {
// 	getLogger().Fatalln(v...)
// }

// // Fatalf log
// func Fatalf(format string, v ...interface{}) {
// 	getLogger().Fatalf(format, v...)
// }
