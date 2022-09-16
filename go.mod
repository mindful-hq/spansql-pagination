module github.com/mindful-hq/spansql-pagination

go 1.19

replace golang.org/x/tools => golang.org/x/tools v0.1.12 // https://github.com/bazelbuild/rules_go/issues/3230#issuecomment-1216728711

require cloud.google.com/go/spanner v1.38.0

require cloud.google.com/go v0.102.1 // indirect
