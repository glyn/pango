# sago - Static Analysis for Go modules

Some approaches for developing Go code, particularly TDD, favour many small types composed together. But when you need to grok the overall structure, it can be laborious to wade through the code.

This is where sago comes in - it discovers modules consisting of one or more packages and analyses their interdependencies.