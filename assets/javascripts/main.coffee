@App = Ember.Application.create()
Config.url = "http://#{Config.host}"

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

App.ApplicationController = Ember.Controller.extend(
  routeChanged: (->
    return unless ga
    
    self = this
    Em.run.next ->
      ga("send", "pageview", "/#{self.get("currentPath")}")
      Config.reload()
    
  ).observes("currentPath")
)

App.IndexController = Ember.ObjectController.extend($.extend(
  needs: "game"
  isLoading: false
  
  actions:
    play: ->
      self = this
      
      # Create game
      self.set("isLoading", true)
      $.post("/game").always(->
        self.set("isLoading", false)
      ).done((data) ->
        controller = self.get("controllers.game")
        controller.set("response", data)
        self.transitionToRoute("game")
      ).fail((xhr, status, error) ->
        alert("Error #{xhr.status}: #{xhr.statusText}. Try refreshing a page")
      )
      
, Config))

App.GameController = Ember.ObjectController.extend($.extend(
  isLoading: false
  response: null
  correct: null
  waitResponse: null
  
  game: (->
    return this.get("waitResponse").game if this.get("waitResponse")
    return this.get("response").game if this.get("response")
    {}
  ).property("response", "waitResponse")
  
  isLive: (->
    return false if this.get("waitResponse") && this.get("waitResponse").game.lives == 0
    return false if this.get("response") && this.get("response").game.lives == 0
    return true
  ).property("response", "waitResponse")
  
  hasSurvived: (->
    return false unless this.get("response")
    this.get("response").status == "survived"
  ).property("response")
  
  highlightedSnippet: (->
    return "" unless this.get("response")
    snippet = this.get("response").variant.snippet
    $("<div />").html(hljs.highlightAuto(snippet).value).html()
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
    ).fail((xhr, status, error) ->
      alert("Error #{xhr.status}: #{xhr.statusText}. Try refreshing a page")
    )
  
  result: (data) ->
    ga("send", "pageview", "/game")
    
    if data.correct
      this.set("correct", data.correct)
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
      ).fail((xhr, status, error) ->
        alert("Error #{xhr.status}: #{xhr.statusText}. Try refreshing a page")
      )
    
    next: ->
      this.set("response", this.get("waitResponse"))
      this.set("correct", null)
      this.set("waitResponse", null)
    
, Config))

App.StatsController = Ember.ObjectController.extend(Config)
