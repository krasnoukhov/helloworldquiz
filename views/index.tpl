<!DOCTYPE html>
<html>
  <head>
    <title>{{ .title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="shortcut icon" href="http://{{ .host }}/public/favicon.ico" />
    <link rel="icon" sizes="16x16 32x32" href="http://{{ .host }}/public/favicon.ico">
    <meta property="og:title" content="{{ .caption }}"/>
    <meta property="og:description" content="{{ .description }}"/>
    <meta property="og:url" content="http://{{ .host }}"/>
    <meta property="og:image" content="http://{{ .host }}/public/icon.png"/>
    
    <script type="text/javascript">
      var Config = {
        title: "{{ .title }}",
        caption: "{{ .caption }}",
        description: "{{ .description }}",
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
