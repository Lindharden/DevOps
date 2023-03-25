# Lecture 4: Continuous Integration (CI), Continuous Delivery (CD), and Continuous Deployment
Week 8, 2023

## Step 1: Complete implementing an API for the simulator in your ITU-MiniTwit.

We continued working on implementing the API for the upcoming simulation. To effectively split our workload, we created GitHub Issues for all of the individual endpoints which are to be implemented. We then decided on who was to implement which endpoints, and we then assigned ourselves to the respective Issues. Some of us worked individually, while others worked in smaller groups. To accurately reflect who did what, we made sure to include all co-authors to all commits we made. 

We made an Issue which represented the "main" Issue for implementing the simulation endpoints. This Issue contains a checklist with a checkbox for each individual endpoint. That way we can check off a single endpoint when we feel like that given endpoint is done. All the work on the simulation endpoints was done on a branch called `feature/simulation`, and all the GitHub Issues reference this branch as the target/development branch. We use this naming scheme to clearly distinguish features which are to be implemented, from bugs which are to be fixed. 

After all the checkboxes in the main Issue were checked off, a pull-request was made, which wanted to merge the `feature/simulation` branch into the `main` branch. This pull-request had all members of our group as reviewers, and it could only be merged when it has at least one of these reviewers accept the changes highlighted in the pull-request. This pull-request also links to all of the individual Issues which relate to the simulation endpoints, such that when merging the pull-request all of these Issues will be closed. When our pull-requests need at least one accepting reviewer (different from the author), it acts as a safe-guard which prevents individuals from pushing directly to the main branch.

After some members of our group accepted the changes in the pull-request, we merged the changes, and our minitwit application is now ready for the simulation.

## Step 2: Creating a CI/CD setup for your ITU-MiniTwit.

A CI/CD setup involves a pipeline which automatically builds, tests and deploys our application, once we make changes to it. Continuous Integration (CI) involves automatic building and testing, once changes are pushed to the code repository. This can be extended to involve Continuous Delivery (CD) which means that after the changes have been built and tested, they are also pushed/delivered to some server (e.g. docker) which runs the application. This way the performance can be monitored and logged before the changes are pushed to the customers/users. This can be extended further with Continuous Deployment (CD), which includes the previous steps, but also automatically deploys the new changes directly to the production (source: [simplilearn.com](https://www.simplilearn.com/tutorials/devops-tutorial/continuous-delivery-and-continuous-deployment)).

For our CI/CD setup, we choose to utilize [GitHub Actions](https://github.com/features/actions). This is because:
 - It is integrated in GitHub, which is the version control system we use for our repository. This way we don't have to involve other platforms, and we can keep everything in one place.
 - Since it is integrated in GitHub, we can see the status of individual commits when we push them. This way we can see whether individual commits build (or contain errors), and we can see whether individual commits pass tests. 
 - It allows for Continuous Delivery and Deployment, which means we can automate the delivery or deployment of our changes.

We choose to go for Continuous Deployment, which means that all of our changes are automatically pushed directly to our virtual machine at DigitalOcean, which runs our minitwit application. We do this such that the users always will have access to the newest changes.

To perform the Continuous Deployment, we set up our GitHub Actions as follows:
``` yml
name: Continuous Deployment

on:
  push:
    # Run workflow every time something is pushed to the main branch
    branches:
      - main
  # allow manual triggers for now too
  workflow_dispatch:
    manual: true

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      #Docker hub setup
      - name: Login to Docker Hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      #Prepare docker build
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v1

      #Configure SSH keys for deployment
      - name: Configure SSH
        run: |
          mkdir -p ~/.ssh/
          echo "$SSH_KEY"
          echo "$SSH_KEY" > ~/.ssh/minitwit.key
          chmod 600 ~/.ssh/minitwit.key
          cat >>~/.ssh/config <<END
          Host minitwit
            HostName $SSH_HOST
            User $SSH_USER
            IdentityFile ~/.ssh/minitwit.key
            StrictHostKeyChecking no
          END
        env:
          SSH_USER: ${{ secrets.SSH_USER }}
          SSH_KEY: ${{ secrets.SSH_KEY }}
          SSH_HOST: ${{ secrets.SSH_HOST }}

      #Install needed tooling for the following steps
      - name: Install Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.20.x

      - name: Install Python
        uses: actions/setup-python@v4
        with:
          python-version: "3.10"
      #Checkout repository files
      - name: Checkout
        uses: actions/checkout@v2
      # Run go unit tests
      - name: Unit testing
        run: |
          go test -v
      # Run simulator integration tests
      - name: Simulator tests
        run: |
          pip install -r ./tools/requirements.txt
          pytest ./tools/test_sim_compliance.py

      # Build and push docker image
      - name: Build and push minitwitimage
        uses: docker/build-push-action@v2
        with:
          context: .
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKER_USERNAME }}/minitwitimage:latest
          cache-from: type=registry,ref=${{ secrets.DOCKER_USERNAME }}/minitwitimage:webbuildcache
          cache-to: type=registry,ref=${{ secrets.DOCKER_USERNAME }}/minitwitimage:webbuildcache,mode=max
      - name: Copy docker-compose.yml to server
        run: scp docker-compose.yml minitwit:/vagrant/docker-compose.yml
      
      - name: Copy deploy.sh to server
        run: >-
          chmod +x tools/deploy.sh &&
          scp tools/deploy.sh minitwit:/vagrant/tools/deploy.sh
      - name: Deploy to server
        # Configure the ~./bash_profile file on the Vagrantfile
        run: ssh minitwit '/vagrant/tools/deploy.sh'
      #Run linter to check if new code follows the style guidelines
      - name: lint
        run: if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then exit 1; fi
  release:
    needs: "build"
    name: "Release"
    runs-on: "ubuntu-latest"
    steps:
      - uses: "marvinpinto/action-automatic-releases@latest"
        with:
          repo_token: "${{ secrets.GH_TOKEN }}"
          automatic_release_tag: "latest"
          prerelease: false
          title: "Automatic release"
```

This script sets up an **Ubuntu machine** and installs the required dependencies. It sets up the required **SSH keys**. It performs **unit tests**, and **simulator tests** (tests API for simulation). It **builds** and **pushes** the minitwit application to our Docker Hub. It **moves the docker compose script** to the DigitalOcean server, and **runs the virtual machine**. After this has completed, it creates an automatic release on GitHub, using a script by *marvinpinto*.
