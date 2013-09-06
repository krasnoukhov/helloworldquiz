App = Ember.Application.create()
App.Router.map ->
  this.route("game")
  this.route("stats")

App.IndexController = Ember.ObjectController.extend(Config)
App.GameController = Ember.ObjectController.extend(Config)
App.StatsController = Ember.ObjectController.extend(Config)

App.IndexRoute = Ember.Route.extend()
App.GameRoute = Ember.Route.extend(
  renderTemplate: ->
    this.render("game")
)
App.StatsRoute = Ember.Route.extend(
  renderTemplate: ->
    this.render("stats")
)

$ ->
  $(".answer-types a").click ->
    $(".answer-types").remove()
    
    if $(this).hasClass("correct")
      $(".result").append('<li><h2>Correct!</h2> <a href="#" class="pure-button pure-button-success">Next</a></li>')
    else
      $(".result").append('<li><h2>Wrong!</h2> <a href="#" class="pure-button pure-button-error">Next</a></li>')
    
    $(".result").fadeIn()
    false
