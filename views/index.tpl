<!DOCTYPE html>
<html>
  <head>
  <title>Programming Languages Quiz</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    {{ "app" | stylesheet_tag }}
    {{ "main" | stylesheet_tag }}
    {{ "app" | javascript_tag }}
    {{ "main" | javascript_tag }}
  </head>
  <body>
    <div class="hero-unit">
      <div class="inner">
        <h1>Programming Languages Quiz</h1>
        <p>Recognize programming language by code snippet</p>
        <a class="btn btn-info btn-large" href="game.html">Play</a>
      </div>
    </div>
    
    <!--
      <div class="container">
        <h1>Programming Languages Quiz</h1>
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
            <li><a class="btn btn-primary btn-large" href="#">C++</a></li>
            <li><a class="btn btn-primary btn-large" href="#">Java</a></li>
            <li><a class="btn btn-primary btn-large correct" href="#">ActionScript</a></li>
          </ul>
          <ul class="result"></ul>
        </div>
        
        
      </div>
    -->
  </body>
</html>
