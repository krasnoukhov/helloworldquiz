<!DOCTYPE html>
<html>
  <head>
  <title>{{ .title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    {{ "app" | stylesheet_tag }}
    {{ "main" | stylesheet_tag }}
    {{ "app" | javascript_tag }}
    {{ "main" | javascript_tag }}
    <link href="http://fonts.googleapis.com/css?family=Gafata" rel="stylesheet" type="text/css">
  </head>
  <body>
    <!-- <div class="hero-unit">
      <div class="inner">
        <h1>{{ .title }}</h1>
        <p>Recognize programming language by code snippet</p>
        <a class="pure-button pure-button-secondary pure-button-large" href="game.html">Play</a>
      </div>
    </div> -->
    
    <div class="container">
      <h1 class="clearfix">{{ .title }} <a href="#">Stats</a></h1>
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
      <p class="pull-left">
        &copy; <a href="http://krasnoukhov.com/">Dmitry Krasnoukhov</a> & <a href="https://github.com/kr3ved">Slava Zagorodniy</a>
      </p>
      
      <iframe class="pull-right" src="//www.facebook.com/plugins/like.php?href=http%3A%2F%2F{{ .host }}&amp;width=107&amp;height=20&amp;colorscheme=light&amp;layout=button_count&amp;action=like&amp;show_faces=true&amp;send=false" allowtransparency="true" frameborder="0" scrolling="no" style="border:none; overflow:hidden; width:107px; height:20px;"></iframe>
      <iframe class="pull-right" src="http://platform.twitter.com/widgets/tweet_button.1378258117.html#_=1378492153846&amp;count=horizontal&amp;id=twitter-widget-0&amp;lang=en&amp;size=m&amp;text={{ .title }}&amp;url=http%3A%2F%2F{{ .host }}%2F" allowtransparency="true" frameborder="0" scrolling="no" style="border:none; overflow:hidden; width:107px; height:20px;"></iframe>
      <iframe class="pull-right" src="http://ghbtns.com/github-btn.html?user=krasnoukhov&repo=langgame&type=watch&count=true" allowtransparency="true" frameborder="0" scrolling="no" style="border:none; overflow:hidden; width:107px; height:20px;"></iframe>
    </footer>
  </body>
</html>
