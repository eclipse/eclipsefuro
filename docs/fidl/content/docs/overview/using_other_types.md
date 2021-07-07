---
weight: 45
title: "Using Other Types"
---
# Using Other Types
You can use other types as field types. You can use every type from your specs and installed
dependencies. You do not need to import them. The imports are resolved and checked by Furo when you translate your µSecs to standard specs with the command `furo muSpec2spec` or with the command `muspec checkImpors`.

{{<hint warning>}}
By using types which are not installed in the dependencies or not from your spec project, do not forget to import them in your protoc command with "-I". 

When you need this types on the client side you also need to import them separately. 
{{</hint>}}


For example, let's say you wanted to include sample.Details in sample.Sample on the field `details`
 – to do this, you can define another type in the same type µSpec. In standard spec you have to use 2 files, 
because the standard spec can only handle 1 type per file due to historical reasons:

*File: muspec/sample.types.yaml*
```yaml
- type: 'sample.Sample #A sample type'
  fields:
    password: '* string:1 #The password.'
    username: '* string:2 #The username or email, or something to identify.'
    details: 'sample.Details:3 #Details.'
  
- type: 'sample.Details #A sample type'
  fields:
    birth_date: '* google.type.Date:1 #The birth date.'
    weight: '* number:2 #The weight.'    
    age: '- number:3 #Calculated field for displaying the age, because the calculations are very hard.'
 

```



