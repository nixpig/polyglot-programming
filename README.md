# Projector

A simple CLI application that stores, deletes, or presents variables based on the current working directory or a path provided.

## Simple Example

### Print All Values

```
> cd /foo/bar
> /foo/bar> projector
> {}
```

### Add/Get Value

```
> /foo/bar> projector add foo bar
> /foo/bar> projector
> {"foo": "bar"}
> /foo/bar> projector foo
> bar
```

### Merging Data

```
> /foo/bar> cd baz
> /foo/bar/baz> projector
> {"foo": "bar"}
> /foo/bar/baz> projector foo
> bar
> /foo/bar/baz> projector add foo twitch
> /foo/bar/baz> projector
> {"foo": "twitch"}
> cd ..
> /foo/bar> projector
> {"foo": "bar"}
> /foo/bar> projector add bar baz
> /foo/bar> cd baz
> /foo/bar/baz> projector
> {
"foo": "twitch", # from /foo/bar/baz
"bar": "baz" # from /foo/bar
}
> /foo/bar/baz> projector bar
> baz
```

## Breaking the problem up

```
   +----------+    +----------+      +----------+    +----------+
   | cli opts | -> | project  | -+-> |  print   | -> | display  |
   +----------+    |  config  |  |   +----------+    +----------+
                   +----------+  |
                                 |   +----------+    +----------+
                                 +-> |   add    | -> |   save   |
                                 |   +----------+    +----------+
                                 |
                                 |   +----------+    +----------+
                                 +-> |    rm    | -> |   save   |
                                     +----------+    +----------+
```
