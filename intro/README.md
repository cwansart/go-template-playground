# Intro

First lection from https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/.

## How to run

Open the terminal in the intro folder and run:

```bash
$ go run .
```

## What you'll see

```html
title: Backslash! An in depth look ath the&#34;\&#34; character
html: &lt;h1&gt;A header&lt;/h1&gt;
SafeHTML: <h1>A Safe header
dot: {&lt;h1&gt;A header&lt;/h1&gt; &lt;h1&gt;A Safe header Backslash! An in depth look ath the&#34;\&#34; character /dashboard/settings {Malou 14} map[key:value other_key:other_value]}

<a title="Backslash! An in depth look ath the&#34;\&#34; character"></a>
<a title="&lt;h1&gt;A header&lt;/h1&gt;"></a>

<a href="%3ch1%3eA%20header%3c/h1%3e"></a>
<a href="?q=%3ch1%3eA%20header%3c%2fh1%3e"></a>
<a href="/dashboard/settings"></a>
<a href="?q=%2fdashboard%2fsettings"></a>

<script>
  var dog = {"Name":"Malou","Age":14};
  var map = {"key":"value","other_key":"other_value"};
  doWork("Backslash! An in depth look ath the\"\\\" character");
</script>
```

It demonstrates how the `template/html` package converts some characters depending on the context in templates.