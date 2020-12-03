Exposes local files over http
=============================

This command exposes local files over http.

Motivation
----------

Popular browsers doen't work when schema is file(file://path/to/your/file.html).
For example, "The image element contains cross-origin data, and may not be loaded."

The problem can be solved by running Apache Httpd, Nginx and more.
But it's overkill to use them to open local html files.
Because they are big and very complicated to set up.

This command provides very very simple way to expose local file over http.

How to Use
----------

if you want to expose index.html.

```
cd /path/to/directory
expose
```

access to http://localhost:8080/index.html

How to Install
--------------

Download a release file and ungzip it.

https://github.com/yuuichi-fujioka/expose/releases/tag/v0.0.1

Misc
----

You can change bind addr.

```
expose -bind 127.0.0.1:9999
```