# TrueAccord

## Git repo

https://github.com/vbeffa/trueaccord

## Process and approach

I first got the debts endpoint consumed and output to the console, then did the same for the other endpoints. After that I refactored the code by moving the types out of the main function so they could be used by the helper funcs/methods. I then output the debts as described in the requirements, adding additional helpers and fixing a bug in the logic. Finally I added a test of the output using ginkgo. With more time I could have dockerized the solution so that go/ginkgo are not required to run/test it.

Note: I assumed biweekly means every two weeks, not twice a week.

## Running and testing

To run:

`go run cmd/trueaccord/main.go`

To test:

`ginkgo -v`
