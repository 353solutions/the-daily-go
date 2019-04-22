# Possible Exercises

## Document Summary

Use the following algorithm to add a `Summarize(text string, count int)
[]string` method to our `nlp` package.

1. For every word in the sentence count how many times it appears in the document.
    - This is known as `token frequency` or
      [`tf`](https://en.wikipedia.org/wiki/Tf%E2%80%93idf).
2. For every sentence in the document, give is a score which is the sum of all the token frequencies for the tokens in the sentence.
    - e.g. if we have the following token frequencies: please=1, reinstall=3,
      universe=3, and=10, reboot=2. Then the score of the sentence `Please
      reinstall universe and reboot` is `1 + 3 + 3 + 10 + 2 = 19`
3. Return the `count` top scored sentences (this is known as `sentence extraction`)

### Tasks
- Test your function. Have directory with texts to summarize and expected output.
- Create testable example for `Summarize`
- Create a benchmark for your server under load, you can use
  [vegeta](https://github.com/tsenart/vegeta) to generate load
- Tag a new version of your code. Create a test to see it's `go get`able
- Publish a docker image for `nlpd` to docker hub
- Add [gRPC](https://grpc.io/docs/tutorials/basic/go.html) front end `nlpd`.
  Have an option to run either HTTP server or gRPC or both.
    - Make sure metrics still work in `gRPC` only mode
- Head over to [godoc](https://godoc.org/) and check how your package
  documentation looks. See it in the eyes of a new user and fix it
