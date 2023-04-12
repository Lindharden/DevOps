# Lecture 3: Provision of local and remote virtual machines
Week 7, 2023

## New github setup

Before we started on any work this week, we decided to make some changes to our Github repository. The first thing we did was that we started utilizing the 'issues' feature in Git. This is very similar to many other Kanban tools like *Jira* or *Trello*, where we can create backlog items known as *issues*, and then assign each other to different tasks. 

The reason we decided on picking Github issues over other Kanban solutions, was because of the tight integration between Git and the Kanban board itself. Github issues allow us to natively link and mention different things inside of our Github repository, such as branches. It's also much simpler to just use one platform for everything, instead of having to switch back and forth between different platforms while working. Meanwhile Github issues support all of the features we need, such as:
 - Creating issues/backlog items.
 - Creating milestones, and linking issues to milestones.
 - Assigning authors to different issues.

 We also changed some rules in our Github repository such that no one can push directly to the main branch. From now on everything has to be implemented on separate branches and then we have to create pull requests, which have to be peer-reviewed, before the content can be merged into the main branch.

## Step 1: Implement an API for the simulator in your ITU-MiniTwit.

To implement the API for the simulator, we went through the simulator file [minitwit_sim_api.py](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_03/API_Spec/minitwit_sim_api.py), and looked at all the routes we have to implement. We then created the issue [Feature: Implement the endpoints for the simulator](https://github.com/Lindharden/DevOps/issues/13), and some sub-issues which are described inside the main issue.

We then implemented the required controllers and handlers, to support all the new API calls.

## Step 2: Continue refactoring of your ITU-MiniTwit.

We created some issues for the missing features in our application (not related to simulation). 

We need to deploy our minitwit platform to a remote virtual machine. We deploy the virtual machine to DigitalOcean. We deploy to DigitalOcean as we get free credit from ITU. The tool we used to create the virtual machine is *Vagrant*. Vagrant is a tool that enables the creation and configuration of lightweight, reproducible and portable virtual machines. We choose to use Vagrant over other tools for creating virtual machines, such as *Docker* or *VirtualBox*, because:
 - Vagrant makes virtual machines easy and fast to setup, as it can be through scripts.
 - Vagrant works on all platforms, that is: Windows, macOS and Linux.
 - Virtual machines created from Vagrant are easy to reproduce, as it's done through a script.
 - Vagrant is very flexible and can utilize a lot of different containerization platforms, and integrates well with different development ecosystems.

We created a Vagrant file for our application, using the following script:
``` ruby
Vagrant.configure("2") do |config|
  config.env.enable
  config.vm.box = 'digital_ocean'
  config.vm.box_url = "https://github.com/devopsgroup-io/vagrant-digitalocean/raw/master/box/digital_ocean.box"
  config.ssh.private_key_path = '~/.ssh/id_rsa'
  config.vm.synced_folder ".", "/vagrant", type: "rsync"

  config.vm.define "webserver", primary: false do |server|

    server.vm.provider :digital_ocean do |provider|
      provider.ssh_key_name = ENV["SSH_KEY_NAME"]
      provider.token = ENV["DIGITAL_OCEAN_TOKEN"]
      provider.image = 'ubuntu-22-04-x64'
      provider.region = 'fra1'
      provider.size = 's-1vcpu-1gb'
      provider.privatenetworking = true
    end

    server.vm.hostname = "webserver"

    server.vm.provision "shell", inline: <<-SHELL
      sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
      cd /vagrant
      docker compose up -d
      echo "================================================================="
      echo "=                            DONE                               ="
      echo "================================================================="
      echo "Navigate in your browser to:"
      THIS_IP=`hostname -I | cut -d" " -f1`
    SHELL
  end
  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    sudo apt-get update
    sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
    sudo mkdir -m 0755 -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt update
  SHELL
end
```

In the script we provide the required information for DigitalOcean. Now we can deploy the virtual machine on DigitalOcean using the command: `vagrant up`. The application is now deployed here: <http://164.90.223.49:8080/public>.
