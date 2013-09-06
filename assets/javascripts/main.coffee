@App = Ember.Application.create()
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
  model: ->
    Ember.$.getJSON "/stats"
  
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
  waitResponse: null
  hasSurvived: (->
    return false unless this.get("response")
    this.get("response").status == "survived"
  ).property("response")
    
  
  load: ->
    self = this
    
    # Load game
    self.set("isLoading", true)
    $.get("/game").always(->
      self.set("isLoading", false)
    ).done((data) ->
      if data.error
        self.transitionToRoute("index")
      else
        self.set("response", data)
    )
  
  result: (data) ->
    if data.correct
      response = $.extend({}, data)
      response.variant = data.correct
      response.variant.live = data.game.lives > 0
      this.set("response", response)
      
      this.set("waitResponse", data)
    else
      this.set("response", data)
    
  actions:
    choose: (option) ->
      self = this
      
      self.set("isLoading", true)
      $.ajax(url: "/game", type: "PUT", data: {
        option: option
      }).always(->
        self.set("isLoading", false)
      ).done((data) ->
        if data.error
          alert(data.error)
        else
          self.result(data)
      )
    
    next: ->
      this.set("response", this.get("waitResponse"))
      this.set("waitResponse", null)
      this.set("correct", null)
    
, Config))

App.StatsController = Ember.ObjectController.extend(Config)
