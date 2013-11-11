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
    <meta name="signet:authors" content="Dmitry Krasnoukhov, Slava Zagorodniy">
    <meta name="signet:links" content="https://github.com/krasnoukhov/helloworldquiz, https://twitter.com/krasnoukhov, https://twitter.com/kr3ved">
    
    <script type="text/javascript">
      var Config = {
        title: "{{ .title }}",
        caption: "{{ .caption }}",
        description: "{{ .description }}",
        host: "{{ .host }}",
        reload: function() {
          (function(d, s, id) {
            var js, fjs = d.getElementsByTagName(s)[0];
            js = d.createElement(s); //js.id = id;
            js.src = "//platform.twitter.com/widgets.js";
            fjs.parentNode.insertBefore(js, fjs);
          }(document, 'script', 'twitter-wjs'));
          
          $("#fb-root").remove();
          (function(d, s, id) {
            var js, fjs = d.getElementsByTagName(s)[0];
            js = d.createElement(s); // js.id = id;
            js.src = "//connect.facebook.net/en_US/all.js#xfbml=1&appId=163504373852504";
            js.onload = function() {
              FB.XFBML.parse();
            } 
            fjs.parentNode.insertBefore(js, fjs);
          }(document, 'script', 'facebook-jssdk'));
        }
      }
    </script>
    
    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
      
      ga('create', 'UA-43847693-1', 'helloworldquiz.com');
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
