# Checkit

## Inspiration

This is inspired by Checkit that written in JS. You can check out the original lib below
<br/>
https://github.com/tgriesser/checkit

I'm interested in Golang while I'm doing a blockchain platform project. My responsibility is contributing a RESTFul module to communicate with the core blockchain
<br/>
I wrote this lib with the first idea that it would help me validate request body.
<br/>
It allows you to seamlessly validate full Golang structs, primitive type values, defining custom messages, labels, and validations.

## Installation

You just open `Terminal` and run this command
```
go get https://github.com/dungntm58/checkit
```
to inject it as a dependency

## Example

### Validate complex structure object

```Golang
r, err := Validator(map[string]Validating{
  "a": Between(0, 2),
  "b": MaxLength(2),
  "c": ExactLength(3),
}).ValidateSync(struct {
  a int
  b *string
  c []interface{}
}{
  a: 1,
  b: &str,
  c: []interface{}{0, 1, 2},
})
fmt.Println(r) // true
```

### Validate single value
```Golang
r, err := Integer().Validate(1)
fmt.Println(r) // true
```

## Available Validators

<table>
  <thead>
    <tr>
      <th style="min-width:250px;">Validation Name</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Accepted</td>
      <td>The value must be yes, on, or 1. This is useful for validating "Terms of Service" acceptance.</td>
    </tr>
    <tr>
      <td>Alpha</td>
      <td>The value must be entirely alphabetic characters.</td>
    </tr>
    <tr>
      <td>AlphaDash</td>
      <td>The value may have alpha-numeric characters, as well as dashes and underscores.</td>
    </tr>
    <tr>
      <td>AlphaNumeric</td>
      <td>The value must be entirely alpha-numeric characters.</td>
    </tr>
    <tr>
      <td>AlphaUnderscore</td>
      <td>The value must be entirely alpha-numeric, with underscores but not dashes.</td>
    </tr>
    <tr>
      <td>Array</td>
      <td>The value must be a valid array object.</td>
    </tr>
    <tr>
      <td>Base64</td>
      <td>The value must be a base64 encoded value.</td>
    </tr>
    <tr>
      <td>Between:min:max</td>
      <td>The value must have a size between the given min and max.</td>
    </tr>
    <tr>
      <td>Boolean</td>
      <td>The value must be a javascript boolean.</td>
    </tr>
    <tr>
      <td>Contains:value</td>
      <td>The value must be a string or an array and contain the value.</td>
    </tr
    <tr>
      <td>Date</td>
      <td>The value must be a valid date object.</td>
    </tr>
    <tr>
      <td>Email</td>
      <td>The field must be a valid formatted e-mail address.</td>
    </tr>
    <tr>
      <td>Empty</td>
      <td>The value under validation must be empty; either an empty string, an empty, array, empty object, or a falsy value.</td>
    </tr>
    <tr>
      <td>ExactLength:value</td>
      <td>The field must have the exact length of "val".</td>
    </tr>
    <tr>
      <td>ExistsNonNil</td>
      <td>The value under validation must be exist or not nil.</td>
    </tr>
    <tr>
      <td>Finite</td>
      <td>The value under validation must be a finite number.</td>
    </tr>
    <tr>
      <td>Function</td>
      <td>The value under validation must be a function.</td>
    </tr>
    <tr>
      <td>GreaterThan:value</td>
      <td>The value under validation must be "greater than" the given value.</td>
    </tr>
    <tr>
      <td>GreaterThanEqualTo:value</td>
      <td>The value under validation must be "greater than" or "equal to" the given value.</td>
    </tr>
    <tr>
      <td>Integer</td>
      <td>The value must have an integer value.</td>
    </tr>
    <tr>
      <td>Ipv4</td>
      <td>The value must be formatted as an IPv4 address.</td>
    </tr>
    <tr>
      <td>Ipv6</td>
      <td>The value must be formatted as an IPv6 address.</td>
    </tr>
    <tr>
      <td>LessThan:value</td>
      <td>The value must be "less than" the specified value.</td>
    </tr>
    <tr>
      <td>LessThanEqualTo:value</td>
      <td>The value must be "less than" or "equal to" the specified value.</td>
    </tr>
    <tr>
      <td>Luhn</td>
      <td>The given value must pass a basic luhn (credit card) check regular expression.</td>
    </tr>
    <tr>
      <td>Max:value</td>
      <td>The value must be less than a maximum value. Strings, numerics, and files are evaluated in the same fashion as the size rule.</td>
    </tr>
    <tr>
      <td>MaxLength:value</td>
      <td>The value must have a length property which is less than or equal to the specified value. Note, this may be used with both arrays and strings.</td>
    </tr>
    <tr>
      <td>Min:value</td>
      <td>The value must have a minimum value. Strings, numerics, and files are evaluated in the same fashion as the size rule.</td>
    </tr>
    <tr>
      <td>MinLength:value</td>
      <td>The value must have a length property which is greater than or equal to the specified value. Note, this may be used with both arrays and strings.</td>
    </tr>
    <tr>
      <td>NaN</td>
      <td>The value must be <tt>NaN</tt>.</td>
    </tr>
    <tr>
      <td>Natural</td>
      <td>The value must be a natural number (a number greater than or equal to 0).</td>
    </tr>
    <tr>
      <td>NaturalNonZero</td>
      <td>The value must be a natural number, greater than or equal to 1.</td>
    </tr>
    <tr>
      <td>Object</td>
      <td>The value must be anything except functions, pointers.</td>
    </tr>
    <tr>
      <td>PlainObject</td>
      <td>The value must be a map.</td>
    </tr>
    <tr>
      <td>Regex</td>
      <td>The value must be a Go <tt>RegExp</tt> object.</td>
    </tr>
    <tr>
      <td>String</td>
      <td>The value must be a string type.</td>
    </tr>
    <tr>
      <td>URL</td>
      <td>The value must be formatted as an URL.</td>
    </tr>
    <tr>
      <td>UUID</td>
      <td>Passes for a validly formatted UUID.</td>
    </tr>
  </tbody>
</table>
