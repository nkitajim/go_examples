[UDHR](https://oreil.ly/-q7Cn) Download

# staticcheck
```
% go install honnef.co/go/tools/cmd/staticcheck@latest
% staticcheck 
main.go:47:10: should use fmt.Errorf(...) instead of errors.New(fmt.Sprintf(...)) (S1028)
```

# cross compile
```
% GOOS=linux GOARCH=amd64 go build 
% GOOS=linux GOARCH=arm64 go build 
% GOOS=windows GOARCH=386 go build
% GOOS=darwin GOARCH=amd64 go build
% GOOS=darwin GOARCH=arm64 go build
```
