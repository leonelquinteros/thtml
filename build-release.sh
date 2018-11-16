go build -o release/darwin-amd64/thtml .
GOOS=android GOARCH=arm go build -o release/android-arm/thtml .
GOOS=darwin GOARCH=arm go build -o release/darwin-arm/thtml .
GOOS=darwin GOARCH=arm64 go build -o release/darwin-arm64/thtml .
GOOS=dragonfly GOARCH=amd64 go build -o release/dragonfly-amd64/thtml .
GOOS=freebsd GOARCH=386 go build -o release/freebsd-386/thtml .
GOOS=freebsd GOARCH=amd64 go build -o release/freebsd-amd64/thtml .
GOOS=freebsd GOARCH=arm go build -o release/freebsd-arm/thtml .
GOOS=linux GOARCH=386 go build -o release/linux-386/thtml .
GOOS=linux GOARCH=amd64 go build -o release/linux-amd64/thtml .
GOOS=linux GOARCH=arm go build -o release/linux-arm/thtml .
GOOS=linux GOARCH=arm64 go build -o release/linux-arm64/thtml .
GOOS=linux GOARCH=ppc64 go build -o release/linux-ppc64/thtml .
GOOS=linux GOARCH=ppc64le go build -o release/linux-ppc64le/thtml .
GOOS=linux GOARCH=mips go build -o release/linux-mips/thtml .
GOOS=linux GOARCH=mipsle go build -o release/linux-mipsle/thtml .
GOOS=linux GOARCH=mips64 go build -o release/linux-mips64/thtml .
GOOS=linux GOARCH=mips64le go build -o release/linux-mips64le/thtml .
GOOS=linux GOARCH=s390x go build -o release/linux-s390x/thtml .
GOOS=netbsd GOARCH=386 go build -o release/netbsd-386/thtml .
GOOS=netbsd GOARCH=amd64 go build -o release/netbsd-amd64/thtml .
GOOS=netbsd GOARCH=arm go build -o release/netbsd-arm/thtml .
GOOS=openbsd GOARCH=386 go build -o release/openbsd-386/thtml .
GOOS=openbsd GOARCH=amd64 go build -o release/openbsd-amd64/thtml .
GOOS=openbsd GOARCH=arm go build -o release/openbsd-arm/thtml .
GOOS=plan9 GOARCH=386 go build -o release/plan9-386/thtml .
GOOS=plan9 GOARCH=amd64 go build -o release/plan9-amd64/thtml .
GOOS=solaris GOARCH=amd64 go build -o release/solaris-amd64/thtml .
GOOS=windows GOARCH=386 go build -o release/windows-386/thtml.exe .
GOOS=windows GOARCH=amd64 go build -o release/windows-amd64/thtml.exe .

for i in release/*; do zip -r "$i".zip "$i"; done