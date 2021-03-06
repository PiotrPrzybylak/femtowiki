// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const baseSrc = `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" type="text/css" href="/static/css/femtowiki.css?v=010">
	<title>
		Femtowiki
	</title>
	{{ block "head" . }}{{ end }}
</head>
<body>
	<div id="header">
		<div class="logo"><a href="/">Femtowiki</a></div> <a href="">Home</a> <a href="">Download</a> <a href="">FAQ</a>
	</div>
	<div id="container">
		<div id="content">
			<div id="section-profile"><a href="/signup">Signup</a> <a href="/login">Login</a></div>
			{{ block "content" . }}{{ end }}
		</div>
		<div id="nav">
			<ul>
				<li>London</li>
				<li>New York</li>
				<li>Bangalore</li>
			</ul>
		</div>
		<div id="footer">
			Privacy Terms Help
		</div>
	</div>
	<script src="/static/js/femtowiki.js?v=010"></script>
</body>
</html>`