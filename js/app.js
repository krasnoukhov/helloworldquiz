App = Ember.Application.create();

App.Router.map(function() {
  // put your routes here
});

App.IndexRoute = Ember.Route.extend({
  model: function() {
    return ['red', 'yellow', 'blue', 'black'];
  }
});

$(function() {
  $(".answer-types a").click(function() {
    $(".answer-types").remove();
    
    if($(this).hasClass("correct")) {
      $(".result").append("<li><h2>Correct!</h2> <a href='#' class='btn btn-success btn-large'>Next</a></li>");
    }else{
      $(".result").append("<li><h2>Wrong!</h2> <a href='#' class='btn btn-danger btn-large'>Next</a></li>");
    }
    
    $(".result").fadeIn();
    return false;
  })
});
