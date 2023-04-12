# Lecture 10: Scaling
[Week 10](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_10/README_TASKS.md)

## Add Scaling to the Project
We created an additional droplet on DigitalOcean, and changed our Vagrantfile to duplicate our application, such that we have one running on each of the droplets. Both of these instances are identical. We will then have one of the instances/replicas acting as the primary server, whereas the other will be our backup server. We added the following to our Vagrantfile:
``` ruby
arr = ["webserver-primary", "webserver-backup"]
servers = []
...

arr.each do |name|
    config.vm.define name, primary: false do |server|
        server.vm.provider :digital_ocean do |provider|
            provider.ssh_key_name = ENV["SSH_KEY_NAME"]
            provider.token = ENV["DIGITAL_OCEAN_TOKEN"]
            provider.image = 'ubuntu-22-04-x64'
            provider.region = 'fra1'
            provider.size = 's-1vcpu-1gb'
            provider.privatenetworking = true
        end
...
```

The next step is to setup our DigitalOcean such that our backup replica will take over once the primary replica fails.

## Rolling Updates
Now that we have two replicas of our application, we need to keep them both updated as we push new changes to minitwit. Here multiple different strategies can be used. 

We utilize the *rolling update* strategy where each replica is shut down, updated, then turned back on. In order to increase our availability, however, we will first be updating the backup replica and make it our primary server, then we can shut down the replica which was previously the primary one and update it as well. DigitalOcean will make this easy as it automatically switches the primary server to be the replica which is running.
