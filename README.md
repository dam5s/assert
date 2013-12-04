Assert
======

Assert is an assertion library for Go

The goal is to be able to write simple assertions in your tests as follows:

```
assert.That(t, foo).IsEqualTo(bar)
```

Syntax should avoid confusion (what order should the parameters be?) and it should output meaningful messages.
It should always be clear what object the assertion is on, and what is the expected value.

