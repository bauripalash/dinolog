# 🦕 Dinolog Protocol
Simple PlainText-based blogging protocol

# Simple Specification

### Blog Parameters
Each `dinolog.txt` file will have few mandatory fields (unless not `origin`)
* `title` => Title of the blog
* `author` => Author of the blog

Apart from these if the blog has the field `origin` the value of this field will 
considered as the first page of a pagination blog.

Fields will be like this and the space between the `=` character does not matter.

```bash
KEY = VALUE 
```
Before setting the values, extra spaces at the end and beginning will be 
trimmed. For example, `  key   =value` will be read as `key=value`.

### Follow (Blogroll)
Each text file can have a "follow" section which will mention the blogs 
they read from, thus it will act as a blogroll. Follow Block will be something
like this

```bash
----
@blogxyz = https://example.com/xyz.dinolog.txt 
@blogabc = https://example.org/blog/abc.dinolog.txt 
----
```

Follow Block starts with `----` (four dashes) and ends with the same four 
dashes. This block is optional but recommended. 

### Pagination
As dinolog is not only limited to micro-blogging or blogging, as a matter of 
fact, the text file can get very big very fast. For that blogs can have
pagination. It can suggest the next page as well as the previous page.

```bas 
>>>nextpage.dinolog.txt 
<<<https://example.org/previous.dinolog.txt
```

The text after `>>>` denotes the next page and the text after `<<<` denotes
the previous page. The **text** can be a relative or absolute URL/path.

Each of these pages will follow the same Specification as other individual
blogs. But it is recommended to put an `origin` field in the blog parameters 
section mentioning the first page.

### Posts
Each post starts with a date field. Date field starts with  `[[` and ends with 
`]]` and the text between them must be IETF RFC 3339 [^1] formatted date.
After that line, every line will be considered as the part of blog post's body.
A line with only `--0--` marks the end of a post.

A single dinolog.txt can hold as many blog posts as the owner likes. But 
It is good practice to use Pagination to make the text file lightweight. 


---
[^1]:[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)


