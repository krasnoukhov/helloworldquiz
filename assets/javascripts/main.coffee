App = Ember.Application.create()

App.Router.map ->
  # put your routes here

App.IndexRoute = Ember.Route.extend(
  model: ->
    ["red", "yellow", "blue", "black"]
)

->
  $(".answer-types a").click ->
    $(".answer-types").remove()
    
    if $(this).hasClass("correct")
      $(".result").append('<li><h2>Correct!</h2> <a href="#" class="btn btn-success btn-large">Next</a></li>')
    else
      $(".result").append('<li><h2>Wrong!</h2> <a href="#" class="btn btn-danger btn-large">Next</a></li>')
    
    $(".result").fadeIn()
    false
