<!--
*** Thanks for checking out this README Template. If you have a suggestion that would
*** make this better, please fork the repo and create a pull request or simply open
*** an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->





<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <!-- <a href="https://github.com/i-am-jeetu/gogeoip">
    <img src="/home/jeetu/Downloads/laptop.svg" alt="Logo" width="80" height="80">
  </a> -->

  <h3 align="center">GoGeoIp</h3>

  <p align="center"> GoGeoIP is an unofficial, simple, authentication free, client for <a href="https://ip-api.com/">ip-api.com</a>
 <br>
Do not use in production as the limit is very low i.e 45  requests/minute
    <br />
    <a href="HREF:TODOCS"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/i-am-jeetu/gogeoip/issues">Report Bug</a>
    ·
    <a href="https://github.com/i-am-jeetu/gogeoip/issues">Request Feature</a>
  </p>
</p>



<!-- ABOUT THE PROJECT -->
<!-- ## About The Project -->



### See Also
* [Official Website](https://ip-api.com)



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

## Usage 
There is nothing much more here, apart from this sample code and a bunch of self explanatory functions. 
```go
package main

import (
	"fmt"

	gogeoip "github.com/i-am-jeetu/gogeoip"
)

func main() {

	ipStats, err := gogeoip.FromString("117.251.211.81")
	
	// Alternatively
	// ipAddr := net.IP{117, 251, 211}
	// ipStats, err := gogeoip.FromIP(&ipAddr)

	if err != nil {
		// Do Something with error
		fmt.Errorf("%w", err)
	}
	fmt.Print(ipStats.GetZip()) // Prints 811202
}
```

Functions Supported are - 
- GetCountryCode()
- GetRegion()
- GetRegionName()
- GetCountry()
- GetZip()
- GetLongitude()
- GetLatitude()
- GetIsp()
- GetCity()
- GetTimezone()
- GetRegionName()


### Installation
You can install it using the following command -

```sh
go get github.com/i-am-jeetu/gogeoip 
```


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact
[Jeetu Kumar](https://www.linkedin.com/in/img2/) 


<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
<div>Icon made by <a href="https://www.flaticon.com/authors/smashicons" title="Smashicons">Smashicons</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a></div>
<!-- * [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Img Shields](https://shields.io)
* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Pages](https://pages.github.com)
* [Animate.css](https://daneden.github.io/animate.css)
* [Loaders.css](https://connoratherton.com/loaders)
* [Slick Carousel](https://kenwheeler.github.io/slick)
* [Smooth Scroll](https://github.com/cferdinandi/smooth-scroll)
* [Sticky Kit](http://leafo.net/sticky-kit)
* [JVectorMap](http://jvectormap.com)
* [Font Awesome](https://fontawesome.com) -->





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/i-am-jeetu/gogeoip?style=flat-square
[contributors-url]: https://github.com/i-am-jeetu/gogeoip/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/i-am-jeetu/gogeoip.svg?style=flat-square
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=flat-square
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=flat-square
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=flat-square
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/img2/
[product-screenshot]: images/screenshot.png
