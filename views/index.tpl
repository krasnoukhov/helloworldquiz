<!DOCTYPE html>
<html>
  <head>
  <title>Programming Languages Quiz</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href='http://yui.yahooapis.com/pure/0.2.1/buttons-min.css' rel='stylesheet' type='text/css'>
    <link href='http://fonts.googleapis.com/css?family=Gafata' rel='stylesheet' type='text/css'>
    {{ "app" | stylesheet_tag }}
    {{ "main" | stylesheet_tag }}
    {{ "app" | javascript_tag }}
    {{ "main" | javascript_tag }}
  </head>
  <body>
    <div id="fb-root"></div>
    <script>(function(d, s, id) {
      var js, fjs = d.getElementsByTagName(s)[0];
      if (d.getElementById(id)) return;
      js = d.createElement(s); js.id = id;
      js.src = "//connect.facebook.net/en_US/all.js#xfbml=1&appId=1374335546120479";
      fjs.parentNode.insertBefore(js, fjs);
    }(document, 'script', 'facebook-jssdk'));</script>
    <!-- <div class="hero-unit">
      <div class="inner">
        <h1>Programming Languages Quiz</h1>
        <p>Recognize programming language by code snippet</p>
        <a class="pure-button pure-button-secondary pure-button-large" href="game.html">Play</a>
      </div>
    </div> -->
    
    <div class="container">
      <h1 class="clearfix">Programming Languages Quiz <a href="#">Stats</a></h1>
      <div class="status">
        <span>Lives: 3</span>
        <span>Score: 100</span>
      </div>
      <pre>
        <code>
          package org.example.dummy {
              import org.dummy.*;
    
              /*define package inline interface*/
              public interface IFooBarzable {
                  public function foo(... pairs):Array;
              }
    
              public class FooBar implements IFooBarzable {
                  static private var cnt:uint = 0;
                  private var bar:String;
    
                  //constructor
                  public function TestBar(bar:String):void {
                      bar = bar;
                      ++cnt;
                  }
    
                  public function foo(... pairs):Array {
                      pairs.push(bar);
                      return pairs;
                  }
              }
          }
        </code>
      </pre>
      
      
      <div class="answers">
        <ul class="answer-types">
          <li><a class="pure-button pure-button-secondary" href="#">C++</a></li>
          <li><a class="pure-button pure-button-secondary" href="#">Java</a></li>
          <li><a class="pure-button pure-button-secondary correct" href="#">ActionScript</a></li>
        </ul>
        <ul class="result"></ul>
      </div>
    </div>
    
    <footer class="container clearfix">
      <p class="pull-left">&copy; <a href="https://github.com/krasnoukhov">Dmitry Krasnoukhov</a> & <a href="https://github.com/kr3ved">Slava Zagorodniy</a></p>
      <div class="fb-like pull-right" data-href="http://developers.facebook.com/docs/reference/plugins/like" data-width="450" data-layout="button_count" data-show-faces="true" data-send="false"></div>
      <a href="https://twitter.com/share" class="twitter-share-button pull-right" data-text="Recognize Programming Language game">Tweet</a>
      <script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+'://platform.twitter.com/widgets.js';fjs.parentNode.insertBefore(js,fjs);}}(document, 'script', 'twitter-wjs');</script>
      
    </footer>
  </body>
</html>
