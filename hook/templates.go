package hook

const (
	postcardDefaultBack = `<html>
<head>
<meta charset="UTF-8">
<link href="https://fonts.googleapis.com/css?family=Allura" rel="stylesheet">
<style>
  *, *:before, *:after {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
  body {
    width: 6.25in;
    height: 4.5in;
    margin: 0;
    padding: 0;
  }
  #outer-wrapper {
    width: 6.25in;
    height: 4.5in;
    background-image: url("https://dmprint.s3.amazonaws.com/resources/template/html/birthday/bigpresent.jpg");
    background-size: 6.25in 4.5in;
    -webkit-filter: blur(3px);
    -moz-filter: blur(3px);
    -o-filter: blur(3px);
    -ms-filter: blur(3px);
    filter: blur(3px);
  }
  #text-area {
    position: absolute;
    width: 5.875in;
    height: 3.875in;
    left: 0.1875in;
    top: 0.25in;
    padding: .1in;
    text-align: left;
    font-size: auto;
    color: #fff;
  }
  pre {
    white-space: pre-wrap;
  }
</style>
</head>
<body>
  <div id="outer-wrapper">
  </div>
  <div id="text-area">
    <pre>{{logMessage}}</pre>
  </div>
</body>
</html>`

	postcardDefaultFront = `<html>
<head>
<meta charset="UTF-8">
<link href="https://fonts.googleapis.com/css?family=Quicksand" rel="stylesheet">
<style>
  *, *:before, *:after {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
  body {
    width: 6.25in;
    height: 4.5in;
    margin: 0;
    padding: 0;
  }
  #safe-area {
    position: absolute;
    width: 5.875in;
    height: 3.875in;
    left: 0.1875in;
    top: 0.1875in;
  }
  #present {
    background-image: url("https://dmprint.s3.amazonaws.com/resources/template/html/birthday/present.jpg");
    background-size: 6.25in 1.5in;
    width:6.25in;
    height:1.5in;
  }
  #message {
    position: absolute;
    width: 2.0in;
    height: 2in;
    top: 1.4in;
    left: .25in;
    font-family: 'Quicksand', sans-serif;
    font-size: 0.12in;
  }
</style>
</head>
<body>
  <div id="present">
  </div>
  <div id="safe-area">
    <div id="message">
  	  Greetings {{recipientName}},
	  <br><br>
	  Here is a log message from <b>{{appName}}</b>.
	  I hope it's useful.
  	  <br><br>
	  Sincerely,
	  <br><br>
	  Your {{appName}} logger
    </div>
  </div>
</body>
</html>`

	letterDefault = `<html>
  <head>
  	<link href="https://fonts.googleapis.com/css?family=Quicksand" rel="stylesheet">
    <meta charset="UTF-8">
    <style>
        *, *:before, *:after {
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;
        box-sizing: border-box;
      }
      body {
        width: 8.5in; 
        height: 14in; 
        margin: 0;
        padding: 0;
       	font-family: 'Quicksand', sans-serif;
        background-size: 8.5in 4.5in;
        background-repeat: no-repeat;
    	background-position: bottom; 
		background-image: url('https://s3-us-west-2.amazonaws.com/dmprint/resources/template/html/newcustomer/background.jpg');
      }
      .sig img{
        width:250px;
        height:80px;
      }
      .wrapper {
        position: absolute;
        top: 3.45in;
      }
      .page-content {
        position: relative;
        width: 7in;
        height: 10.625in;
        left: 0.75in;
        top: 0.1875in;
      }
	  pre {
	  	white-space: pre-wrap;
	  }
    </style>
  </head>
  <body>
    <div class="page">
      <div class="page-content">
        <div class="wrapper">
          <p>Hello {{recipientName}},</p>
		  <p>Here is a log message for you from <b>{{appName}}</b>. I hope it's useful.</p>
		  <p><pre>{{logMessage}}</pre></p>
          <br><br><br>
          <p>Sincerely,</p>
          <p>{{appName}} logger<p>
          </div>
      </div>
    </div>
  </body>
</html>`
)
