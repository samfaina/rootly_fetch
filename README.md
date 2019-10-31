# rootly_fetch

Clone trefle.io plants database

# build commands
* Windows amd64: ```env GOOS=windows GOARCH=amd64 go build -o builds/rootly_win64.exe``` 
* Windows 386: ```env GOOS=windows GOARCH=386 go build -o builds/rootly_win64.exe``` 
* Linux amd64: ```env GOOS=linux GOARCH=amd64 go build -o builds/rootly_linin64``` 
* Linux 386: ```env GOOS=linux GOARCH=386 go build -o builds/rootly_lin32``` 
* Linux arm: ```env GOOS=linux GOARCH=arm go build -o builds/rootly_lin_arm``` 
* Darwin amd64: ```env GOOS=darwin GOARCH=amd64 go build -o builds/rootly_osx64``` 
* Darwin 386: ```env GOOS=darwin GOARCH=386 go build -o builds/rootly_osx32``` 
* Darmac arm: ```env GOOS=darwin GOARCH=arm go build -o builds/rootly_osx_arm``` 
* Android arm: ```env GOOS=android GOARCH=arm go build -o builds/rootly_and64.apk``` 
* Solaris amd64: ```env GOOS=solaris GOARCH=amd64 go build -o builds/rootly_sol64``` 
* Plan9 amd64: ```env GOOS=plan9 GOARCH=amd64 go build -o builds/rootly_pla64``` 
* Plan9 amd386: ```env GOOS=solaris GOARCH=386 go build -o builds/rootly_pla32``` 
* OpenBSD amd64: ```env GOOS=openbsd GOARCH=amd64 go build -o builds/rootly_obsd64``` 
* OpenBSD 386: ```env GOOS=openbsd GOARCH=386 go build -o builds/rootly_obsd32``` 
* OpenBSD arm: ```env GOOS=openbsd GOARCH=arm go build -o builds/rootly_obsd_arm``` 
* NetBSD amd64: ```env GOOS=netbsd GOARCH=amd64 go build -o builds/rootly_nbsd64``` 
* NetBSD 386: ```env GOOS=netbsd GOARCH=386 go build -o builds/rootly_nbsd32``` 
* NetBSD arm: ```env GOOS=netbsd GOARCH=arm go build -o builds/rootly_nbsd_arm``` 
* FreeBSD amd64: ```env GOOS=freebsd GOARCH=amd64 go build -o builds/rootly_fbsd64``` 
* FreeBSD 386: ```env GOOS=freebsd GOARCH=386 go build -o builds/rootly_fbsd32``` 
* FreeBSD arm: ```env GOOS=freebsd GOARCH=arm go build -o builds/rootly_fbsd_arm``` 