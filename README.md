<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![Build][build-shield]][build-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/lindharden/devops">
    <img src="/static/img/minitwit.png" alt="Logo" width="200" height="200">
  </a>
<br />
<br />
  <p align="center">
    An application maintained during the DevOps course at the IT University of Copenhagen
    <br />
    <a href="http://157.230.76.157:8080/public">Deployment</a>
    ·
    <a href="https://github.com/Lindharden/DevOps/issues/new/choose">Report Bug</a>
    ·
    <a href="https://github.com/Lindharden/DevOps/issues/new/choose">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

The project is built follwing the course outlined in the following repository: <https://github.com/itu-devops/lecture_notes> 
<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

The project is built with the following languages and frameworks
* [![Golang][Golang-badge]][Golang-url]
* [![Gin web framework][Gin-badge]][Gin-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* golang > 1.16
  ```sh
  https://go.dev/
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/Lindharden/DevOps.git
   ```
2. Install dependencies
   ```sh
   go mod download
   go mod tidy
   ```
3. Run the project
   ```sh
   go run minitwit.go
   ```
### Running in docker
   ```sh
   docker compose up -d
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/lindharden/devops/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

See the contributing guide for this project in [CONTRIBUTING](https://github.com/Lindharden/DevOps/blob/main/CONTRIBUTING.md).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the GPL License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Jeppe Lindhard - [@lindharden](https://github.com/Lindharden) - jepli@itu.dk

Johan Flensmark - [@bitknox](https://github.com/bitknox) - jokf@itu.dk

Nikolaj Sørensen - [@nikso-itu](https://github.com/nikso-itu) - nikso@itu.dk

Benjamin Hammer Thygesen - [@Mansin-itu](https://github.com/Mansin-ITU) - beth@itu.dk

Daniel Gjesse Kjellberg - [@KjellbergD](https://github.com/KjellbergD) - dakj@itu.dk

Project Link: [https://github.com/Lindharden/DevOps](https://github.com/Lindharden/DevOps)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

[DALL-E 2](https://openai.com/product/dall-e-2)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/lindharden/devops?style=for-the-badge
[contributors-url]: https://github.com/Lindharden/DevOps/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/lindharden/devops?style=for-the-badge
[forks-url]: https://github.com/Lindharden/DevOps/forks
[stars-shield]: https://img.shields.io/github/stars/lindharden/devops?style=for-the-badge
[stars-url]: https://github.com/Lindharden/DevOps/stargazers
[issues-shield]: https://img.shields.io/github/issues/lindharden/devops?style=for-the-badge
[issues-url]: https://github.com/Lindharden/DevOps/issues
[build-shield]: https://img.shields.io/github/actions/workflow/status/lindharden/devops/cd.yml?style=for-the-badge&logo=appveyor
[build-url]: https://github.com/Lindharden/DevOps/actions/workflows/cd.yml
[license-shield]: https://img.shields.io/github/license/lindharden/devops?style=for-the-badge
[license-url]: https://github.com/Lindharden/DevOps/blob/main/LICENSE
[product-screenshot]: images/screenshot.png
[Golang-url]: https://go.dev/
[Golang-badge]: https://img.shields.io/badge/golang-29BEB0?style=for-the-badge&logo=go&logoColor=white
[Gin-url]: https://go.dev/
[Gin-badge]: https://img.shields.io/badge/gin-29BEB0?style=for-the-badge
