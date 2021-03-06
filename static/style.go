// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package static

const StyleSrc = `
* {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
}
html, body {
	margin: 0;
	padding: 0;
}
body {
	font-family: Arial, "Helvetica Neue", Helvetica, sans-serif;
	-webkit-font-smoothing: antialiased;
	text-rendering: optimizeLegibility;
	line-height: 1.58;
	margin-bottom: 50px;
	background-color: #f6f6f6;
}
.clearfix {
	overflow: auto;
	zoom: 1;
}
a {
	text-decoration: none;
}
a:link {
	color: #07C;
}
a:hover, a:active {
	color: #3af;
}
a:visited {
	color: #005999;
}
#container {
	position: relative;
	padding: 10px;
}
#header {
	text-align: right;
	padding: 18px;
	background-color: #333333;
	border-bottom: 5px #08C solid;
}
#header .logo {
	float: left;
}
#header a {
	color: #FFFFFF;
	font-size: 16px;
	padding: 0 10px;
	font-weight: bold;
}
#header a:link, #header a:visited, #header a:hover, #header a:active {
	text-decoration: none;
}
#section-profile {
	text-align: right;
	padding: 10px 0px;
}
#section-profile a {
	padding: 0 10px;
}
@media screen and (min-width:600px) {
	#header a {
		padding: 0 20px;
	}
	#nav {
		position: absolute;
		top: 0px;
		width: 180px;
	}
	#content {
		margin-left: 200px;
		min-height: 480px;
	}
	#footer {
		margin-left: 200px;
	}
}
`