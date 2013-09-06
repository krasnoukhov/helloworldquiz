<!DOCTYPE html>
<html>
  <head>
  <title>{{ .title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    
    <script type="text/javascript">
      var Config = {
        title: "{{ .title }}",
        host: "{{ .host }}"
      }
    </script>
    
    {{ "app" | stylesheet_tag }}
    {{ "main" | stylesheet_tag }}
    {{ "app" | javascript_tag }}
    {{ "main" | javascript_tag }}
    <link href="http://fonts.googleapis.com/css?family=Gafata" rel="stylesheet" type="text/css">
  </head>
  <body>
    {{ str2html app }}
  </body>
</html>
