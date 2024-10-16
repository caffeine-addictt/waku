###### _<div align="right"><sub>// made with <3</sub></div>_

<div align="center">

<!-- Project Banner -->

<a href="https://github.com/{{{ .User }}}/{{{ .Repo }}}">
  <img src="https://github.com/{{{ .User }}}/{{{ .Repo }}}/blob/main/assets/logo.svg" width="750" height="300" alt="">
</a>

<br>

<!-- Badges -->

![badge-workflow]
[![badge-license]][license]
![badge-language]
[![badge-pr]][prs]
[![badge-issues]][issues]

<br><br>

<!-- Description -->

{{{ .Description }}}

<br><br>

---

<!-- TOC -->

**[<kbd> <br> Quick Start <br> </kbd>](#quick-start)**
{{{ if .Docs }}}**[<kbd> <br> Documentation <br> </kbd>]({{{ .Docs }}})**{{{ end }}}
**[<kbd> <br> Thanks <br> </kbd>](#special-thanks)**
**[<kbd> <br> Contribute <br> </kbd>][contribute]**

---

<br>

</div>

# Quick Start

_This is an example of how you can set up your project locally.
To get a local copy up and running, follow these simple steps._

## Prerequisites

- Language 1

## Installation

_Below is an example of how you can install and use {{ .Name }}_

1. Step 1
2. Step 2

<div align="right">
  <br>
  <a href="#-made-with-3"><kbd> <br> 🡅 <br> </kbd></a>
</div>

## Special Thanks

- **[Caffeine-addictt][template-repo]** - _For the repository template_
- **[Img Shields][img-shields]** - _For the awesome README badges_
- **[Hyprland][hyprland]** - _For showing how to make beautiful READMEs_
- **[Hyprdots][hyprdots]** - _For showing how to make beautiful READMEs_

---

![stars-graph]

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[stars-graph]: https://starchart.cc/{{{ .User }}}/{{{ .Repo }}}.svg?variant=adaptive
[prs]: https://github.com/{{{ .User }}}/{{{ .Repo }}}/pulls
[issues]: https://github.com/{{{ .User }}}/{{{ .Repo }}}/issues
[license]: https://github.com/{{{ .User }}}/{{{ .Repo }}}/blob/main/LICENSE

<!---------------- {Links} ---------------->

[contribute]: https://github.com/{{{ .User }}}/{{{ .Repo }}}/blob/main/CONTRIBUTING.md

<!---------------- {Thanks} ---------------->

[template-repo]: https://github.com/caffeine-addictt/waku
[hyprland]: https://github.com/hyprwm/Hyprland
[hyprdots]: https://github.com/prasanthrangan/hyprdots
[img-shields]: https://shields.io

<!---------------- {Badges} ---------------->

[badge-workflow]: https://github.com/{{{ .User }}}/{{{ .Repo }}}/actions/workflows/test-worker.yml/badge.svg
[badge-issues]: https://img.shields.io/github/issues/{{{ .User }}}/{{{ .Repo }}}
[badge-pr]: https://img.shields.io/github/issues-pr/{{{ .User }}}/{{{ .Repo }}}
[badge-language]: https://img.shields.io/github/languages/top/{{{ .User }}}/{{{ .Repo }}}
[badge-license]: https://img.shields.io/github/license/{{{ .User }}}/{{{ .Repo }}}
