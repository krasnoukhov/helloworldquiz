App = Ember.Application.create()
App.Router.map ->
  this.route("game")
  this.route("stats")

App.IndexRoute = Ember.Route.extend()
App.GameRoute = Ember.Route.extend(
  setupController: (controller) ->
    controller.load()
    
  renderTemplate: ->
    this.render("game")
)
App.StatsRoute = Ember.Route.extend(
  renderTemplate: ->
    this.render("stats")
)

App.IndexController = Ember.ObjectController.extend($.extend(
  isLoading: false
  
  actions:
    play: ->
      self = this
      
      # Create game
      self.set("isLoading", true)
      $.post("/game").always(->
        self.set("isLoading", false)
      ).done((data) ->
        controller = self.controllerFor("game")
        controller.set("response", data)
        self.transitionToRoute("game")
      )
      
, Config))

App.GameController = Ember.ObjectController.extend($.extend(
  isLoading: false
  response: null
  
  load: ->
    self = this
    
    # Load game
    self.set("isLoading", true)
    $.get("/game").always(->
      self.set("isLoading", false)
    ).done((data) ->
      if data.Error
        self.transitionToRoute("index")
      else
        self.set("response", data)
    )
  
  actions:
    lol: ->
      # ...
      
, Config))

App.StatsController = Ember.ObjectController.extend(Config)

$ ->
  $(".answer-types a").click ->
    $(".answer-types").remove()
    
    if $(this).hasClass("correct")
      $(".result").append('<li><h2>Correct!</h2> <a href="#" class="pure-button pure-button-success">Next</a></li>')
    else
      $(".result").append('<li><h2>Wrong!</h2> <a href="#" class="pure-button pure-button-error">Next</a></li>')
    
    $(".result").fadeIn()
    false
