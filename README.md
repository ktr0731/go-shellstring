go-shellstring
===


``` 
"foo", []string{"foo"}
"foo bar", []string{"foo", "bar"}
`"foo"`, []string{"foo"}, false}
`"foo bar"`, []string{"foo bar"}
`"foo" "bar"`, []string{"foo", "bar"}
`'foo'`, []string{"foo"}, false}
`'foo bar'`, []string{"foo bar"}
`'foo' 'bar'`, []string{"foo", "bar"}
`'foo' 'bar'`, []string{"foo", "bar"}
```
