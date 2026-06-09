---
theme: ./theme.json
author: Mateusz Hawrus
date: February 13, 2025
paging: Slide %d / %d
---

# How I learned to stop worrying and love generics in Go

Exploring validation API solutions in modern Go.

---

## Who am I?

👨 💼

### Where can you find me?

1. https://github.com/nieomylnieja/
2. https://www.linkedin.com/in/mateusz-hawrus-038463174/

---

## Topic

What this presentation is about?

- Writing (runtime) validation in Go.
- Using generics to create a validation API.
- Govy!

What this presentation is **NOT** about?

- Drum roll... Generics!
- Writing static validation :)

### Legend

1. The basics, how would one go about writing validation code. The simple way.
2. Getting fancier, utilizing reflection.
3. Enter generics, combining type safety and ergonomics.
4. Beyond type safety. What _govy_ brings to the table?

---

## Disclaimer

Code snippets in this presentation have been trimmed, removing some parts of the code like:

- package declaration
- imports
- repetitive code (introduced in previous slides)

If you see `// ...` comment, it means I only trimmed some part of a declaration.

Furthermore, I have introduced questionable code in terms of:

- its logic
- runtime safety (casting)
- formatting

All that in an effort to make the snippets more concise.

---

## KISS 💋

Keep it simple, stupid!

If you need to check just a couple of things and the validation logic is simple, you're most likely better off writing the checks yourself :)

Let's explore how that would look like!

---

### Simple checks

```go
~~~./code-to-slide.awk code/01-simple/main.go

~~~
```

---

### Simple checks aggregated

What if we want to aggregate all the errors?

```go
~~~./code-to-slide.awk code/02-simple-aggregate/main.go

~~~
```

NOTE: There's an error here! (can you spot it?)

---

### Simple checks with slice of students

```go
~~~./code-to-slide.awk code/03-simple-slice/main.go

~~~
```

---

### Simple approach overview

#### Summary:

1. Works perfectly fine with simple structures.
2. It's as fast as it gets.
3. Does not provide a level of abstraction sufficient enough to handle complex validation logic and multiple entities.
4. Keeping consistent messages across multiple APIs can become a challenge. \
.  What if our models' names change? \
.  Example: `Teacher.Name` --> `Teacher.FirstName` \
.  Can we make sure messages reflect this change?

#### Things we could've improved upon:

1. DRY, extract common functions for validating specific fields.
2. Flow control, whether to fail fast or aggregate all errors.

If we go with these improvements, we're essentially creating a new validation library.
One might argue, it would no longer be _the simple_ approach.

---

## Reflection

Most of the established validation libraries in Go's ecosystem rely on reflection.
On top of that, many solutions utilise struct tags for associating rules with specific struct fields.

In the following slides we'll explore how such an API could be written.

---

### Reflection based checks

```go
~~~./code-to-slide.awk code/04-reflection/main.go

~~~
```

---

### Using struct tags

```go
~~~./code-to-slide.awk code/05-struct-tags/main.go

~~~
```

---

### Reflection based validation overview

#### Advantages over simple validation:

1. Scales to large structures.
2. High reusability.
3. Keeps error messages consistent:
    - reusability results in the messages consistency
    - properties' names in messages reflect code state
4. Easy of use with struct tags.

#### Problems:

1. ❗**We're no longer type safe!** ❗
2. Runtime overhead incurred by reflection.

---

## Generics

Before generics, there really wasn't any other way to write a robust validation API without reflection.

However, since Go 1.18 we're no longer limited!

In the following slides we'll explore how generics in Go allow us to design a reusable validation API from scratch :)

---

### Type definitions

Let us first define basic building blocks for our API.

```go
~~~./code-to-slide.awk code/06-generics-type-definitions/main.go

~~~
```

---

### Validation functions

Now, we'll define validation functions for each of the previously defined building blocks.

```go
~~~./code-to-slide.awk code/07-generics-functions/main.go

~~~
```

---

### Brining it all together

```go
~~~./code-to-slide.awk code/08-generics-full/main.go

~~~
```

---

### Generics overview

#### Advantages over reflection based validation:

1. **We're type safe!**
2. There's no runtime overhead coming from reflection anymore.

#### Problems:

1. The API might be harder to write implementation wise.
2. The code our users write will be more verbose.
3. We have to define field names manually.

---

## Thus govy was born

The three building blocks we've seen are the backbone of govy.

```
~~~graph-easy --as=boxart
[ Validator\[S\] ] - aggregates -> [ FieldRules\[T, S\] ]  - aggregates -> [ Rule\[T\] ]
~~~
```

- `Validator` instance ties our rules to specific struct type `S`
- `FieldRules` defines a single struct field, it is a glue between `Validator` struct type `S` and `Rule` value of type `T`
- `Rule` defines a single rule for type `T`

### What else does it bring to the table?

1. Advanced flow control (error cascade modes).
2. Conditional execution (predicates).
3. Powerful composition options (include validators).
4. Functional API with its main building blocks being immutable.
5. Error message templates.
6. Validation plan.
7. Automatic property name inference (experimental).
8. And more!

---

### Nuts which govy cracks!

Let's first see what govy was designed to validate:

```yaml
apiVersion: n9/v1alpha
kind: SLO
metadata:
  name: api-server-slo
  displayName: API Server SLO
  project: default
spec:
  description: Example Prometheus SLO
  indicator:
    metricSource:
      name: prometheus
      project: default
      kind: Agent
  budgetingMethod: Occurrences
  objectives:
  - displayName: Good response (200)
    value: 200.0
    name: ok
    target: 0.95
    rawMetric:
      query:
        prometheus:
          promql: api_server_requestMsec{host="*",job="nginx"}
    op: lte
    primary: true
  service: api-server
  timeWindows:
  - unit: Month
    count: 1
    isRolling: false
    calendar:
      startTime: "2022-12-01 00:00:00"
      timeZone: UTC
```

---

### Govy usage example

Let's explore how govy is used in the wild :)

- https://github.com/nobl9/govy/blob/main/docs/validator-comparison/example_test.go
- https://github.com/OpenSLO/go-sdk/blob/main/pkg/openslo/v1/service.go

---

### Trivia

Why have we burned time and money to create it?
Was it only about type safety?

Truth be told, the main reason had nothing to do with type safety.
Our main pain point with `go-playground/validator` library (which at that time we used) were the error messages.

`validator` library produces very poor messages which were not user friendly and often left our customers clueless as to what they've done wrong.

For instance:

> $ Key: 'Service.kind' Error:Field validation for 'kind' failed on the 'eq' tag

Compared to what govy would spit:

> $ Validation for Service has failed for the following properties:
> $     - 'kind' with value 'Project':
> $         - should be equal to 'Service'

In time however, we discovered govy can help us achieve much more, especially with its validation plan.
We could now generate documentation from our code and allow other tools to utilise that information (like IDE integrations).

#### Alternatives

JSON schema was a consideration, in fact I've done a PoC for our most complex API object --> SLO, and I ended up with a monster...

Cue was another candidate, although it lacked robust cross field validation options.

---

## Wrapping up

Thank you for your attention!

The presentation can be found at https://github.com/nieomylnieja/go-meetup-govy-presentation.

### More on govy

Repository link: https://github.com/nobl9/govy
Blog post: https://www.nobl9.com/resources/type-safe-validation-in-go-with-govy

### Real life usage

- https://github.com/nobl9/nobl9-go
- https://github.com/OpenSLO/go-sdk
