load "deploy" if respond_to?(:namespace)
require "capistrano_colors"
require "rvm/capistrano"
default_run_options[:pty] = true

# Application
set   :application, "helloworldquiz.com"
set   :project, "langgame"
set   :domain, "helloworldquiz.com"
set   :deploy_to, "/home/krasnoukhov/#{project}/"
role  :web, domain

# Source
set   :scm, "git"
set   :repository, "git@github.com:krasnoukhov/#{project}.git"
set   :branch, "master"
set   :repository_cache, "git"
set   :deploy_via, :remote_cache
set   :git_enable_submodules, 1

# Options
set   :user, "krasnoukhov"
set   :rvm_ruby_string, `cat .ruby-version`
set   :rvm_type, :user
set   :use_sudo, false
set   :keep_releases, 5

# Hooks
after "deploy:update", "deploy:cleanup"
namespace :deploy do
  task :restart, roles: :web do
    run "cd #{current_path} && ./build"
  end
end
