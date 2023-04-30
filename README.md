[![wakatime](https://wakatime.com/badge/github/Uyutaka/introduction_to_testing_in_go.svg)](https://wakatime.com/badge/github/Uyutaka/introduction_to_testing_in_go)

Based on [this lecture](https://www.udemy.com/course/introduction-to-testing-in-go-golang/)

### Commands

```
# init module
go mod init github.com/uyutaka/xxxx/xxxx

# run test
go test -v . 

# show % of coverage
go test -cover .

# generate html about test coverage
go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

# run specific test
go test -run <test func>
```
