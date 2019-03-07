<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.code}}</title>
    <link href="/static/css/error.min.css?v=20181121" rel="stylesheet">
</head>

<body>
    <section id="not-found">
        <div id="title">{{.text}}</div>
        <div class="circles">
            <p>{{.code}}<br><small>{{.msg}}</small></p><span class="circle big"></span> <span class="circle med"></span>
            <span class="circle small"></span>
        </div>
    </section>
</body>

</html>